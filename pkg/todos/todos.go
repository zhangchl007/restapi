package todos

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"github.com/zhangchl007/restapi/pkg/models"
	"github.com/zhangchl007/restapi/pkg/db"
)

// createTodo add a new todo
func CreateTodo(c *gin.Context) {
	var db = db.GetDB()
	completed, _ := strconv.Atoi(c.PostForm("completed"))
	todo := models.TodoModel{Title: c.PostForm("title"), Completed: completed}
	db.Save(&todo)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Todo item created successfully!", "resourceId": todo.ID})
}

// fetchAllTodo fetch all todos
func FetchAllTodo(c *gin.Context) {
	var todos []models.TodoModel
	var _todos []models.TransformedTodo
	var db = db.GetDB()
	db.Find(&todos)

	if len(todos) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
		return
	}

	//transforms the todos for building a good response
	for _, item := range todos {
		completed := false
		if item.Completed == 1 {
			completed = true
		} else {
			completed = false
		}
		_todos = append(_todos, models.TransformedTodo{ID: item.ID, Title: item.Title, Completed: completed})
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _todos})
}

// fetchSingleTodo fetch a single todo
func FetchSingleTodo(c *gin.Context) {
	var todo models.TodoModel
	var db = db.GetDB()
	todoID := c.Param("id")

	db.First(&todo, todoID)

	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
		return
	}

	completed := false
	if todo.Completed == 1 {
		completed = true
	} else {
		completed = false
	}

	_todo := models.TransformedTodo{ID: todo.ID, Title: todo.Title, Completed: completed}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _todo})
}

// updateTodo update a todo
func UpdateTodo(c *gin.Context) {
	var todo models.TodoModel
	var db = db.GetDB()
	todoID := c.Param("id")

	db.First(&todo, todoID)

	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
		return
	}

	db.Model(&todo).Update("title", c.PostForm("title"))
	completed, _ := strconv.Atoi(c.PostForm("completed"))
	db.Model(&todo).Update("completed", completed)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Todo updated successfully!"})
}

// deleteTodo remove a todo
func DeleteTodo(c *gin.Context) {
	var todo models.TodoModel
	var db = db.GetDB()
	todoID := c.Param("id")

	db.First(&todo, todoID)

	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
		return
	}

	db.Delete(&todo)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Todo deleted successfully!"})
}

