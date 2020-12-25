package repo

import "github.com/raveger/NuxtTodoApp/backend/domain/model"

type User interface {
	Users() ([]model.User, error)
	User(id int) (*model.User, error)
	// INSERT用
	DoAdd(name string, todo string) error
	// DELETE用
	DoRemove(id int) error
}
