package services

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

	mongom "github.com/kamva/mgm/v3"
	mgm "github.com/joaomlucio/projeto/api/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	
	"github.com/joaomlucio/projeto/api/user/models"
	"github.com/joaomlucio/projeto/api/user/models/dtos"
)

var collection *mongom.Collection
var ctx context.Context
var validate *validator.Validate

func init(){
	collection = mgm.Collection
	ctx = context.Background()
	validate = validator.New()
}

func ValidateStruct(user interface{}) []*dtos.ErrorResponse {
    var errors []*dtos.ErrorResponse
    err := validate.Struct(user)
    if err != nil {
        for _, err := range err.(validator.ValidationErrors) {
            var element dtos.ErrorResponse
            element.FailedField = err.StructNamespace()
            element.Tag = err.Tag()
            element.Value = err.Param()
            errors = append(errors, &element)
        }
    }
    return errors
}

func CreateUser(user *dtos.CreateUser) (*models.User, error) {
	u := &models.User{Name: user.Name, IsActive: &user.IsActive}
	err := collection.CreateWithCtx(ctx, u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func UpdateUser(id string, user *dtos.UpdateUser) (*models.User, error) {
	u := new(models.User)
	err := collection.FindByIDWithCtx(ctx, id, u)
	if err != nil {
		return nil, err
	}
	u.Name = user.Name
	u.IsActive = user.IsActive
	err = collection.UpdateWithCtx(ctx, u)
	if err != nil {
		return nil, fiber.ErrNotFound
	}
	return u, nil
}

func DeleteUser(id string) error {
	user, err := FindUser(id)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fiber.ErrNotFound
		}
		return err
	}
	err = collection.DeleteWithCtx(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func FindUsers() ([]models.User, error) {
	users := []models.User{}
	filter := bson.D{{}}	
	err := collection.SimpleFindWithCtx(ctx, &users, filter)

	if err != nil {
		return users, err
	}

	if len(users) == 0 {
		return users, fiber.ErrNotFound
	}

	return users, nil
}

func FindUser(id string) (*models.User, error) {
	user := new(models.User)
	err := collection.FindByIDWithCtx(ctx, id, user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return user, fiber.ErrNotFound
		}
		return user, err
	}
	return user, nil
}