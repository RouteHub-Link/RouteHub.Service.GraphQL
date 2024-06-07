package services_domain_utils

import "errors"

const (
	ValidateURL  = "validate/url"
	ValidateDNS  = "validate/dns"
	ValidateSite = "validate/site"
)

func GetRoute(route string, id string) (getRoute string, err error) {
	if id == "" {
		err = errors.New("id is required")
		return
	}

	getRoute = route + "/" + id
	return
}
