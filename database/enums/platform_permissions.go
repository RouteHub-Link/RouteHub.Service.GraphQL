package database_enums

import (
	"fmt"
	"io"
	"strconv"
)

type PlatformPermission string

const (
	PlatformPermissionLinkCreate PlatformPermission = "LINK_CREATE"
	PlatformPermissionLinkUpdate PlatformPermission = "LINK_UPDATE"
	PlatformPermissionLinkDelete PlatformPermission = "LINK_DELETE"
	PlatformPermissionLinkRead   PlatformPermission = "LINK_READ"
	PlatformPermissionRead       PlatformPermission = "PLATFORM_READ"
	PlatformPermissionUpdate     PlatformPermission = "PLATFORM_UPDATE"
	PlatformPermissionDelete     PlatformPermission = "PLATFORM_DELETE"
)

var AllPlatformPermission = []PlatformPermission{
	PlatformPermissionLinkCreate,
	PlatformPermissionLinkUpdate,
	PlatformPermissionLinkDelete,
	PlatformPermissionLinkRead,
	PlatformPermissionUpdate,
	PlatformPermissionDelete,
}

func (e PlatformPermission) IsValid() bool {
	switch e {
	case PlatformPermissionLinkCreate, PlatformPermissionLinkUpdate, PlatformPermissionLinkDelete, PlatformPermissionLinkRead, PlatformPermissionRead, PlatformPermissionUpdate, PlatformPermissionDelete:
		return true
	}
	return false
}

func (e PlatformPermission) String() string {
	return string(e)
}

func (e *PlatformPermission) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = PlatformPermission(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid PlatformPermission", str)
	}
	return nil
}

func (e PlatformPermission) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
