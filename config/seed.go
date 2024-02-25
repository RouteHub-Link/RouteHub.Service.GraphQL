package config

type Seed struct {
	Admin        SeedAdmin        `koanf:"admin"`
	Organization SeedOrganization `koanf:"organization"`
	Domain       SeedDomain       `koanf:"domain"`
}

type SeedAdmin struct {
	Email    string `koanf:"email"`
	Password string `koanf:"password"`
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
