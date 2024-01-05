package loaders

import (
	"context"
	"log"
	"net/http"
	"sync"

	"github.com/RouteHub-Link/routehub-service-graphql/config"
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
	_config     = config.ConfigurationService{}.Get()
	cacheSize   = _config.GraphQL.Dataloader.Lrue.Size
	cacheExpire = _config.GraphQL.Dataloader.Lrue.Expire
)

func NewLoaders() *Loaders {
	loadersOnce.Do(func() {
		loaders = NewLoaderContainer()
		caches = NewCacheContainer()
	})

	userOptions := []dataloader.Option[uuid.UUID, *database_models.User]{dataloader.WithWait[uuid.UUID, *database_models.User](_config.GraphQL.Dataloader.Wait)}

	if _config.GraphQL.Dataloader.Cache {
		userOptions = append(userOptions, dataloader.WithCache[uuid.UUID, *database_models.User](caches.User))
	} else {
		log.Printf("LRU Timed Cache is disabled")
	}

	loaders := &Loaders{
		User: dataloader.NewBatchedLoader(loaders.User.Loader, userOptions...),
	}

	return loaders
}

func NewLoaderContainer() *LoaderContainer {
	return &LoaderContainer{
		User: &userLoader{},
	}
}

func NewCacheContainer() *cacheContainer {
	user := &services_utils.LRUExpirableCache[uuid.UUID, dataloader.Thunk[*database_models.User]]{}

	user.Expire = cacheExpire
	user.Size = cacheSize
	user.Init()

	// Simple LRU Cache exmaple
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
