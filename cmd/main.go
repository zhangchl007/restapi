package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zhangchl007/restapi/pkg/db"
	"github.com/zhangchl007/restapi/pkg/todos"
	"log"
)


func main() {
	log.Println("Starting server..")
	db.Init()
	router := gin.Default()

	v1 := router.Group("/api/v1/todos")
	{
		v1.POST("/", todos.CreateTodo)
		v1.GET("/", todos.FetchAllTodo)
		v1.GET("/:id", todos.FetchSingleTodo)
		v1.PUT("/:id", todos.UpdateTodo)
		v1.DELETE("/:id", todos.DeleteTodo)
	}
	router.Run(":8080")
}
