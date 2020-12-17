package service

import (
	"github.com/raveger/NuxtTodoApp/backend/domain/model"
	"github.com/raveger/NuxtTodoApp/backend/domain/repo"
)

type User interface {
	Users() ([]model.User, error)
	User(id int) (*model.User, error)
}

type user struct {
	repo repo.User
}

func NewUser(r repo.User) User {
	return &user{repo: r}
}

func (u *user) Users() ([]model.User, error) {
	return u.repo.Users()
}

func (u *user) User(id int) (*model.User, error) {
	return u.repo.User(id)
}
