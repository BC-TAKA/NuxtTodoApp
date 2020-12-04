package api

import (
	"database/sql"

	"github.com/raveger/VueTODO/ToDo/model"
)

func UpdateTODO(todo model.TODO, DB *sql.DB) error {
	_, err := DB.Exec(
		`UPDATE todolist SET name = ?, TODO = ? WHERE id = ?`,
		todo.Name, todo.Todo, todo.ID,
	)
	if err != nil {
		return err
	}
	return nil
}
