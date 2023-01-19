package services

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

func createUser(user *dtos.CreateUser) (*mongo.InsertOneResult, error) {
	return collection.InsertOne(ctx, user)
}

func updateUser(id string, user *dtos.UpdateUser) (*mongo.UpdateResult, error) {
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

func findUsers() ([]*models.User, error) {
	var users []*models.User
	filter := bson.D{{}}
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return users, err
	}

	for cursor.Next(ctx) {
		var u models.User
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

