package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/raveger/NuxtTodoApp/backend/app/server/handler"
	"github.com/raveger/NuxtTodoApp/backend/config"
	"github.com/raveger/NuxtTodoApp/backend/domain/service"
	"github.com/raveger/NuxtTodoApp/backend/infra"
	"xorm.io/xorm"
)

func main() {
	var err error
	config, err := config.Readconfig()
	if err != nil {
		log.Fatal(err)
	}
	e := echo.New()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=Asia%%2FTokyo",
		config.DB.ID, config.DB.Password, config.DB.Host, config.DB.Port, config.DB.DB)
	engine, err := xorm.NewEngine("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	userRepo := infra.NewUser(engine)
	userService := service.NewUser(userRepo)
	h := handler.NewUserHandler(userService)

	router := httprouter.New()

	e.Start()
}
