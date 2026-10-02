package users

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type UserRepository struct {
	DB     *mongo.Client
	DBName string
}

func (ur *UserRepository) Create(ctx context.Context, user User) error {
	database := ur.DB.Database(ur.DBName)

	collection := database.Collection("users")

	_, err := collection.InsertOne(ctx, user)

	if err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) GetAll(ctx context.Context) ([]User, error) {
	database := ur.DB.Database(ur.DBName)

	collection := database.Collection("users")

	cursor, err := collection.Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}

	var users []User
	err = cursor.All(ctx, &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (ur *UserRepository) GetUserByID(ctx context.Context, id bson.ObjectID) (User, error) {
	database := ur.DB.Database("test")

	collection := database.Collection("users")

	result := collection.FindOne(ctx, bson.M{"_id": id})

	var user User

	err := result.Decode(&user)

	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (ur *UserRepository) UpdateUser(ctx context.Context, id bson.ObjectID, updateData bson.M) error {
	database := ur.DB.Database(ur.DBName)

	collection := database.Collection("users")

	update := bson.M{"$set": updateData}

	_, err := collection.UpdateOne(ctx, bson.M{"_id": id}, update)

	return err
}

func (ur *UserRepository) DeleteUser(ctx context.Context, id bson.ObjectID) error {
	database := ur.DB.Database(ur.DBName)

	collection := database.Collection("users")

	_, err := collection.DeleteOne(ctx, bson.M{"_id": id})

	if err != nil {
		return err
	}
	return nil
}
