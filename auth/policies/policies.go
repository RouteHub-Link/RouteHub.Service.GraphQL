package policies

import (
	"log"

	"github.com/casbin/casbin/v2"
	"github.com/google/uuid"
)

type PolicyBuilder struct {
	e       *casbin.SyncedCachedEnforcer
	sub     uuid.UUID
	enforce bool
	eft     string
}

type PermissionActExplained struct {
	Permission string
	Exp        []string
	Act        string
}

func NewPolicyBuilder(e *casbin.SyncedCachedEnforcer, sub uuid.UUID, _eft string) *PolicyBuilder {
	return &PolicyBuilder{
		e:       e,
		sub:     sub,
		enforce: false,
		eft:     _eft,
	}
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

func EnforcePermissions(e *casbin.SyncedCachedEnforcer, userId uuid.UUID, platformId uuid.UUID, permissions []string) ([]PermissionActExplained, error) {
	permissionActExplained := []PermissionActExplained{}

	// BatchEnforce does some magic (for in enforce) to checking multiple permissions
	for _, permission := range permissions {
		hasPermission, exp, err := e.EnforceEx(userId.String(), platformId.String(), permission)
		log.Printf("\nEnforcePermissions;\nhasPermission: %+v\nexp: %+v\nerr: %+v\n\n", hasPermission, exp, err)
		if err != nil {
			return nil, err
		}

		if hasPermission {
			_act := exp[len(exp)-1]
			permissionActExplained = append(permissionActExplained, PermissionActExplained{Permission: permission, Exp: exp, Act: _act})
		}
	}

	return permissionActExplained, nil
}
