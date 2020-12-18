package infra

import (
	"database/sql"
	"log"

	"github.com/raveger/NuxtTodoApp/backend/domain/model"
	"github.com/raveger/NuxtTodoApp/backend/domain/repo"
)

type user struct {
	db *sql.DB
}

func NewUser(db *sql.DB) repo.User {
	return &user{db: db}
}

func (u *user) Users() ([]model.User, error) {
	var (
		id   int
		name string
		TODO string
	)
	users := []model.User{}

	rows, _ := db.Query("SELECT * FROM todolist")
	for rows.Next() {
		err := rows.Scan(&id, &name, &TODO)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		list := model.User{
			ID:   id,
			Name: name,
			Todo: TODO,
		}
		users = append(users, list)
	}
	return users, nil
}

func (u *user) User(id int) (*model.User, error) {
	user := model.User{
		ID: id,
	}
	if err := u.db.Find(&user).Error; err != nil {
		log.Println(err)
		return nil, err
	}
	return &user, nil
}
