package database_enums

import (
	"fmt"
	"io"
	"strconv"
)

type OrganizationPermission string

const (
	OrganizationPermissionDomainCreate   OrganizationPermission = "DOMAIN_CREATE"
	OrganizationPermissionDomainUpdate   OrganizationPermission = "DOMAIN_UPDATE"
	OrganizationPermissionDomainDelete   OrganizationPermission = "DOMAIN_DELETE"
	OrganizationPermissionRead           OrganizationPermission = "ORGANIZATION_READ"
	OrganizationPermissionUpdate         OrganizationPermission = "ORGANIZATION_UPDATE"
	OrganizationPermissionDelete         OrganizationPermission = "ORGANIZATION_DELETE"
	OrganizationPermissionUserInvite     OrganizationPermission = "ORGANIZATION_USER_INVITE"
	OrganizationPermissionPlatformCreate OrganizationPermission = "PLATFORM_CREATE"
)

var AllOrganizationPermission = []OrganizationPermission{
	OrganizationPermissionDomainCreate,
	OrganizationPermissionDomainUpdate,
	OrganizationPermissionDomainDelete,
	OrganizationPermissionUpdate,
	OrganizationPermissionDelete,
	OrganizationPermissionPlatformCreate,
	OrganizationPermissionUserInvite,
}

func (e OrganizationPermission) IsValid() bool {
	switch e {
	case OrganizationPermissionDomainCreate, OrganizationPermissionDomainUpdate, OrganizationPermissionDomainDelete, OrganizationPermissionUpdate, OrganizationPermissionDelete, OrganizationPermissionPlatformCreate, OrganizationPermissionUserInvite, OrganizationPermissionRead:
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
