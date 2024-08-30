package database_types

import "encoding/json"

type LinkContent struct {
	Title              string
	Subtitle           string
	ContentContainer   string
	RedirectionURLText string
	RedirectionDelay   *int
	MetaDescription    *MetaDescription
	AdditionalHead     *string
	AdditionalFooter   *string
}

type MetaDescription struct {
	Title         string
	FavIcon       string
	Description   string
	Locale        string
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
	OGCard        string
	OGSite        string
	OGType        string
	OGCreator     string
}

func (og *MetaDescription) ParseFromJson(jsonOpenGraph string) {
	json.Unmarshal([]byte(jsonOpenGraph), &og)
}

func (og *MetaDescription) AsJson() (string, error) {
	jsonOpengraph, err := json.Marshal(og)
	if err != nil {
		return "", err
	}
	return string(jsonOpengraph), nil
}
