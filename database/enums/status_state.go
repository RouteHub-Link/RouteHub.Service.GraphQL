package database_enums

import (
	"fmt"
	"io"
	"strconv"
)

type StatusState string

const (
	StatusStatePasive StatusState = "PASIVE"
	StatusStateActive StatusState = "ACTIVE"
)

var AllStatusState = []StatusState{
	StatusStatePasive,
	StatusStateActive,
}

func (e StatusState) IsValid() bool {
	switch e {
	case StatusStatePasive, StatusStateActive:
		return true
	}
	return false
}

func (e StatusState) String() string {
	return string(e)
}

func (e *StatusState) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = StatusState(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid StatusState", str)
	}
	return nil
}

func (e StatusState) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
