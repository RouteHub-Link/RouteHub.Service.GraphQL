package link

import (
	database_enums "github.com/RouteHub-Link/routehub-service-graphql/database/enums"
	database_models "github.com/RouteHub-Link/routehub-service-graphql/database/models"
	"github.com/RouteHub-Link/routehub-service-graphql/services/hub/layout"
	"github.com/RouteHub-Link/routehub-service-graphql/services/hub/redirection"
	"github.com/google/uuid"
)

type Link struct {
	ID      uuid.UUID
	Key     string
	Options redirection.Option
	Content *LinkContent
}

type LinkContent struct {
	PageTitle          string
	Title              string
	Description        string
	RedirectionDetails string
	RedirectionURL     string
	RedirectionURLText string
	RedirectionDelay   string
	HTML               string
	MetaDescription    *layout.MetaDescription
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

	l.MapFromLinkOG(dbLink)
}

func (l *Link) MapFromLinkOG(dbLink database_models.Link) {
	/*

		lcs.RedirectionURL = dbLink.Target
		lcs.RedirectionURLText = dbLink.TargetText
		lcs.MetaDescription = &layout.MetaDescription{
			Title: dbLink.PageTitle,
		}

		if dbLink.OpenGraph == nil {
			l.Content = &lcs
			return
		}

		lcs.PageTitle = dbLink.OpenGraph.Title
		lcs.Title = dbLink.OpenGraph.Title
		lcs.Description = dbLink.OpenGraph.Description
		//lcs.RedirectionDetails = dbLink.OpenGraph.RedirectionDetails
		lcs.RedirectionDetails = dbLink.TargetDetails
		if dbLink.RedirectionChoice == database_enums.RedirectionOptionsCustom {
			lcs.HTML = dbLink.TargetDetails
		}

					type MetaDescription struct {
					Title         string
					FavIcon       string
					Description   string
					OGTitle       string
					OGDescription string
					OGURL         string
					OGSiteName    string
					OGMetaType    string
					OGLocale      string
					OGBigImage    string
					OGBigWidth    string
					OGBigHeight   string
					OGSmallImage  string
					OGSmallWidth  string
					OGSmallHeight string
				}


			lcs.MetaDescription = &layout.MetaDescription{
				Title:         dbLink.OpenGraph.Title,
				Description:   dbLink.OpenGraph.Description,
				FavIcon:       dbLink.OpenGraph.FavIcon,
				OGTitle:       dbLink.OpenGraph.Title,
				OGCard:        dbLink.OpenGraph.X.Card,
				OGSite:        dbLink.OpenGraph.X.Site,
				OGType:        dbLink.OpenGraph.X.Type,
				OGCreator:     dbLink.OpenGraph.X.Creator,
				OGDescription: dbLink.OpenGraph.Description,
				OGURL:         dbLink.OpenGraph.URL,
				OGSiteName:    dbLink.OpenGraph.SiteName,
				OGMetaType:    dbLink.OpenGraph.Type,
				OGLocale:      dbLink.OpenGraph.Locale,
				OGBigImage:    dbLink.OpenGraph.Image,
				OGBigWidth:    "",
				OGBigHeight:   "",
				OGSmallImage:  dbLink.OpenGraph.AlternateImage,
				OGSmallWidth:  "",
				OGSmallHeight: "",
			}
	*/

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
	return dbLink
	/*
		dbLink.Target = l.Content.RedirectionURL
		dbLink.TargetText = l.Content.RedirectionURLText
		dbLink.PageTitle = l.Content.PageTitle
		dbLink.TargetDetails = l.Content.RedirectionDetails

		if l.Content.MetaDescription != nil {
			dbLink.OpenGraph = &database_types.MetaDescription{}
		}

		return dbLink
	*/
}
