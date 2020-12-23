package infra

import (
	"log"

	"github.com/raveger/NuxtTodoApp/backend/domain/model"
	"github.com/raveger/NuxtTodoApp/backend/domain/repo"
	"xorm.io/xorm"
)

type user struct {
	engine *xorm.Engine
}

func NewUser(engine *xorm.Engine) repo.User {
	return &user{engine: engine}
}

func (u *user) Users() ([]model.User, error) {
	users := []model.User{}

	if err := u.engine.Table("todolist").Find(&users); err != nil {
		log.Println(err)
		return nil, err
	}
	return users, nil
}

func (u *user) User(id int) (*model.User, error) {
	user := model.User{
		ID: id,
	}
	if err := u.engine.Table("todolist").Find(&user); err != nil {
		log.Println(err)
		return nil, err
	}
	return &user, nil
}

// DELETE用
// func (u *user) Delete(id int) (err error) {
// 	err = u.engine.ID(id).Delete(&user)

// 	if err != nil {
// 		log.Println(err)
// 	}
// 	return
// }
