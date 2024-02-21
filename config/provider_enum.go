package config

import (
	"fmt"
	"io"
	"strconv"
)

type Provider string

const (
	Postgres Provider = "postgres"
	Embed    Provider = "embed"
)

var AllProviders = []Provider{
	Postgres,
	Embed,
}

func (e Provider) IsValid() bool {
	switch e {
	case Postgres, Embed:
		return true
	}
	return false
}

func (e Provider) String() string {
	return string(e)
}

func (e *Provider) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Provider(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid DNSStatus", str)
	}
	return nil
}

func (e Provider) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
