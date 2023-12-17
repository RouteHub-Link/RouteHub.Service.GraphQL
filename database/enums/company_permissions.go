package database_enums

import (
	"fmt"
	"io"
	"strconv"
)

type CompanyPermission string

const (
	CompanyPermissionDomainCreate       CompanyPermission = "DOMAIN_CREATE"
	CompanyPermissionDomainUpdate       CompanyPermission = "DOMAIN_UPDATE"
	CompanyPermissionDomainDelete       CompanyPermission = "DOMAIN_DELETE"
	CompanyPermissionCompanyUpdate      CompanyPermission = "COMPANY_UPDATE"
	CompanyPermissionCompanyDelete      CompanyPermission = "COMPANY_DELETE"
	CompanyPermissionPlatformCreate     CompanyPermission = "PLATFORM_CREATE"
	CompanyPermissionPlatformUpdate     CompanyPermission = "PLATFORM_UPDATE"
	CompanyPermissionPlatformDelete     CompanyPermission = "PLATFORM_DELETE"
	CompanyPermissionUserInvite         CompanyPermission = "USER_INVITE"
	CompanyPermissionPlatformUserAdd    CompanyPermission = "PLATFORM_USER_ADD"
	CompanyPermissionPlatformUserRemove CompanyPermission = "PLATFORM_USER_REMOVE"
	CompanyPermissionPlatformUserUpdate CompanyPermission = "PLATFORM_USER_UPDATE"
)

var AllCompanyPermission = []CompanyPermission{
	CompanyPermissionDomainCreate,
	CompanyPermissionDomainUpdate,
	CompanyPermissionDomainDelete,
	CompanyPermissionCompanyUpdate,
	CompanyPermissionCompanyDelete,
	CompanyPermissionPlatformCreate,
	CompanyPermissionPlatformUpdate,
	CompanyPermissionPlatformDelete,
	CompanyPermissionUserInvite,
	CompanyPermissionPlatformUserAdd,
	CompanyPermissionPlatformUserRemove,
	CompanyPermissionPlatformUserUpdate,
}

func (e CompanyPermission) IsValid() bool {
	switch e {
	case CompanyPermissionDomainCreate, CompanyPermissionDomainUpdate, CompanyPermissionDomainDelete, CompanyPermissionCompanyUpdate, CompanyPermissionCompanyDelete, CompanyPermissionPlatformCreate, CompanyPermissionPlatformUpdate, CompanyPermissionPlatformDelete, CompanyPermissionUserInvite, CompanyPermissionPlatformUserAdd, CompanyPermissionPlatformUserRemove, CompanyPermissionPlatformUserUpdate:
		return true
	}
	return false
}

func (e CompanyPermission) String() string {
	return string(e)
}

func (e *CompanyPermission) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = CompanyPermission(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid CompanyPermission", str)
	}
	return nil
}

func (e CompanyPermission) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
