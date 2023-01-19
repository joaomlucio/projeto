package dtos

type UpdateUser struct {
	Name     string `validate:"omitempty,min=6,max=255" json:"name,omitempty" bson:"name,omitempty"`
	IsActive *bool   `validate:"omitempty,boolean" json:"isActive,omitempty" bson:"isActive,omitempty"`
}