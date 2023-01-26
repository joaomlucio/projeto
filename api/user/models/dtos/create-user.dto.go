package dtos

type CreateUser struct {
	Name     string `validate:"required,min=6,max=255" json:"nome" bson:"nome"`
	IsActive bool   `default:"false" validate:"boolean" json:"ativo" bson:"ativo"`
}