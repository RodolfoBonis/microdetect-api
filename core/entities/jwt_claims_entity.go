package entities

import (
	"github.com/RodolfoBonis/microdetect-api/core/types"
	"github.com/google/uuid"
)

type JWTClaim struct {
	ID             uuid.UUID              `json:"sub"`
	Verified       bool                   `json:"email_verified"`
	Name           string                 `json:"name"`
	Username       string                 `json:"preferred_username"`
	FirstName      string                 `json:"given_name"`
	FamilyName     string                 `json:"family_name"`
	Email          string                 `json:"email"`
	ResourceAccess map[string]interface{} `json:"resource_access,omitempty"`
	Roles          types.Array            `json:"roles,omitempty"`
}
