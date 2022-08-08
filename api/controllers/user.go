package controllers

import (
	"echo/repositories"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetUsers(c echo.Context) error {
	users := repositories.GetUsers()
	return c.JSON(http.StatusOK, users)
}

func GetUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	user := repositories.GetUser(id)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, nil)
	}

	if user.Name == "" {
		return c.JSON(http.StatusOK, nil)
	} else {
		return c.JSON(http.StatusOK, user)
	}
}

func CreateUser(c echo.Context) error {
	user, err := repositories.CreateUser(c)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, nil)
	}
	return c.JSON(http.StatusOK, user)
}

func UpdateUser(c echo.Context) error {
	user, err := repositories.UpdateUser(c)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, nil)
	}
	return c.JSON(http.StatusOK, user)
}

func DeleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	repositories.DeleteUser(id)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, nil)
	}
	return c.NoContent(http.StatusNoContent)
}
