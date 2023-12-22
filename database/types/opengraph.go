package database_types

import "encoding/json"

type OpenGraph struct {
	Title          *string     `json:"title,omitempty"`
	Description    *string     `json:"description,omitempty"`
	FavIcon        *string     `json:"favIcon,omitempty"`
	Image          *string     `json:"image,omitempty"`
	AlternateImage *string     `json:"alternateImage,omitempty"`
	URL            *string     `json:"url,omitempty"`
	SiteName       *string     `json:"siteName,omitempty"`
	Type           *string     `json:"type,omitempty"`
	Locale         *string     `json:"locale,omitempty"`
	X              *OpenGraphX `json:"x,omitempty"`
}

type OpenGraphX struct {
	Card        *string `json:"card,omitempty"`
	Site        *string `json:"site,omitempty"`
	Title       *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
	Image       *string `json:"image,omitempty"`
	URL         *string `json:"url,omitempty"`
	Type        *string `json:"type,omitempty"`
	Creator     *string `json:"creator,omitempty"`
}

func (og *OpenGraph) ParseFromJson(jsonOpenGraph string) {
	json.Unmarshal([]byte(jsonOpenGraph), &og)
}
