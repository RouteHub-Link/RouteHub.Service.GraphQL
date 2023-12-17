package database_enums

import (
	"fmt"
	"io"
	"strconv"
)

type DNSStatus string

const (
	DNSStatusPending  DNSStatus = "PENDING"
	DNSStatusVerified DNSStatus = "VERIFIED"
	DNSStatusFailed   DNSStatus = "FAILED"
)

var AllDNSStatus = []DNSStatus{
	DNSStatusPending,
	DNSStatusVerified,
	DNSStatusFailed,
}

func (e DNSStatus) IsValid() bool {
	switch e {
	case DNSStatusPending, DNSStatusVerified, DNSStatusFailed:
		return true
	}
	return false
}

func (e DNSStatus) String() string {
	return string(e)
}

func (e *DNSStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = DNSStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid DNSStatus", str)
	}
	return nil
}

func (e DNSStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
