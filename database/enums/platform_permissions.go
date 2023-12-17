package database_enums

import (
	"fmt"
	"io"
	"strconv"
)

type PlatformPermission string

const (
	PlatformPermissionLinkCreate     PlatformPermission = "LINK_CREATE"
	PlatformPermissionLinkUpdate     PlatformPermission = "LINK_UPDATE"
	PlatformPermissionLinkDelete     PlatformPermission = "LINK_DELETE"
	PlatformPermissionPlatformCreate PlatformPermission = "PLATFORM_CREATE"
	PlatformPermissionPlatformUpdate PlatformPermission = "PLATFORM_UPDATE"
	PlatformPermissionPlatformDelete PlatformPermission = "PLATFORM_DELETE"
)

var AllPlatformPermission = []PlatformPermission{
	PlatformPermissionLinkCreate,
	PlatformPermissionLinkUpdate,
	PlatformPermissionLinkDelete,
	PlatformPermissionPlatformCreate,
	PlatformPermissionPlatformUpdate,
	PlatformPermissionPlatformDelete,
}

func (e PlatformPermission) IsValid() bool {
	switch e {
	case PlatformPermissionLinkCreate, PlatformPermissionLinkUpdate, PlatformPermissionLinkDelete, PlatformPermissionPlatformCreate, PlatformPermissionPlatformUpdate, PlatformPermissionPlatformDelete:
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
