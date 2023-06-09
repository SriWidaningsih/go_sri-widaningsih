package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
}

type Book struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserId int    `json:"user_id"`
}
type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {

	e := echo.New()
	e.GET("/books", GetBookController)
	e.GET("/books/:id", GetBookByIdController)
	e.POST("/books/post", PostBookController)
	e.DELETE("/Del/:id", DeleteBookController)
	e.Start(":8000")

}

// Untuk mendapatkan semua Data
func GetBookController(c echo.Context) error {
	response, _ := http.Get("https://jsonplaceholder.typicode.com/posts")

	responseData, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	var books []Book
	json.Unmarshal(responseData, &books)
	return c.JSON(http.StatusOK, Response{
		Code:    http.StatusOK,
		Message: "Success to get Book",
		Status:  "Success",
		Data:    books,
	})
}

// Untuk mendapatkan Data, dengan ID tertentu
func GetBookByIdController(c echo.Context) error {
	id := c.Param("id")
	response, _ := http.Get("https://jsonplaceholder.typicode.com/posts" + "/" + id)

	responseData, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	var books Book
	json.Unmarshal(responseData, &books)
	return c.JSON(http.StatusOK, Response{
		Code:    http.StatusOK,
		Message: "Success to get Book",
		Status:  "Success",
		Data:    books,
	})
}

// Untuk menambahkan Data dan disimpan didalamnya
func PostBookController(c echo.Context) error {
	post := Post{
		Title:  "Judul Oleh Wida",
		Body:   "Isi nya, ahirnya berhasil",
		UserID: 1,
	}
	postBody, _ := json.Marshal(post)

	response, _ := http.Post("https://jsonplaceholder.typicode.com/posts", "application/json", bytes.NewBuffer(postBody))
	responseData, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	var newPost Post
	json.Unmarshal(responseData, &newPost)
	return c.JSON(http.StatusOK, Response{
		Code:    http.StatusOK,
		Message: "Success to create new post",
		Status:  "Success",
		Data:    newPost,
	})
}

// Untuk menghapus Data
func DeleteBookController(c echo.Context) error {
	id := c.Param("id")
	response, _ := http.Get("https://jsonplaceholder.typicode.com/posts" + "/" + id)

	responseData, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	var books Book
	json.Unmarshal(responseData, &books)
	return c.JSON(http.StatusOK, Response{
		Code:    http.StatusOK,
		Message: "Success to Delete By Id",
		Status:  "Success",
		Data:    books,
	})
}
