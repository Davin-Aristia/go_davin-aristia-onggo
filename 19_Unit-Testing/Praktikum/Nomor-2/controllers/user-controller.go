package controllers

import (
	"net/http"
	"strconv"

	"test/mvc/models"
	"test/mvc/repositories"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

// get all users
func GetUsersController(c echo.Context) error {
	responseData, err := repositories.GetAllUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "error get data",
			"error":   err,
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success get all users",
		"users":   responseData,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	// your solution here
	id, _ := strconv.Atoi(c.Param("id"))
	responseData, err := repositories.GetUserByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "error get data",
			"error":   err,
		})
	}
	
	return c.JSON(http.StatusOK, map[string]any{
		"message": "success get user with id " + strconv.Itoa(id),
		"user":    responseData,
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	user := models.User{}
	errBind := c.Bind(&user)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error bind data",
			"error":   errBind.Error(),
		})
	}

	responseData, err := repositories.InsertUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "error insert data",
			"error":   err,
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success create new user",
		"user":    responseData,
	})

}

// delete user by id
func DeleteUserController(c echo.Context) error {
	// your solution here
	id, _ := strconv.Atoi(c.Param("id"))

	// Pengecekan apakah user dengan id tersebut ada dalam database
	_, err := repositories.GetUserByID(id)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return c.JSON(http.StatusNotFound, map[string]any{"message": "no user found for id "+strconv.Itoa(id)})
		}
	}

	err = repositories.DeleteUserByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "error delete data",
			"error":   err,
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success delete user with id " + strconv.Itoa(id),
	})
}

// update user by id
func UpdateUserController(c echo.Context) error {
	// your solution here
	id, _ := strconv.Atoi(c.Param("id"))

	updateData := models.User{}
	errBind := c.Bind(&updateData)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error bind data",
			"error":   errBind.Error(),
		})
	}

	// Pengecekan apakah user dengan id tersebut ada dalam database
	user, err := repositories.GetUserByID(id)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return c.JSON(http.StatusNotFound, map[string]any{"message": "no user found for id "+strconv.Itoa(id)})
		}
	}

	responseData, err := repositories.UpdateUserByID(id, user, updateData)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "error update data",
			"error":   err,
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success update user with id " + strconv.Itoa(id),
		"user":    responseData,
	})
}
