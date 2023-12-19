package database_enums

import (
	"fmt"
	"io"
	"strconv"
)

type OrganizationPermission string

const (
	OrganizationPermissionDomainCreate       OrganizationPermission = "DOMAIN_CREATE"
	OrganizationPermissionDomainUpdate       OrganizationPermission = "DOMAIN_UPDATE"
	OrganizationPermissionDomainDelete       OrganizationPermission = "DOMAIN_DELETE"
	OrganizationPermissionOrganizationUpdate OrganizationPermission = "Organization_UPDATE"
	OrganizationPermissionOrganizationDelete OrganizationPermission = "Organization_DELETE"
	OrganizationPermissionPlatformCreate     OrganizationPermission = "PLATFORM_CREATE"
	OrganizationPermissionPlatformUpdate     OrganizationPermission = "PLATFORM_UPDATE"
	OrganizationPermissionPlatformDelete     OrganizationPermission = "PLATFORM_DELETE"
	OrganizationPermissionUserInvite         OrganizationPermission = "USER_INVITE"
	OrganizationPermissionPlatformUserAdd    OrganizationPermission = "PLATFORM_USER_ADD"
	OrganizationPermissionPlatformUserRemove OrganizationPermission = "PLATFORM_USER_REMOVE"
	OrganizationPermissionPlatformUserUpdate OrganizationPermission = "PLATFORM_USER_UPDATE"
)

var AllOrganizationPermission = []OrganizationPermission{
	OrganizationPermissionDomainCreate,
	OrganizationPermissionDomainUpdate,
	OrganizationPermissionDomainDelete,
	OrganizationPermissionOrganizationUpdate,
	OrganizationPermissionOrganizationDelete,
	OrganizationPermissionPlatformCreate,
	OrganizationPermissionPlatformUpdate,
	OrganizationPermissionPlatformDelete,
	OrganizationPermissionUserInvite,
	OrganizationPermissionPlatformUserAdd,
	OrganizationPermissionPlatformUserRemove,
	OrganizationPermissionPlatformUserUpdate,
}

func (e OrganizationPermission) IsValid() bool {
	switch e {
	case OrganizationPermissionDomainCreate, OrganizationPermissionDomainUpdate, OrganizationPermissionDomainDelete, OrganizationPermissionOrganizationUpdate, OrganizationPermissionOrganizationDelete, OrganizationPermissionPlatformCreate, OrganizationPermissionPlatformUpdate, OrganizationPermissionPlatformDelete, OrganizationPermissionUserInvite, OrganizationPermissionPlatformUserAdd, OrganizationPermissionPlatformUserRemove, OrganizationPermissionPlatformUserUpdate:
		return true
	}
	return false
}

func (e OrganizationPermission) String() string {
	return string(e)
}

func (e *OrganizationPermission) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = OrganizationPermission(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid OrganizationPermission", str)
	}
	return nil
}

func (e OrganizationPermission) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
