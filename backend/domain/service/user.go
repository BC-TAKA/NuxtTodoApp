package service

import (
	"github.com/raveger/NuxtTodoApp/backend/domain/model"
	"github.com/raveger/NuxtTodoApp/backend/domain/repo"
)

type User interface {
	Users() ([]model.User, error)
	User(id int) (*model.User, error)
	// INSERT用
	DoAdd(name string, todo string) error
	// DELETE用
	DoRemove(id int) error
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

// INSERT用
func (u *user) DoAdd(name string, todo string) error {
	return u.repo.DoAdd(name, todo)
}

// DELETE用
func (u *user) DoRemove(id int) error {
	return u.repo.DoRemove(id)
}
