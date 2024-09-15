package database_types

type AccountPhone struct {
	Number   string `json:"number"`
	Verified bool   `json:"verified"`
}
