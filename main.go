package main

import (
	"log"
	

	"github.com/dayiamin/Fiber_UserManagement_API/internal/database"
	"github.com/dayiamin/Fiber_UserManagement_API/internal/user"
	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
	// "github.com/gofiber/swagger"
	// _ "github.com/dayiamin/Fiber_UserManagement_API/docs"
	
)

func init(){
	var envErr = godotenv.Load()
	if envErr != nil{
		log.Fatal("env error")
	}
}

// @title           Fiber User Management API
// @version         1.0
// @description     REST API for user registration, login, profile access, and admin user listing
// @host            localhost:8080
// @BasePath        /
// @schemes         http
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	app := fiber.New()
	// app.Get("/swagger/*", swagger.HandlerDefault)
	database.ConnectToDb()
	api := app.Group("/api")

	user.RegisterRoutes(api)
	
	log.Fatal(app.Listen(":8080"))
}
