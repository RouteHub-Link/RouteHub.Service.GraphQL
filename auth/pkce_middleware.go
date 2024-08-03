package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	applicationConfig "github.com/RouteHub-Link/routehub-service-graphql/config"
	"github.com/google/uuid"
	"github.com/zitadel/oidc/v3/pkg/client/rp"
	httphelper "github.com/zitadel/oidc/v3/pkg/http"
	"github.com/zitadel/oidc/v3/pkg/oidc"
)

var (
	provider            rp.RelyingParty
	authConfig          applicationConfig.AuthConfig
	onceConfigureOauth2 sync.Once
)

func PKCEMiddleware(next http.Handler) http.Handler {
	ConfigureOauth(next)

	return rp.CodeExchangeHandler(rp.UserinfoCallback(func(w http.ResponseWriter, r *http.Request, tokens *oidc.Tokens[*oidc.IDTokenClaims], state string, rp rp.RelyingParty, info *oidc.UserInfo) {
		if info == nil {
			next.ServeHTTP(w, r)
		}

		user := new(UserSession)
		user.ParseFromIdTokenClaims(info)

		newCtx := context.WithValue(r.Context(), userCtxKey, user)
		log.Printf("user session on context \nname : %s\nuser : %+v ", userCtxKey, user)

		r = r.WithContext(newCtx)
		next.ServeHTTP(w, r)

	}), provider)
	//return authenticationInterceptor.CheckAuthentication()(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//
	//	if authentication.IsAuthenticated(r.Context()) {
	//		authCtx := authenticationInterceptor.Context(r.Context())
	//		userInfo := authCtx.GetUserInfo()
	//
	//		user := new(UserSession)
	//		user.ParseFromIdTokenClaims(userInfo)
	//
	//		newCtx := context.WithValue(r.Context(), userCtxKey, user)
	//		log.Printf("user session on context \nname : %s\nuser : %+v \n userInfo: %+v ", userCtxKey, user, userInfo)
	//
	//		r = r.WithContext(newCtx)
	//	}
	//
	//	next.ServeHTTP(w, r)
	//}))

	//return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//ctx := r.Context()
	//
	////check if the cookie is set
	//cookie, err := r.Cookie("token")
	//
	//if err != nil {
	//	setCallbackCookie(w, r, "token", "")
	//	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	//	return
	//}
	//
	//oauth2Token, err := oauth2Config.Exchange(ctx, cookie.Value)
	//if err != nil {
	//	http.Error(w, "Failed to exchange token: "+err.Error(), http.StatusInternalServerError)
	//	return
	//}
	//rawIDToken, ok := oauth2Token.Extra("id_token").(string)
	//if !ok {
	//	http.Error(w, "No id_token field in oauth2 token.", http.StatusInternalServerError)
	//	return
	//}
	//idToken, err := verifier.Verify(ctx, rawIDToken)
	//if err != nil {
	//	http.Error(w, "Failed to verify ID Token: "+err.Error(), http.StatusInternalServerError)
	//	return
	//}
	//
	//nonce, err := r.Cookie("nonce")
	//if err != nil {
	//	http.Error(w, "nonce not found", http.StatusBadRequest)
	//	return
	//}
	//if idToken.Nonce != nonce.Value {
	//	http.Error(w, "nonce did not match", http.StatusBadRequest)
	//	return
	//}
	//resp := struct {
	//	OAuth2Token   *oauth2.Token
	//	IDTokenClaims *json.RawMessage // ID Token payload is just JSON.
	//	CustomClaims  *IdTokenClaims
	//}{oauth2Token, new(json.RawMessage), new(IdTokenClaims)}
	//
	//if err := idToken.Claims(&resp.IDTokenClaims); err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}
	//
	//if err := idToken.Claims(&resp.CustomClaims); err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}
	//
	//log.Printf("PKCEMiddleware: %+v", resp)
	//
	//user := new(UserSession)
	//user.ParseFromIdTokenClaims(resp.CustomClaims)
	//
	//newCtx := context.WithValue(r.Context(), userCtxKey, user)
	//log.Printf("user session on context \nname : %s\nuser : %+v ", userCtxKey, user)
	//
	//r = r.WithContext(newCtx)
	//next.ServeHTTP(w, r)
	//})
}

func ConfigureOauth(h http.Handler) (err error) {
	onceConfigureOauth2.Do(func() {
		ctx := context.Background()
		_applicationConfig := applicationConfig.ConfigurationService{}.Get()
		authConfig = _applicationConfig.AuthConfig

		host := strings.Join([]string{"http://localhost:", _applicationConfig.GraphQL.PortAsString}, "")
		redirectURL := strings.Join([]string{host, authConfig.Callback}, "")

		cookieHandler := httphelper.NewCookieHandler([]byte(authConfig.CookieKey), []byte(authConfig.CookieKey), httphelper.WithUnsecure())

		options := []rp.Option{
			rp.WithCookieHandler(cookieHandler),
			rp.WithVerifierOpts(rp.WithIssuedAtOffset(5 * time.Second)),
			rp.WithPKCE(cookieHandler),
		}

		provider, err = rp.NewRelyingPartyOIDC(ctx, authConfig.Issuer, authConfig.ClientID, authConfig.ClientSecret, redirectURL, authConfig.Scopes, options...)
		if err != nil {
			log.Fatal(err)
		}

		state := func() string {
			return uuid.New().String()
		}

		urlOptions := []rp.URLParamOpt{
			rp.WithPromptURLParam("Welcome back!"),
		}

		if authConfig.ResponseMode != "" {
			urlOptions = append(urlOptions, rp.WithResponseModeURLParam(oidc.ResponseMode(authConfig.ResponseMode)))
		}

		http.Handle("/oauth2/login", rp.AuthURLHandler(
			state,
			provider,
			urlOptions...,
		))

		http.Handle(authConfig.Callback, rp.CodeExchangeHandler(rp.UserinfoCallback(marshalUserinfo), provider))

		// Initiate the authentication by providing a zitadel configuration and handler.
		// This example will use OIDC/OAuth2 PKCE Flow, therefore you will also need to initialize that with the generated client_id:

		//if authConfig.Insecure {
		//	zitadelInstance = zitadel.New(authConfig.Domain, zitadel.WithInsecure(authConfig.Port))
		//} else {
		//	zitadelInstance = zitadel.New(authConfig.Domain)
		//}
		//
		//authN, err = authentication.New(ctx, zitadelInstance, authConfig.JsonTokenPrivateKey,
		//	openid.DefaultAuthentication(authConfig.ClientID, redirectURL, authConfig.JsonTokenPrivateKey),
		//)
		//if err != nil {
		//	slog.Error("zitadel sdk could not initialize", "error", err)
		//	os.Exit(1)
		//}
		//
		//authenticationInterceptor = authentication.Middleware(authN)
		//
		//http.Handle("/authN/", authN)
	})

	return
}

func marshalToken(w http.ResponseWriter, r *http.Request, tokens *oidc.Tokens[*oidc.IDTokenClaims], state string, rp rp.RelyingParty) {
	data, err := json.Marshal(tokens)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

func marshalUserinfo(w http.ResponseWriter, r *http.Request, tokens *oidc.Tokens[*oidc.IDTokenClaims], state string, rp rp.RelyingParty, info *oidc.UserInfo) {
	fmt.Println("access token", tokens.AccessToken)
	fmt.Println("refresh token", tokens.RefreshToken)
	fmt.Println("id token", tokens.IDToken)

	data, err := json.Marshal(info)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(data)
}

func requestTokenExchange(w http.ResponseWriter, r *http.Request, tokens *oidc.Tokens[*oidc.IDTokenClaims], state string, rp rp.RelyingParty, info oidc.UserInfo) {
	data := make(url.Values)
	data.Set("grant_type", string(oidc.GrantTypeTokenExchange))
	data.Set("requested_token_type", string(oidc.IDTokenType))
	data.Set("subject_token", tokens.RefreshToken)
	data.Set("subject_token_type", string(oidc.RefreshTokenType))
	data.Add("scope", "profile custom_scope:impersonate:id2")

	client := &http.Client{}
	r2, _ := http.NewRequest(http.MethodPost, authConfig.Issuer+"/oauth/token", strings.NewReader(data.Encode()))
	// r2.Header.Add("Authorization", "Basic "+"d2ViOnNlY3JldA==")
	r2.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r2.SetBasicAuth("web", "secret")

	resp, _ := client.Do(r2)
	fmt.Println(resp.Status)

	b, _ := io.ReadAll(resp.Body)
	resp.Body.Close()

	w.Write(b)
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
