package config

type Seed struct {
	Admins       *[]SeedAdmin     `koanf:"admins"`
	Second       SeedAdmin        `koanf:"second"`
	Organization SeedOrganization `koanf:"organization"`
	Domain       SeedDomain       `koanf:"domain"`
}
type SeedAdmin struct {
	Subject string `koanf:"subject"`
}

type SeedOrganization struct {
	Name        string `koanf:"name"`
	Description string `koanf:"description"`
	Url         string `koanf:"url"`
}

type SeedDomain struct {
	Name        string `koanf:"name"`
	Description string `koanf:"description"`
	Url         string `koanf:"url"`
}
