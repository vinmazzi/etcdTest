package core

import (
	"context"
	"log"
)

type User struct {
	Id   int    `json:"userId"`
	Name string `json:"userName"`
}

type UserFactory struct {
	Database
}

func NewUserFactory(database Database) *UserFactory {
	return &UserFactory{
		Database: database,
	}
}

func (uf *UserFactory) GetUser(ctx context.Context, id int) (*User, error) {
	user, err := uf.Database.GetUser(ctx, id)
	if err != nil {
		log.Println("Could not get the user: ", err)
		return nil, err
	}

	return user, nil
}
