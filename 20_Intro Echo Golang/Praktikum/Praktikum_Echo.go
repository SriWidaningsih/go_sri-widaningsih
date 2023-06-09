package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type User struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var users []User
var mapUsers map[int]User
var lastId int

// -------------------- controller --------------------

// get all users
func GetUsersController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all users",
		"users":    users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	userId := GetIdFromParameter(c)
	if val, ok := mapUsers[userId]; ok {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"messages": "Success Get user by Id",
			"user":     val,
		})

	}
	return c.JSON(http.StatusNotFound, map[string]interface{}{
		"messages": "Error Get user by Id",
		"user":     nil,
	})

}

// delete user by id
func DeleteUserController(c echo.Context) error {
	userId := GetIdFromParameter(c)
	for index, val := range users {
		if val.Id == userId {
			users = append(users[:index], users[index+1:]...)
			delete(mapUsers, userId)

			return c.JSON(http.StatusOK, map[string]interface{}{
				"messages": "Success Delete user by Id",
			})
		}
	}
	return c.JSON(http.StatusNotFound, map[string]interface{}{
		"messages":         "Error Delete user by Id",
		"ErrorDescription": "ID Not Found",
	})

}

// update user by id
func UpdateUserController(c echo.Context) error {
	userId := GetIdFromParameter(c)
	updateUser := User{}
	c.Bind(&updateUser)
	updateUser.Id = userId

	for _, val := range users {
		if val.Id == userId {
			val = updateUser
			mapUsers[userId] = updateUser
			return c.JSON(http.StatusOK, map[string]interface{}{
				"messages": "Succsess Update user by Id",
			})
		}
	}
	return c.JSON(http.StatusBadRequest, map[string]interface{}{
		"messages":         "Error Update user by Id",
		"ErrorDescription": "Not Found",
	})

	// your solution here
}

// create new user
func CreateUserController(c echo.Context) error {
	// binding data
	user := User{}
	c.Bind(&user)

	if len(users) == 0 {
		user.Id = 1
		mapUsers = make(map[int]User)
	} else {
		newId := users[len(users)-1].Id + 1
		user.Id = newId
	}
	users = append(users, user)
	mapUsers[user.Id] = user
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create user",
		"user":     user,
	})
}

func GetIdFromParameter(c echo.Context) int {
	userId, _ := strconv.Atoi(c.Param("id"))
	return userId
}

// ---------------------------------------------------
func main() {
	e := echo.New()
	e.GET("/users", GetUsersController)
	e.GET("/users/:id", GetUserController)
	e.DELETE("/users/:id", DeleteUserController)
	e.PUT("/users/:id", UpdateUserController)
	e.POST("/users", CreateUserController)

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":1010"))
}
