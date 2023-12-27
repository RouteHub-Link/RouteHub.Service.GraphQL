package loaders

import (
	"context"
	"log"

	"github.com/RouteHub-Link/routehub-service-graphql/database"
	database_models "github.com/RouteHub-Link/routehub-service-graphql/database/models"
	services_user "github.com/RouteHub-Link/routehub-service-graphql/services/user"
	"github.com/google/uuid"
	"github.com/graph-gophers/dataloader/v7"
)

type userLoader struct{}

func (*userLoader) Loader(ctx context.Context, userIds []uuid.UUID) []*dataloader.Result[*database_models.User] {
	result := make([]*dataloader.Result[*database_models.User], len(userIds))

	log.Printf("userLoader.Loader: %v", userIds)

	userService := services_user.UserService{DB: database.DB}

	users, _ := userService.UsersByIds(userIds)

	userMap := make(map[uuid.UUID]*database_models.User)
	for _, user := range users {
		userMap[user.ID] = user
	}

	for i, userId := range userIds {
		result[i] = &dataloader.Result[*database_models.User]{Data: userMap[userId]}
	}

	return result
}

func (*userLoader) Get(ctx context.Context, userId uuid.UUID) (*database_models.User, error) {
	loaders := For(ctx)
	return loaders.User.Load(ctx, userId)()
}

func (*userLoader) GetBatch(ctx context.Context, userIds []uuid.UUID) (data []*database_models.User, err error) {
	loaders := For(ctx)
	data, errs := loaders.User.LoadMany(ctx, userIds)()

	if len(errs) > 0 {
		return nil, errs[0]
	}

	return
}
