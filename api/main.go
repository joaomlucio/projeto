package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

	"context"
	"encoding/json"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/joaomlucio/projeto/api/user/models/user"
)



type CreateUser struct {
	Name     string `validate:"required,min=6,max=255" json:"name" bson:"name"`
	IsActive *bool   `validate:"required,boolean" json:"isActive" bson:"isActive"`
}

type UpdateUser struct {
	Name     string `validate:"omitempty,min=6,max=255" json:"name,omitempty" bson:"name,omitempty"`
	IsActive *bool   `validate:"omitempty,boolean" json:"isActive,omitempty" bson:"isActive,omitempty"`
}

type ErrorResponse struct {
    FailedField string
    Tag         string
    Value       string
}

var collection *mongo.Collection
var ctx = context.TODO()
var validate = validator.New()

func ValidateStruct(user interface{}) []*ErrorResponse {
    var errors []*ErrorResponse
    err := validate.Struct(user)
    if err != nil {
        for _, err := range err.(validator.ValidationErrors) {
            var element ErrorResponse
            element.FailedField = err.StructNamespace()
            element.Tag = err.Tag()
            element.Value = err.Param()
            errors = append(errors, &element)
        }
    }
    return errors
}

func createUser(user *CreateUser) (*mongo.InsertOneResult, error) {
	return collection.InsertOne(ctx, user)
}

func updateUser(id string, user *UpdateUser) (*mongo.UpdateResult, error) {
	objectID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{primitive.E{Key: "_id", Value: objectID}}
	update := bson.D{primitive.E{Key: "$set", Value: user}}
	return collection.UpdateOne(ctx, filter, update)
}

func deleteUser(id string) (*mongo.DeleteResult, error) {
	objectID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{primitive.E{Key: "_id", Value: objectID}}
	return collection.DeleteOne(ctx, filter)
}

func findUsers() ([]*User, error) {
	var users []*User
	filter := bson.D{{}}
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return users, err
	}

	for cursor.Next(ctx) {
		var u User
		err := cursor.Decode(&u)
		if err != nil {
			return users, err
		}
		users = append(users, &u)
	}

	if err := cursor.Err(); err != nil {
		return users, err
	}

	cursor.Close(ctx)

	if len(users) == 0 {
		return users, mongo.ErrNoDocuments
	}

	return users, nil
}

func init() {
	//clientOptions := options.Client().ApplyURI("mongodb://user:password@mongo:27017/")
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database("local").Collection("projeto")
}

func main() {
	app := fiber.New()

	app.Get("/api/user", func(c *fiber.Ctx) error {
		users, err := findUsers()
		if err != nil {
			return c.Status(404).SendString(err.Error())
		}
		
		blob, err := json.Marshal(users)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		
		c.Response().BodyWriter().Write(blob)
		return c.SendStatus(200)
	})

	app.Post("/api/user", func(c *fiber.Ctx) error {
		user := new(CreateUser)
		if err := c.BodyParser(user); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			}) 
		}

		errors := ValidateStruct(*user)
		if errors != nil {
       		return c.Status(fiber.StatusBadRequest).JSON(errors)        
    	}

		u, err := createUser(user)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			}) 
		}
		return c.Status(201).JSON(u)
	})

	app.Patch("/api/user/:id", func(c *fiber.Ctx) error {
		user := new(UpdateUser)
		id := c.Params("id")
		if err := c.BodyParser(user); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			}) 
		}

		errors := ValidateStruct(*user)
		if errors != nil {
       		return c.Status(fiber.StatusBadRequest).JSON(errors)        
    	}

		u, err := updateUser(id, user)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			}) 
		}
		return c.Status(200).JSON(u)
	})

	app.Delete("/api/user/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		u, err := deleteUser(id)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			}) 
		}
		return c.Status(200).JSON(u)
	})

	app.Listen(":3000")
}
