package api

import (
	"database/sql"

	"github.com/raveger/VueTODO/ToDo/model"
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
