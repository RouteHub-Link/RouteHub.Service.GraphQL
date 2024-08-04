package auth

import (
	"context"
	"encoding/json"
	"log"
	"log/slog"
	"net/http"
	"os"
	"strings"
	"sync"

	applicationConfig "github.com/RouteHub-Link/routehub-service-graphql/config"
	"github.com/RouteHub-Link/routehub-service-graphql/database"
	services_user "github.com/RouteHub-Link/routehub-service-graphql/services/user"
	"github.com/zitadel/oidc/v3/pkg/oidc"
	"github.com/zitadel/zitadel-go/v3/pkg/authentication"
	openid "github.com/zitadel/zitadel-go/v3/pkg/authentication/oidc"
	"github.com/zitadel/zitadel-go/v3/pkg/zitadel"
	"golang.org/x/exp/rand"
)

var (
	authConfig          applicationConfig.AuthConfig
	onceConfigureOauth2 sync.Once

	zitadelInstance           *zitadel.Zitadel
	authN                     *authentication.Authenticator[*openid.UserInfoContext[*oidc.IDTokenClaims, *oidc.UserInfo]]
	authenticationInterceptor *authentication.Interceptor[*openid.UserInfoContext[*oidc.IDTokenClaims, *oidc.UserInfo]]
	symmetricEncryptionKey    string
	userService               *services_user.UserService
)

func PKCEMiddleware(next http.Handler) http.Handler {
	ConfigureOauth(next)

	return authenticationInterceptor.CheckAuthentication()(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if authentication.IsAuthenticated(r.Context()) {
			authCtx := authenticationInterceptor.Context(r.Context())
			userInfo := authCtx.GetUserInfo()

			user := new(UserSession)
			user.ParseFromIdTokenClaims(userInfo)

			dbUser := userService.GetOrCreateUser(userInfo)
			user.ID = dbUser.ID

			newCtx := context.WithValue(r.Context(), userCtxKey, user)
			log.Printf("user session on context \nname : %s\nuser : %+v \n userInfo: %+v ", userCtxKey, user, userInfo)

			r = r.WithContext(newCtx)
		}
		next.ServeHTTP(w, r)
	}))
}

func ConfigureOauth(h http.Handler) (err error) {
	onceConfigureOauth2.Do(func() {
		userService = &services_user.UserService{DB: database.DB}
		symmetricEncryptionKey = randomString(32)

		ctx := context.Background()
		_applicationConfig := applicationConfig.ConfigurationService{}.Get()
		authConfig = _applicationConfig.AuthConfig

		host := strings.Join([]string{"http://localhost:", _applicationConfig.GraphQL.PortAsString}, "")
		redirectURL := strings.Join([]string{host, authConfig.Callback}, "")

		if authConfig.Insecure {
			zitadelInstance = zitadel.New(authConfig.Domain, zitadel.WithInsecure(authConfig.Port))
		} else {
			zitadelInstance = zitadel.New(authConfig.Domain)
		}

		authN, err = authentication.New(ctx, zitadelInstance, symmetricEncryptionKey,
			openid.DefaultAuthentication(authConfig.ClientID, redirectURL, symmetricEncryptionKey),
		)

		if err != nil {
			slog.Error("zitadel sdk could not initialize", "error", err)
			os.Exit(1)
		}

		authenticationInterceptor = authentication.Middleware(authN)

		http.Handle("/", authenticationInterceptor.CheckAuthentication()(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			if authentication.IsAuthenticated(req.Context()) {
				authCtx := authenticationInterceptor.Context(req.Context())
				data, err := json.MarshalIndent(authCtx.UserInfo, "", "	")

				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				w.Header().Set("content-type", "application/json")
				w.Write(data)
			} else {
				http.Error(w, "not authenticated", http.StatusUnauthorized)
			}
		})))

		http.Handle("/oauth2/login", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			authN.Authenticate(w, req, "")
		}))
		http.Handle(authConfig.Callback, http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			authN.Callback(w, req)
		}))
		http.Handle("/oauth2/logout", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			authN.Logout(w, req)
		}))
	})

	return
}

func randomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
