package usecase

import (
	"errors"
	"project/lib/database"
	"project/models"
)

func CreateBlog(blog *models.Blog) error {
	if blog.Title == "" {
		return errors.New("blog title cannot be empty")
	}
	if blog.Content == "" {
		return errors.New("blog content cannot be empty")
	}
	err := database.CreateBlog(blog)
	if err != nil {
		return err
	}
	return nil
}

func GetBlog(id uint) (blog models.Blog, err error) {
	blog, err = database.GetBlog(id)
	return
}

func GetListBlogs() (blogs []models.Blog, err error) {
	blogs, err = database.GetBlogs()
	return
}

func UpdateBlog(blog *models.Blog) error {
	err := database.UpdateBlog(blog)
	if err != nil {
		return err
	}
	return nil
}

func DeleteBlog(id uint) error {
	blog := models.Blog{}
	blog.ID = id
	err := database.DeleteBlog(&blog)
	if err != nil {
		return err
	}
	return nil
}
