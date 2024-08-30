package layout

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
