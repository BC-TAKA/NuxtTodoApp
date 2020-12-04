package api

import "database/sql"

func DeleteTODO(id int, DB *sql.DB) error {
	_, err := DB.Exec("DELETE FROM todolist WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}
