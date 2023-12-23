package clients

import (
	database_types "github.com/RouteHub-Link/routehub-service-graphql/database/types"
	"github.com/gocolly/colly/v2"
)

type CollyClient struct{}

func (CollyClient) VisitScrapeOG(url string) (scrapedResult *database_types.OpenGraph) {
	c := colly.NewCollector()

	c.OnHTML("head", func(e *colly.HTMLElement) {

		scrapedResult = &database_types.OpenGraph{
			Title:          e.ChildAttr("meta[property='og:title']", "content"),
			Description:    e.ChildAttr("meta[property='og:description']", "content"),
			FavIcon:        e.ChildAttr("link[rel='shortcut icon']", "href"),
			Image:          e.ChildAttr("meta[property='og:image']", "content"),
			AlternateImage: e.ChildAttr("meta[name='twitter:image']", "content"),
			URL:            e.ChildAttr("meta[property='og:url']", "content"),
			SiteName:       e.ChildAttr("meta[property='og:site_name']", "content"),
			Type:           e.ChildAttr("meta[property='og:type']", "content"),
			Locale:         e.ChildAttr("meta[property='og:locale']", "content"),
			X: &database_types.OpenGraphX{
				Card:        e.ChildAttr("meta[name='twitter:card']", "content"),
				Site:        e.ChildAttr("meta[name='twitter:site']", "content"),
				Title:       e.ChildAttr("meta[name='twitter:title']", "content"),
				Description: e.ChildAttr("meta[name='twitter:description']", "content"),
				Image:       e.ChildAttr("meta[name='twitter:image']", "content"),
				URL:         e.ChildAttr("meta[name='twitter:url']", "content"),
				Creator:     e.ChildAttr("meta[name='twitter:creator']", "content"),
			},
		}
	})

	c.Visit(url)
	c.Wait()

	return
}

type ScrapeResult struct {
	OG *database_types.OpenGraph
}
