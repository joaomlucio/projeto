package services

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	mgm "github.com/joaomlucio/projeto/api/mongo"
	"github.com/joaomlucio/projeto/api/user/models"
	"github.com/joaomlucio/projeto/api/user/models/dtos"
	mongom "github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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
	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	id, _ := result.InsertedID.(primitive.ObjectID)
	new_user, _ := FindUser(id.Hex())
	return new_user, nil
}

func UpdateUser(id string, user *dtos.UpdateUser) (*models.User, error) {
	objectID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{primitive.E{Key: "_id", Value: objectID}}
	update := bson.D{primitive.E{Key: "$set", Value: user}}
	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}
	if result.MatchedCount == 0 {
		return nil, fiber.ErrNotFound
	}
	new_user, _ := FindUser(id)
	return new_user, nil
}

func DeleteUser(id string) error {
	user, err := FindUser(id)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fiber.ErrNotFound
		}
		return err
	}
	err = collection.Delete(user)
	if err != nil {
		return err
	}
	return nil
}

func FindUsers() ([]models.User, error) {
	users := []models.User{}
	filter := bson.D{{}}	
	err := collection.SimpleFind(&users, filter)

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
	err := collection.FindByID(id, user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return user, fiber.ErrNotFound
		}
		return user, err
	}
	return user, nil
}