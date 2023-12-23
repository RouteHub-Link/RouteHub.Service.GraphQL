package database_enums

import (
	"fmt"
	"io"
	"strconv"
)

type CrawlStatus string

const (
	CrawlStatusRequested CrawlStatus = "REQUESTED"
	CrawlStatusStarted   CrawlStatus = "STARTED"
	CrawlStatusSuccess   CrawlStatus = "SUCCESS"
	CrawlStatusFailed    CrawlStatus = "FAILED"
)

var AllCrawlStatus = []CrawlStatus{
	CrawlStatusRequested,
	CrawlStatusStarted,
	CrawlStatusSuccess,
	CrawlStatusFailed,
}

func (e CrawlStatus) IsValid() bool {
	switch e {
	case CrawlStatusRequested, CrawlStatusStarted, CrawlStatusSuccess, CrawlStatusFailed:
		return true
	}
	return false
}

func (e CrawlStatus) String() string {
	return string(e)
}

func (e *CrawlStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = CrawlStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid CrawlStatus", str)
	}
	return nil
}

func (e CrawlStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
