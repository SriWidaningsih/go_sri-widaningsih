package usecase

import (
	"errors"
	"fmt"
	"project/lib/database"
	"project/models"

	"project/middlewares"
)

func LoginUser(user *models.User) (err error) {
	err = database.GetUser(user)
	if err != nil {
		fmt.Println("GetUser: Error getting user from database")
		return
	}

	token, err := middlewares.CreateToken(int(user.ID))
	if err != nil {
		fmt.Println("GetUser: Error generate token")
		return

	}
	user.Token = token
	return
}

func CreateUser(user *models.User) error {
	if user.Name == "" {
		return errors.New("user name cannot be empty")
	}
	if user.Email == "" {
		return errors.New("user email cannot be empty")
	}
	err := database.CreateUser(user)
	if err != nil {
		return err
	}
	return nil
}

func GetUser(id uint) (user models.User, err error) {
	user.ID = id

	err = database.GetUser(&user)
	if err != nil {
		fmt.Println("GetUser : Error getting user from database")
		return
	}
	blog, err := database.GetBlog(id)
	if err != nil {
		fmt.Println("GetUser: Error getting user from database")
		return
	}
	user.Blogs = append(user.Blogs, blog)
	return
}

func GetListUsers() (users []models.User, err error) {
	users, err = database.GetUsers()
	return
}

func UpdateUser(user *models.User) error {
	err := database.UpdateUser(user)
	if err != nil {
		return err
	}
	return nil
}

func DeleteUser(id uint) error {
	user := models.User{}
	user.ID = id
	err := database.DeleteUser(&user)
	if err != nil {
		return err
	}
	return nil
}
