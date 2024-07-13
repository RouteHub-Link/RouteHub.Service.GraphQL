package policies

import (
	"log"

	database_enums "github.com/RouteHub-Link/routehub-service-graphql/database/enums"
	"github.com/casbin/casbin/v2"
	"github.com/google/uuid"
)

type PolicyBuilder struct {
	e       *casbin.Enforcer
	sub     uuid.UUID
	enforce bool
	eft     string
}

func NewPolicyBuilder(e *casbin.Enforcer, sub uuid.UUID, obj uuid.UUID, _eft string) *PolicyBuilder {
	return &PolicyBuilder{
		e:       e,
		sub:     sub,
		enforce: false,
		eft:     _eft,
	}
}

func (pb *PolicyBuilder) OrganizationRead(organizationId uuid.UUID) *PolicyBuilder {
	pb.AddPolicy(pb.sub.String(), organizationId.String(), database_enums.OrganizationPermissionRead.String())
	return pb
}

func (pb *PolicyBuilder) OrganizationUpdate(organizationId uuid.UUID) *PolicyBuilder {
	pb.AddPolicy(pb.sub.String(), organizationId.String(), database_enums.OrganizationPermissionUpdate.String())
	return pb
}

func (pb *PolicyBuilder) OrganizationDelete(organizationId uuid.UUID) *PolicyBuilder {
	pb.AddPolicy(pb.sub.String(), organizationId.String(), database_enums.OrganizationPermissionDelete.String())
	return pb
}

func (pb *PolicyBuilder) OrganizationPlatformCreate(organizationId uuid.UUID) *PolicyBuilder {
	pb.AddPolicy(pb.sub.String(), organizationId.String(), database_enums.OrganizationPermissionPlatformCreate.String())
	return pb
}

func (pb *PolicyBuilder) OrganizationUserInvite(organizationId uuid.UUID) *PolicyBuilder {
	pb.AddPolicy(pb.sub.String(), organizationId.String(), database_enums.OrganizationPermissionUserInvite.String())
	return pb
}

func (pb *PolicyBuilder) PlatformRead(platformId uuid.UUID) *PolicyBuilder {
	pb.AddPolicy(pb.sub.String(), platformId.String(), database_enums.PlatformPermissionRead.String())
	return pb
}

func (pb *PolicyBuilder) PlatformUpdate(platformId uuid.UUID) *PolicyBuilder {
	pb.AddPolicy(pb.sub.String(), platformId.String(), database_enums.PlatformPermissionUpdate.String())
	return pb
}

func (pb *PolicyBuilder) LinkCreate(platformId uuid.UUID) *PolicyBuilder {
	pb.AddPolicy(pb.sub.String(), platformId.String(), database_enums.PlatformPermissionLinkCreate.String())
	return pb
}

func (pb *PolicyBuilder) LinkRead(platformId uuid.UUID) *PolicyBuilder {
	pb.AddPolicy(pb.sub.String(), platformId.String(), database_enums.PlatformPermissionLinkRead.String())
	return pb
}

func (pb *PolicyBuilder) LinkDelete(platformId uuid.UUID) *PolicyBuilder {
	pb.AddPolicy(pb.sub.String(), platformId.String(), database_enums.PlatformPermissionLinkDelete.String())
	return pb
}

func (pb *PolicyBuilder) LinkUpdate(platformId uuid.UUID) *PolicyBuilder {
	pb.AddPolicy(pb.sub.String(), platformId.String(), database_enums.PlatformPermissionLinkUpdate.String())
	return pb
}

func (pb *PolicyBuilder) Build() {
	pb.e.SavePolicy()
}

func (pb *PolicyBuilder) AddPolicy(sub string, obj string, act string) {
	pb.e.AddPolicy(sub, obj, act, pb.eft)

	if pb.enforce {
		res, _ := pb.e.Enforce(sub, obj, act)
		log.Printf("Enforce policy; \nsub: %v \nobj: %v \nact: %v \neft: %v \nresult: %v", sub, obj, act, pb.eft, res)
	}
}

func (pb *PolicyBuilder) EnforceWhenAdded(enforce bool) *PolicyBuilder {
	pb.enforce = enforce
	return pb
}
