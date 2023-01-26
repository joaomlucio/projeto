package models

import (
	"github.com/kamva/mgm/v3"
)

type User struct {
	mgm.DefaultModel `bson:",inline"`
	Name     string `json:"name,omitempty" bson:"name,omitempty"`
	IsActive *bool   `json:"isActive,omitempty" bson:"isActive"`
}