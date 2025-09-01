package main

import (
	"tasks-backend/internal/db"
	"tasks-backend/internal/handlers"
	"tasks-backend/internal/middleware"
	"tasks-backend/internal/repositories"

	"github.com/gin-gonic/gin"
)

func init() {
	db.Connect("192.168.0.142", "5432", "postgres", "postgres", "tasks")
}

func main() {

	r := gin.Default()

	r.Use(middleware.AddLoggingMiddleWare())

	userRepo := repositories.NewUserRepository()
	userHandler := handlers.NewUserHandler(userRepo)

	userHandler.RegisterRoutes(r.Group("users"))
	//err := userService.CreateUser()
	//err := taskRepo.UpdateTask(&models.Task{ID: 2, Title: "Städa ditt rum", Description: "Städa rummet"})

	r.Run(":8080")
	//fmt.Print("task: ", task.Title)
}
