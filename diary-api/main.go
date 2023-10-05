package main

import (
	"diary-api/controller"
	"diary-api/database"
	"diary-api/middleware"
	"diary-api/model"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	loadDatabase()
	serveApplication()
}

func loadEnv() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func loadDatabase() {
	database.Connect()
	database.Database.AutoMigrate(&model.User{})
	database.Database.AutoMigrate(&model.Entry{})
	database.Database.AutoMigrate(&model.Film{})
}

func serveApplication() {
	router := gin.Default()

	publicRoutes := router.Group("/auth")
	publicRoutes.POST("/register", controller.Register)
	publicRoutes.POST("/login", controller.Login)
	publicRoutes.GET("/GetAllUser", controller.GetAllUser)
	publicRoutes.GET("/users/:username", controller.GetUserbyUsername)
	publicRoutes.DELETE("/users/:username", controller.DeleteUser)
	publicRoutes.PATCH("/changepass/:username", controller.ChangePassword)
	

	publicRoutes.POST("/film", controller.AddFilm)
	publicRoutes.GET("/getallfilm", controller.GetAllFilm)
	publicRoutes.GET("/ratingscore/:filmID", controller.RatingScore)

	protectedRoutes := router.Group("/api")
	protectedRoutes.Use(middleware.JWTAuthMiddleware())
	protectedRoutes.POST("/entry", controller.AddEntry)
	protectedRoutes.GET("/entry", controller.GetAllEntries)
	protectedRoutes.DELETE("/:ID", controller.DeleteEntry)
	protectedRoutes.GET("/entry/:user_id", controller.CountEntry)
	protectedRoutes.POST("/vote", controller.AddVote)
	protectedRoutes.GET("/:ID", controller.GetEntrybyId)

	router.Run(":8000")
	fmt.Println("Server running on port 8000")
}
