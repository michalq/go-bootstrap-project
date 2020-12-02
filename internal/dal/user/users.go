package user

import (
	"context"
	"fmt"

	"github.com/michalq/go-bootstrap-project/internal/modules/user"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Users mongo implementatin of users repository
type Users struct {
	mongoDB *mongo.Database
}

// NewUsers creates new users repository
func NewUsers(mongoDB *mongo.Database) *Users {
	return &Users{mongoDB}
}

func (u *Users) dbCollection() *mongo.Collection {
	return u.mongoDB.Collection("user")
}

// FindOneByID Returns user by id
func (u *Users) FindOneByID(userID string) (*user.User, error) {
	var userEntity *user.User
	filter := bson.D{{"id", userID}}
	err := u.dbCollection().FindOne(context.TODO(), filter).Decode(userEntity)
	if err != nil {
		return nil, err
	}

	return userEntity, nil
}

// FindAll finds all created users
func (u *Users) FindAll() ([]*user.User, error) {
	return nil, nil
}

// Save saves user in database
func (u *Users) Save(userEntity *user.User) (string, error) {
	result, err := u.dbCollection().InsertOne(context.TODO(), userEntity)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s", result.InsertedID), nil
}

// Update saves user in database
func (u *Users) Update(userEntity *user.User) error {
	_, err := u.dbCollection().InsertOne(context.TODO(), userEntity)
	if err != nil {
		return err
	}

	return nil
}
