package database_enums

import (
	"fmt"
	"io"
	"strconv"
)

type TicketSender string

const (
	TicketSenderUser  TicketSender = "USER"
	TicketSenderAdmin TicketSender = "Admin"
)

var AllTicketSender = []TicketSender{
	TicketSenderUser,
	TicketSenderAdmin,
}

func (e TicketSender) IsValid() bool {
	switch e {
	case TicketSenderUser, TicketSenderAdmin:
		return true
	}
	return false
}

func (e TicketSender) String() string {
	return string(e)
}

func (e *TicketSender) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = TicketSender(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid TicketSender", str)
	}
	return nil
}

func (e TicketSender) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
