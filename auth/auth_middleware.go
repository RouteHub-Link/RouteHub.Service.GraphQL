package auth

import (
	"context"
	"log"
	"net/http"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")

		if header == "" {
			next.ServeHTTP(w, r)
			return
		}

		getAuthToken := header[len("Bearer "):]

		tokenStr := getAuthToken
		claims, err := ParseToken(tokenStr)

		if err != nil {
			http.Error(w, "Invalid token", http.StatusForbidden)
			return
		}

		user := new(UserSession)
		user.ParseFromClaims(claims)

		ctx := context.WithValue(r.Context(), userCtxKey, user)
		log.Printf("user session on context \nname : %s\nuser : %+v ", userCtxKey, user)

		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func ForContext(ctx context.Context) *UserSession {
	raw, err := ctx.Value(userCtxKey).(*UserSession)
	if !err {
		log.Printf("session not found from context %s", userCtxKey)
	}

	return raw
}
