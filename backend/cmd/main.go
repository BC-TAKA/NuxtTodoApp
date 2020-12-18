package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/raveger/NuxtTodoApp/backend/domain/service"
	"github.com/raveger/NuxtTodoApp/backend/infra"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func main() {
	var err error
	config, err := config.Readconfig()
	if err != nil {
		log.Fatal(err)
	}
	e := echo.New()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=Asia%%2FTokyo",
		config.DB.ID, config.DB.Password, config.DB.Host, config.DB.Port, config.DB.DB)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	userRepo := infra.NewUser(db)
	userService := service.NewUser(userRepo)
	h := handler.NewHandler(userService)
	server := server.New(e).RegistHandler(h)

	log.Fatal(server.Start())
}
