package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"github.com/raveger/NuxtTodoApp/backend/domain/service"
)

type UserHandler interface {
	Users(c echo.Context) error
	DoRemove(c echo.Context) error
}

type userHandler struct {
	user service.User
}

type ErrorResponse struct {
	Message string `xorm:"message"`
}

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

// DELETE用
func (u *userHandler) DoRemove(c echo.Context) error {
	paramID := c.Param("ID")

	id, _ := strconv.Atoi(paramID)
	err := u.user.DoRemove(id)
	if err != nil {
		log.Println("Delete failed")
		return c.JSON(http.StatusBadRequest, NewErrorResponse(err))
	}
	return c.JSON(http.StatusOK, "OK")
}
