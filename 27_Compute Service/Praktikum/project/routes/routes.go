package routes

import (
	"project/constants"
	"project/controllers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	e.POST("/login", controllers.LoginUserController)

	//user
	user := e.Group("/users", middleware.JWT([]byte(constants.SECRET_JWT)))

	user.GET("", controllers.GetUsersController)
	user.GET("/:id", controllers.GetUserController)
	user.POST("", controllers.CreateUserController)
	user.DELETE("/:id", controllers.DeleteUserController)
	user.PUT("/:id", controllers.UpdateUserController)

	//Book
	book := e.Group("/books", middleware.JWT([]byte(constants.SECRET_JWT)))
	book.GET("", controllers.GetBooksController)
	book.GET("/:id", controllers.GetBookController)
	book.POST("", controllers.CreateBookController)
	book.DELETE("/:id", controllers.DeleteBookController)
	book.PUT("/:id", controllers.UpdateBookController)

	//Blog
	e.GET("/blogs", controllers.GetBlogsController)
	e.GET("/blogs/:id", controllers.GetBlogController)
	e.POST("/blogs", controllers.CreateBlogController)
	e.DELETE("/blogs/:id", controllers.DeleteBlogController)
	e.PUT("/blogs/:id", controllers.UpdateBlogController)

	return e
}
