package dtos

type CreateUser struct {
	Name     string `validate:"required,min=6,max=255" json:"name" bson:"name"`
	IsActive *bool   `validate:"required,boolean" json:"isActive" bson:"isActive"`
}