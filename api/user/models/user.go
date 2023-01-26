package models

import (
	"github.com/kamva/mgm/v3"
)

type User struct {
	mgm.DefaultModel `bson:",inline"`
	Name     string `json:"nome,omitempty" bson:"nome,omitempty"`
	IsActive *bool   `json:"ativo,omitempty" bson:"ativo"`
}