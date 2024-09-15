package database_types

type PlatformDescription struct {
	MetaDescription   MetaDescription
	NavbarDescription NavbarDescription
	FooterDescription FooterDescription
}

type NavbarDescription struct {
	BrandImg   *ImageDescription
	BrandURL   string
	BrandName  string
	Target     string
	StartItems *[]NavbarItem
	EndButtons *[]NavbarButton
}

type NavbarItem struct {
	Text     string
	URL      string
	Target   string
	Icon     string
	Dropdown *[]NavbarItem
}

type NavbarButton struct {
	Text       string
	URL        string
	Icon       string
	Target     string
	ColorClass string
}

type ImageDescription struct {
	SRC    string
	Alt    string
	Height string
	Width  string
}

type FooterDescription struct {
	ShowRouteHubBranding bool
	CompanyBrandingHtml  string
	SocialMediaContainer *SocialMediaContainer
}

type SocialMediaContainer struct {
	SocialMediaPeddingClass string
	SocialMediaColorClass   string
	SocialMediaSizeClass    string
	SocialMediaLinks        *[]ASocialMedia
}

type ASocialMedia struct {
	Icon   string
	Link   string
	Target string
}
