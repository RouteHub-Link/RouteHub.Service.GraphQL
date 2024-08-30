package link

import (
	database_enums "github.com/RouteHub-Link/routehub-service-graphql/database/enums"
	database_models "github.com/RouteHub-Link/routehub-service-graphql/database/models"
	database_types "github.com/RouteHub-Link/routehub-service-graphql/database/types"
	"github.com/RouteHub-Link/routehub-service-graphql/services/hub/redirection"
	"github.com/google/uuid"
)

type Link struct {
	ID      uuid.UUID
	Key     string
	Options redirection.Option
	Content *database_types.LinkContent
}

func (l *Link) MapFromDatabaseLink(dbLink database_models.Link) {
	l.ID = dbLink.ID
	l.Key = dbLink.Path

	switch dbLink.RedirectionChoice {
	case database_enums.RedirectionOptionsTimed:
		l.Options = redirection.OptionTimed
	case database_enums.RedirectionOptionsDirectHTTPRedirect:
		l.Options = redirection.OptionDirectHTTP
	case database_enums.RedirectionOptionsConfirmRedirect:
		l.Options = redirection.OptionConfirm
	case database_enums.RedirectionOptionsNotAutoRedirect:
		l.Options = redirection.OptionNotAuto
	case database_enums.RedirectionOptionsCustom:
		l.Options = redirection.OptionCustom
	}

	l.Content = dbLink.LinkContent
}

func (l *Link) ToDatabaseLink() database_models.Link {
	dbLink := database_models.Link{}

	dbLink.ID = l.ID
	dbLink.Path = l.Key

	switch l.Options {
	case redirection.OptionTimed:
		dbLink.RedirectionChoice = database_enums.RedirectionOptionsTimed
	case redirection.OptionDirectHTTP:
		dbLink.RedirectionChoice = database_enums.RedirectionOptionsDirectHTTPRedirect
	case redirection.OptionConfirm:
		dbLink.RedirectionChoice = database_enums.RedirectionOptionsConfirmRedirect
	case redirection.OptionNotAuto:
		dbLink.RedirectionChoice = database_enums.RedirectionOptionsNotAutoRedirect
	case redirection.OptionCustom:
		dbLink.RedirectionChoice = database_enums.RedirectionOptionsCustom
	}

	l.Content = dbLink.LinkContent

	return dbLink
}
