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

// INSERT用
func (u *user) DoAdd(name string, todo string) error {
	user := model.User{
		Name: name,
		Todo: todo,
	}
	_, err := u.engine.Table("todolist").Insert(&user)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// UPDATE用
func (u *user) DoUpdate(id int, name string, todo string) error {
	user := model.User{
		Name: name,
		Todo: todo,
	}
	_, err := u.engine.Table("todolist").Where("id = ?", id).Update(&user)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// DELETE用
func (u *user) DoRemove(id int) error {
	user := model.User{
		ID: id,
	}
	_, err := u.engine.Table("todolist").Where("id = ?", id).Delete(&user)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
