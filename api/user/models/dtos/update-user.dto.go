package dtos

type UpdateUser struct {
	Name     string `validate:"omitempty,min=6,max=255" json:"nome,omitempty" bson:"nome,omitempty"`
	IsActive *bool   `validate:"omitempty,boolean" json:"ativo,omitempty" bson:"ativo,omitempty"`
}