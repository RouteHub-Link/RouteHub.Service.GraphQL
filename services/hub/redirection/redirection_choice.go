package redirection

import (
	"fmt"
	"io"
	"strconv"
)

type Option string

const (
	OptionTimed      Option = "TIMED"
	OptionNotAuto    Option = "NOT_AUTO_REDIRECT"
	OptionDirectHTTP Option = "DIRECT_HTTP_REDIRECT"
	OptionConfirm    Option = "CONFIRM_REDIRECT"
	OptionCustom     Option = "CUSTOM"
)

var AllOption = []Option{
	OptionTimed,
	OptionNotAuto,
	OptionDirectHTTP,
	OptionConfirm,
	OptionCustom,
}

func (e Option) IsValid() bool {
	switch e {
	case OptionTimed, OptionNotAuto, OptionDirectHTTP, OptionConfirm, OptionCustom:
		return true
	}
	return false
}

func (e Option) String() string {
	return string(e)
}

func (e *Option) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Option(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Option", str)
	}
	return nil
}

func (e Option) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
