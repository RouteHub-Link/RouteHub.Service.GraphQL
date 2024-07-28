package auth

import (
	"context"
	"log"
)

func ForContext(ctx context.Context) *UserSession {
	raw, err := ctx.Value(userCtxKey).(*UserSession)
	if !err {
		log.Printf("session not found from context %s", userCtxKey)
	}

	return raw
}
