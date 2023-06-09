package main

import (
	"project/config"
	"project/middlewares"
	"project/routes"
)

func main() {
	config.InitDB()
	e := routes.New()
	middlewares.LogMiddlewares(e)

	e.Logger.Fatal(e.Start(":8000"))
}
