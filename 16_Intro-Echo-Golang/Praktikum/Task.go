package main

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"strconv"
)

type User struct {
	Id int `json:"id" form:"id"`
	Name string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var users []User


// -------------------- controller --------------------

// get all users
func GetUsersController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all users",
		"users": users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	id,_ := strconv.Atoi(c.Param("id"))

	for _,user := range users{
		if user.Id == id{
			return c.JSON(200, map[string]interface{}{
				"messages": "success get user with id " + strconv.Itoa(id),
				"user": user,
			})
		}
	}
	return c.JSON(404, map[string]string{
		"messages": "no user found for id " + strconv.Itoa(id),
	})
	
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	id,_ := strconv.Atoi(c.Param("id"))
	
	for index,user := range users{
		if user.Id == id{
			users = append(users[:index], users[index+1:]...)
			return c.JSON(200, map[string]interface{}{
				"messages": "success delete user with id " + strconv.Itoa(id),
			})
		}
	}
	return c.JSON(404, map[string]string{
		"messages": "no user found for id " + strconv.Itoa(id),
	})
}

// update user by id
func UpdateUserController(c echo.Context) error {
	id,_ := strconv.Atoi(c.Param("id"))
	userParam := User{}
	c.Bind(&userParam)

	for index,user := range users{
		if user.Id == id{
			users[index] = userParam
			users[index].Id = id
			return c.JSON(200, map[string]interface{}{
				"messages": "success update user with id " + strconv.Itoa(id),
				"user": users[index],
			})
		}
	}
	return c.JSON(404, map[string]string{
		"messages": "no user found for id " + strconv.Itoa(id),
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	// binding data
	user := User{}
	c.Bind(&user)

	if len(users) == 0 {
		user.Id = 1
	} else {
		newId := users[len(users)-1].Id + 1
		user.Id = newId
	}

	users = append(users, user)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create user",
		"user": user,
	})
}

// ---------------------------------------------------
func main() {
	e := echo.New()

	// routing with query parameter
	e.GET("/users", GetUsersController)
	e.POST("/users", CreateUserController)

	e.GET("/users/:id", GetUserController)
	e.DELETE("/users/:id", DeleteUserController)
	e.PUT("/users/:id", UpdateUserController)

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}
