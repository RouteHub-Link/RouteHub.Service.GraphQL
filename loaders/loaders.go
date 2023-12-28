package loaders

import (
	"context"
	"net/http"
	"sync"
	"time"

	database_models "github.com/RouteHub-Link/routehub-service-graphql/database/models"
	services_utils "github.com/RouteHub-Link/routehub-service-graphql/services/utils"
	"github.com/google/uuid"
	"github.com/graph-gophers/dataloader/v7"
)

type ctxKey string

type Loaders struct {
	User *dataloader.Loader[uuid.UUID, *database_models.User]
}

type LoaderContainer struct {
	User *userLoader
}

type cacheContainer struct {
	User *services_utils.LRUExpirableCache[uuid.UUID, dataloader.Thunk[*database_models.User]]
}

const (
	loadersKey = ctxKey("dataloaders")
)

var (
	loadersOnce sync.Once
	loaders     *LoaderContainer
	caches      *cacheContainer
)

func NewLoaders() *Loaders {
	loadersOnce.Do(func() {
		loaders = NewLoaderContainer()
		caches = NewCacheContainer()
	})

	return &Loaders{
		User: dataloader.NewBatchedLoader(loaders.User.Loader,
			dataloader.WithWait[uuid.UUID, *database_models.User](time.Millisecond),
			dataloader.WithCache[uuid.UUID, *database_models.User](caches.User),
		),
	}
}

func NewLoaderContainer() *LoaderContainer {
	return &LoaderContainer{
		User: &userLoader{},
	}
}

func NewCacheContainer() *cacheContainer {
	user := &services_utils.LRUExpirableCache[uuid.UUID, dataloader.Thunk[*database_models.User]]{}
	user.Init()

	// Simple LRU Cache
	// cache := &services_utils.LRUCache[uuid.UUID, *database_models.User]{Size: 100}

	return &cacheContainer{
		User: user,
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
