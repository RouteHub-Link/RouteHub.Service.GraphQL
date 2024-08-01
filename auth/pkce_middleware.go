package auth

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	applicationConfig "github.com/RouteHub-Link/routehub-service-graphql/config"
	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/oauth2"
)

var (
	oauth2Config        *oauth2.Config
	provider            *oidc.Provider
	oidcConfig          *oidc.Config
	verifier            *oidc.IDTokenVerifier
	authorizerConfig    applicationConfig.AuthorizerConfig
	onceConfigureOauth2 sync.Once
)

type IdTokenClaims struct {
	Sub   string `json:"sub"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func PKCEMiddleware(next http.Handler) http.Handler {
	ConfigureOauth(next)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		//check if the cookie is set
		cookie, err := r.Cookie("token")

		if err != nil {
			setCallbackCookie(w, r, "token", "")
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			return
		}

		oauth2Token, err := oauth2Config.Exchange(ctx, cookie.Value)
		if err != nil {
			http.Error(w, "Failed to exchange token: "+err.Error(), http.StatusInternalServerError)
			return
		}
		rawIDToken, ok := oauth2Token.Extra("id_token").(string)
		if !ok {
			http.Error(w, "No id_token field in oauth2 token.", http.StatusInternalServerError)
			return
		}
		idToken, err := verifier.Verify(ctx, rawIDToken)
		if err != nil {
			http.Error(w, "Failed to verify ID Token: "+err.Error(), http.StatusInternalServerError)
			return
		}

		nonce, err := r.Cookie("nonce")
		if err != nil {
			http.Error(w, "nonce not found", http.StatusBadRequest)
			return
		}
		if idToken.Nonce != nonce.Value {
			http.Error(w, "nonce did not match", http.StatusBadRequest)
			return
		}
		resp := struct {
			OAuth2Token   *oauth2.Token
			IDTokenClaims *json.RawMessage // ID Token payload is just JSON.
			CustomClaims  *IdTokenClaims
		}{oauth2Token, new(json.RawMessage), new(IdTokenClaims)}

		if err := idToken.Claims(&resp.IDTokenClaims); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := idToken.Claims(&resp.CustomClaims); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		log.Printf("PKCEMiddleware: %+v", resp)

		user := new(UserSession)
		user.ParseFromIdTokenClaims(resp.CustomClaims)

		newCtx := context.WithValue(r.Context(), userCtxKey, user)
		log.Printf("user session on context \nname : %s\nuser : %+v ", userCtxKey, user)

		r = r.WithContext(newCtx)
		next.ServeHTTP(w, r)
	})
}

func ConfigureOauth(h http.Handler) (err error) {
	onceConfigureOauth2.Do(func() {
		ctx := context.Background()
		_applicationConfig := applicationConfig.ConfigurationService{}.Get()
		authorizerConfig = _applicationConfig.AuthorizerConfig

		provider, err = oidc.NewProvider(ctx, authorizerConfig.Issuer)
		if err != nil {
			log.Fatal(err)
		}

		oidcConfig = &oidc.Config{
			ClientID: authorizerConfig.ClientID,
		}

		verifier = provider.Verifier(oidcConfig)

		host := strings.Join([]string{"http://localhost:", _applicationConfig.GraphQL.PortAsString}, "")
		redirectURL := strings.Join([]string{host, authorizerConfig.Callback}, "")

		oauth2Config = &oauth2.Config{
			ClientID:     authorizerConfig.ClientID,
			ClientSecret: authorizerConfig.ClientSecret,
			RedirectURL:  redirectURL,
			Endpoint: oauth2.Endpoint{
				AuthURL:  authorizerConfig.AuthorizerURL,
				TokenURL: authorizerConfig.TokenURL,
			},
			Scopes: authorizerConfig.Scopes,
		}

		http.Handle("/oauth2/login", OauthHttpHandler())
		http.Handle(authorizerConfig.Callback, callbackFunction())
	})

	return
}

func OauthHttpHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		setCallbackCookie(w, r, "state", "")
		setCallbackCookie(w, r, "nonce", "")

		state, err := randString(16)
		if err != nil {
			http.Error(w, "Internal error", http.StatusInternalServerError)
			return
		}
		nonce, err := randString(16)
		if err != nil {
			http.Error(w, "Internal error", http.StatusInternalServerError)
			return
		}

		setCallbackCookie(w, r, "state", state)
		setCallbackCookie(w, r, "nonce", nonce)

		http.Redirect(w, r, oauth2Config.AuthCodeURL(state, oidc.Nonce(nonce)), http.StatusFound)
	}
}

func callbackFunction() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		state, err := r.Cookie("state")
		if err != nil {
			http.Error(w, "state not found", http.StatusBadRequest)
			return
		}

		if r.FormValue("state") != state.Value {
			http.Error(w, "state did not match", http.StatusBadRequest)
			return
		}

		if r.URL.Query().Get("code") == "" {
			http.Error(w, "code not found", http.StatusBadRequest)
			return
		}

		code := r.URL.Query().Get("code")

		oauth2Token, err := oauth2Config.Exchange(ctx, code)
		if err != nil {
			http.Error(w, "Failed to exchange token: "+err.Error(), http.StatusInternalServerError)
			return
		}
		rawIDToken, ok := oauth2Token.Extra("id_token").(string)
		if !ok {
			http.Error(w, "No id_token field in oauth2 token.", http.StatusInternalServerError)
			return
		}
		idToken, err := verifier.Verify(ctx, rawIDToken)
		if err != nil {
			http.Error(w, "Failed to verify ID Token: "+err.Error(), http.StatusInternalServerError)
			return
		}

		nonce, err := r.Cookie("nonce")
		if err != nil {
			http.Error(w, "nonce not found", http.StatusBadRequest)
			return
		}
		if idToken.Nonce != nonce.Value {
			http.Error(w, "nonce did not match", http.StatusBadRequest)
			return
		}

		oauth2Token.AccessToken = "*REDACTED*"

		resp := struct {
			OAuth2Token   *oauth2.Token
			IDTokenClaims *json.RawMessage // ID Token payload is just JSON.
		}{oauth2Token, new(json.RawMessage)}

		if err := idToken.Claims(&resp.IDTokenClaims); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		setCallbackCookie(w, r, "token", rawIDToken)

		data, err := json.MarshalIndent(resp, "", "    ")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		//w.Write(data)
		log.Printf("Callback: %s", data)

		http.Redirect(w, r, "/playground", http.StatusTemporaryRedirect)
	}
}

func randString(nByte int) (string, error) {
	b := make([]byte, nByte)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(b), nil
}

func setCallbackCookie(w http.ResponseWriter, r *http.Request, name, value string) {
	c := &http.Cookie{
		Name:     name,
		Value:    value,
		MaxAge:   int(time.Hour.Seconds()),
		Secure:   r.TLS != nil,
		HttpOnly: true,
		Path:     "/",
	}
	http.SetCookie(w, c)
}

/*
http://localhost:8080/authorize?
client_id=b20008b0-adec-499a-ac91-272d4c92cdcb
&redirect_uri=https%3A%2F%2Foauthdebugger.com%2Fdebug&scope=openid%20email%20profile
&response_type=code
&response_mode=form_post
&code_challenge_method=S256
&code_challenge=6KS3YloVy7KrM-FuPwC-H8kx3mrTfGT_USXI7aJYd3M
&state=cjx6cdpopi
&nonce=4xxtqfw9lql
*/
