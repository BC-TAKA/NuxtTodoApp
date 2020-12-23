package repo

import "github.com/raveger/NuxtTodoApp/backend/domain/model"

type User interface {
	Users() ([]model.User, error)
	User(id int) (*model.User, error)
	// DELETE用
	Delete(id int) (int64, error)
}
