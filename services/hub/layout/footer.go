package layout

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
