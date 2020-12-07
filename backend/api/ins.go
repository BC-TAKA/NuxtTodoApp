package api

import (
	"database/sql"

	"github.com/raveger/NuxtTodoApp/backend/model"
)

func CreateTODO(todo model.TODO, DB *sql.DB) error {
	_, err := DB.Exec(
		`INSERT INTO todolist(name,TODO) VALUES( ? , ? )`,
		todo.Name, todo.Todo,
	)
	if err != nil {
		return err
	}
	return nil
}
