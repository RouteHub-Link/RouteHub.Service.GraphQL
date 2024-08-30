package status

import (
	"fmt"
	"io"
	"strconv"
)

type State string

const (
	StatePasive State = "PASIVE"
	StateActive State = "ACTIVE"
)

var AllState = []State{
	StatePasive,
	StateActive,
}

func (e State) IsValid() bool {
	switch e {
	case StatePasive, StateActive:
		return true
	}
	return false
}

func (e State) String() string {
	return string(e)
}

func (e *State) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = State(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid State", str)
	}
	return nil
}

func (e State) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
