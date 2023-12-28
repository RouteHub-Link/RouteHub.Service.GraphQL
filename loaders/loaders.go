package loaders

import (
	"context"
	"net/http"

	database_models "github.com/RouteHub-Link/routehub-service-graphql/database/models"
	services_utils "github.com/RouteHub-Link/routehub-service-graphql/services/utils"
	"github.com/google/uuid"
	"github.com/graph-gophers/dataloader/v7"
)

type Loaders struct {
	User *dataloader.Loader[uuid.UUID, *database_models.User]
}

type LoaderContainer struct {
	User *userLoader
}

type ctxKey string

const (
	loadersKey = ctxKey("dataloaders")
)

func NewLoaders() *Loaders {
	loaders := NewLoaderContainer()

	cache := &services_utils.LRUCache[uuid.UUID, *database_models.User]{Size: 100}

	return &Loaders{
		User: dataloader.NewBatchedLoader(loaders.User.Loader,
			dataloader.WithCache[uuid.UUID, *database_models.User](cache),
		),
	}
}

func NewLoaderContainer() *LoaderContainer {
	return &LoaderContainer{
		User: &userLoader{},
	}
}

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		loader := NewLoaders()
		r = r.WithContext(context.WithValue(r.Context(), loadersKey, loader))
		next.ServeHTTP(w, r)
	})
}

// For returns the dataloader for a given context
func For(ctx context.Context) *Loaders {
	return ctx.Value(loadersKey).(*Loaders)
}

// handleError creates array of result with the same error repeated for as many items requested
func handleError[T any](itemsLength int, err error) []*dataloader.Result[T] {
	result := make([]*dataloader.Result[T], itemsLength)
	for i := 0; i < itemsLength; i++ {
		result[i] = &dataloader.Result[T]{Error: err}
	}
	return result
}
