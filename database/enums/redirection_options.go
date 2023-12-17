package database_enums

import (
	"fmt"
	"io"
	"strconv"
)

type RedirectionOptions string

const (
	RedirectionOptionsTimed              RedirectionOptions = "TIMED"
	RedirectionOptionsNotAutoRedirect    RedirectionOptions = "NOT_AUTO_REDIRECT"
	RedirectionOptionsDirectHTTPRedirect RedirectionOptions = "DIRECT_HTTP_REDIRECT"
	RedirectionOptionsConfirmRedirect    RedirectionOptions = "CONFIRM_REDIRECT"
)

var AllRedirectionOptions = []RedirectionOptions{
	RedirectionOptionsTimed,
	RedirectionOptionsNotAutoRedirect,
	RedirectionOptionsDirectHTTPRedirect,
	RedirectionOptionsConfirmRedirect,
}

func (e RedirectionOptions) IsValid() bool {
	switch e {
	case RedirectionOptionsTimed, RedirectionOptionsNotAutoRedirect, RedirectionOptionsDirectHTTPRedirect, RedirectionOptionsConfirmRedirect:
		return true
	}
	return false
}

func (e RedirectionOptions) String() string {
	return string(e)
}

func (e *RedirectionOptions) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = RedirectionOptions(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid RedirectionOptions", str)
	}
	return nil
}

func (e RedirectionOptions) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
