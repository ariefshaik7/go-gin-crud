package users

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type UserService struct {
	*UserRepository
}

func (us *UserService) CreateUser(ctx context.Context, user User) error {
	err := us.Create(ctx, user)
	return err
}

func (us *UserService) GetAll(ctx context.Context) ([]User, error) {
	users, err := us.UserRepository.GetAll(ctx)
	return users, err
}

func (us *UserService) GetUserByID(ctx context.Context, id bson.ObjectID) (User, error) {
	user, err := us.UserRepository.GetUserByID(ctx, id)
	return user, err
}

func (us *UserService) UpdateUser(ctx context.Context, id bson.ObjectID, updateData bson.M) error {
	err := us.UserRepository.UpdateUser(ctx, id, updateData)
	return err
}

func (us *UserService) DeleteUser(ctx context.Context, id bson.ObjectID) error {
	err := us.UserRepository.DeleteUser(ctx, id)
	return err
}
