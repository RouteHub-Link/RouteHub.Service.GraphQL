package database_enums

import (
	"fmt"
	"io"
	"strconv"
)

type TicketStatus string

const (
	TicketStatusOpen   TicketStatus = "OPEN"
	TicketStatusClosed TicketStatus = "CLOSED"
)

var AllTicketStatus = []TicketStatus{
	TicketStatusOpen,
	TicketStatusClosed,
}

func (e TicketStatus) IsValid() bool {
	switch e {
	case TicketStatusOpen, TicketStatusClosed:
		return true
	}
	return false
}

func (e TicketStatus) String() string {
	return string(e)
}

func (e *TicketStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = TicketStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid TicketStatus", str)
	}
	return nil
}

func (e TicketStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
