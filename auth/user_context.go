package auth

import (
	"log"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/zitadel/oidc/v3/pkg/oidc"
)

// to prevent collisions between different context uses
var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

// A stand-in for our database backed user object
type UserSession struct {
	ID       uuid.UUID
	UserInfo *oidc.UserInfo
}

func (u *UserSession) ParseFromClaims(claims jwt.MapClaims) {
	log.Printf("claims:%v", claims)
	u.ID = uuid.MustParse(claims["jti"].(string))
}

func (u *UserSession) ParseFromIdTokenClaims(userInfo *oidc.UserInfo) {
	u.UserInfo = userInfo
}

/*
func (u *UserSession) ToClaims() *jwt.MapClaims {
	return &jwt.MapClaims{
		"jti":      u.ID.String(),
		"username": u.Name,
		"exp":      jwt.TimeFunc().Add(time.Hour * 24 * 365).Unix(),
	}
}
*/
