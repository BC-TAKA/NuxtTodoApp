package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"github.com/raveger/NuxtTodoApp/backend/domain/model"
	"github.com/raveger/NuxtTodoApp/backend/domain/service"
)

type UserHandler interface {
	Users(c echo.Context) error
	DoAdd(c echo.Context) error
	DoUpdate(c echo.Context) error
	DoRemove(c echo.Context) error
}

type userHandler struct {
	user service.User
}

type ErrorResponse struct {
	Message string `xorm:"message"`
}

// 書く場所はmodel？

func NewErrorResponse(err error) ErrorResponse {
	return ErrorResponse{
		Message: err.Error(),
	}
}

func NewUserHandler(user service.User) UserHandler {
	return &userHandler{user: user}
}

func (u *userHandler) Users(c echo.Context) error {
	paramID := c.QueryParam("id")

	if paramID == "" {
		users, err := u.user.Users()
		if err != nil {
			err = errors.Wrap(err, "cannot get users")
			log.Println(err)
			return c.JSON(http.StatusInternalServerError, NewErrorResponse(err))
		}
		return c.JSON(http.StatusOK, users)
	}
	id, err := strconv.Atoi(paramID)
	if err != nil {
		err = errors.Wrap(err, "Invalid user id")
		log.Println(err)
		return c.JSON(http.StatusBadRequest, NewErrorResponse(err))
	}
	user, err := u.user.User(id)
	if err != nil {
		err = errors.Wrap(err, fmt.Sprintf("cannot get user %d", id))
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, NewErrorResponse(err))
	}
	return c.JSON(http.StatusOK, user)
}

// INSERT用
func (u *userHandler) DoAdd(c echo.Context) error {
	// 構造体の定義する場所について要確認
	var i model.InsertParams
	err := c.Bind(&i)
	log.Println(i)

	if err != nil {
		log.Println("INSERT failed")
		return c.JSON(http.StatusBadRequest, NewErrorResponse(err))
	}
	if err := u.user.DoAdd(i.Name, i.Todo); err != nil {
		log.Println("Insert failed")
		return c.JSON(http.StatusBadRequest, NewErrorResponse(err))
	}
	return c.JSON(http.StatusOK, "INSERT OK")
}

// UPDATE用
func (u *userHandler) DoUpdate(c echo.Context) error {
	var i model.User
	err := c.Bind(&i)
	log.Println(i)

	if err != nil {
		log.Println("UPDATE failed")
		return c.JSON(http.StatusBadRequest, NewErrorResponse(err))
	}
	if err := u.user.DoUpdate(i.ID, i.Name, i.Todo); err != nil {
		log.Println("Update failed")
		return c.JSON(http.StatusBadRequest, NewErrorResponse(err))
	}
	return c.JSON(http.StatusOK, "UPDATE OK")
}

// DELETE用
func (u *userHandler) DoRemove(c echo.Context) error {
	paramID := c.Param("ID")

	id, _ := strconv.Atoi(paramID)
	err := u.user.DoRemove(id)
	if err != nil {
		log.Println("Delete failed")
		return c.JSON(http.StatusBadRequest, NewErrorResponse(err))
	}
	return c.JSON(http.StatusOK, "DELETE OK")
}
