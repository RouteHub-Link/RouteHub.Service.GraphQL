package database_enums

import (
	"fmt"
	"io"
	"strconv"
)

type InvitationStatus string

const (
	InvitationStatusPending  InvitationStatus = "PENDING"
	InvitationStatusAccepted InvitationStatus = "ACCEPTED"
	InvitationStatusRejected InvitationStatus = "REJECTED"
)

var AllInvitationStatus = []InvitationStatus{
	InvitationStatusPending,
	InvitationStatusAccepted,
	InvitationStatusRejected,
}

func (e InvitationStatus) IsValid() bool {
	switch e {
	case InvitationStatusPending, InvitationStatusAccepted, InvitationStatusRejected:
		return true
	}
	return false
}

func (e InvitationStatus) String() string {
	return string(e)
}

func (e *InvitationStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = InvitationStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid InvitationStatus", str)
	}
	return nil
}

func (e InvitationStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
