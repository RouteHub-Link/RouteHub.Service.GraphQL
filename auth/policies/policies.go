package policies

import (
	"log"

	"github.com/casbin/casbin/v2"
	"github.com/google/uuid"
)

type GeneratedPolicies struct {
	Policies      [][]string
	GroupedPolicy [][]string
}

func NewOrganizationPolicies(userId uuid.UUID, org uuid.UUID, platformId uuid.UUID) GeneratedPolicies {
	policies := [][]string{
		{"admin", org.String(), platformId.String(), "read"},
		{"admin", org.String(), platformId.String(), "write"},
		{"admin", org.String(), platformId.String(), "delete"},
		{"admin", org.String(), "invitation", "create"},
		{"admin", org.String(), "invitation", "read"},
		{"admin", org.String(), "invitation", "delete"},
	}

	groupedPolicy := [][]string{{userId.String(), "admin", org.String()}}

	return GeneratedPolicies{Policies: policies, GroupedPolicy: groupedPolicy}
}

type PolicyBuilder struct {
	e              *casbin.Enforcer
	userId         uuid.UUID
	organizationId uuid.UUID
	platformId     uuid.UUID
}

func NewPolicyBuilder(e *casbin.Enforcer, userId, organizationId, platformId uuid.UUID) *PolicyBuilder {
	return &PolicyBuilder{
		e:              e,
		userId:         userId,
		organizationId: organizationId,
		platformId:     platformId,
	}
}

func (pb *PolicyBuilder) OrganizationRead() *PolicyBuilder {
	pb.e.AddPolicy("admin", pb.organizationId.String(), pb.organizationId.String(), "read")
	return pb
}

func (pb *PolicyBuilder) OrganizationUpdate() *PolicyBuilder {
	pb.e.AddPolicy("admin", pb.organizationId.String(), pb.organizationId.String(), "update")
	return pb
}

func (pb *PolicyBuilder) OrganizationDelete() *PolicyBuilder {
	pb.e.AddPolicy("admin", pb.organizationId.String(), pb.organizationId.String(), "delete")
	return pb
}

func (pb *PolicyBuilder) PlatformRead() *PolicyBuilder {
	pb.e.AddPolicy("admin", pb.organizationId.String(), pb.platformId.String(), "read")
	return pb
}

func (pb *PolicyBuilder) PlatformUpdate() *PolicyBuilder {
	pb.e.AddPolicy("admin", pb.organizationId.String(), pb.platformId.String(), "update")
	return pb
}

func (pb *PolicyBuilder) LinkCreate() *PolicyBuilder {
	pb.e.AddPolicy("admin", pb.organizationId.String(), pb.platformId.String(), "link_create")
	return pb
}

func (pb *PolicyBuilder) LinkRead() *PolicyBuilder {
	pb.e.AddPolicy("admin", pb.organizationId.String(), pb.platformId.String(), "link_read")
	return pb
}

func (pb *PolicyBuilder) LinkDelete() *PolicyBuilder {
	pb.e.AddPolicy("admin", pb.organizationId.String(), pb.platformId.String(), "link_delete")
	return pb
}

func (pb *PolicyBuilder) LinkUpdate() *PolicyBuilder {
	pb.e.AddPolicy("admin", pb.organizationId.String(), pb.platformId.String(), "link_update")
	return pb
}

func (pb *PolicyBuilder) InvitationCreate() *PolicyBuilder {
	pb.e.AddPolicy("admin", pb.organizationId.String(), "invitation", "create")
	return pb
}

func (pb *PolicyBuilder) InvitationRead() *PolicyBuilder {
	pb.e.AddPolicy("admin", pb.organizationId.String(), "invitation", "read")
	return pb
}

func (pb *PolicyBuilder) InvitationDelete() *PolicyBuilder {
	pb.e.AddPolicy("admin", pb.organizationId.String(), "invitation", "delete")
	return pb
}

func (pb *PolicyBuilder) GroupingPolicy() *PolicyBuilder {
	pb.e.AddGroupingPolicy(pb.userId.String(), "admin", pb.organizationId.String())
	return pb
}

func (pb *PolicyBuilder) Build() {
	pb.e.SavePolicy()
}

func (pb *PolicyBuilder) Enforce() {
	res, _ := pb.e.Enforce(pb.userId.String(), pb.organizationId.String(), pb.organizationId.String(), "read")
	log.Printf("userId, organizationId, organizationId, 'read' res : %v\n", res)

	res, _ = pb.e.Enforce(pb.userId.String(), pb.organizationId.String(), pb.organizationId.String(), "update")
	log.Printf("userId, organizationId, organizationId, 'update' res : %v\n", res)

	res, _ = pb.e.Enforce(pb.userId.String(), pb.organizationId.String(), pb.organizationId.String(), "delete")
	log.Printf("userId, organizationId, organizationId, 'delete' res : %v\n", res)

	res, _ = pb.e.Enforce(pb.userId.String(), pb.organizationId.String(), pb.platformId.String(), "read")
	log.Printf("userId, organizationId, platformId, 'read' res : %v\n", res)

	res, _ = pb.e.Enforce(pb.userId.String(), pb.organizationId.String(), pb.platformId.String(), "update")
	log.Printf("userId, organizationId, platformId, 'update' res : %v\n", res)

	res, _ = pb.e.Enforce(pb.userId.String(), pb.organizationId.String(), pb.platformId.String(), "delete")
	log.Printf("userId, organizationId, platformId, 'delete' res : %v\n", res)

	res, _ = pb.e.Enforce(pb.userId.String(), pb.organizationId.String(), pb.platformId.String(), "link_create")
	log.Printf("userId, organizationId, platformId, 'link_create' res : %v\n", res)

	res, _ = pb.e.Enforce(pb.userId.String(), pb.organizationId.String(), pb.platformId.String(), "link_read")
	log.Printf("userId, organizationId, platformId, 'link_read' res : %v\n", res)

	res, _ = pb.e.Enforce(pb.userId.String(), pb.organizationId.String(), pb.platformId.String(), "link_delete")
	log.Printf("userId, organizationId, platformId, 'link_delete' res : %v\n", res)

	res, _ = pb.e.Enforce(pb.userId.String(), pb.organizationId.String(), pb.platformId.String(), "link_update")
	log.Printf("userId, organizationId, platformId, 'link_update' res : %v\n", res)

	res, _ = pb.e.Enforce(pb.userId.String(), pb.organizationId.String(), "invitation", "create")
	log.Printf("userId, organizationId, 'invitation', 'create' res : %v\n", res)

	res, _ = pb.e.Enforce(pb.userId.String(), pb.organizationId.String(), "invitation", "read")
	log.Printf("userId, organizationId, 'invitation', 'read' res : %v\n", res)

	res, _ = pb.e.Enforce(pb.userId.String(), pb.organizationId.String(), "invitation", "delete")
	log.Printf("userId, organizationId, 'invitation', 'delete' res : %v\n", res)
}
