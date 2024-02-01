package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID string `json:"id"`
	Task string `json:"task"`
	Status string `json:"status"`
}

var todos = []todo{
	{ID: "1", Task: "Learn Go", Status: "Ongoing"},
	{ID: "2", Task: "Learn Gin", Status: "Ongoing"},
	{ID: "3", Task: "Build REST APIs", Status: "Ongoing"},
}

func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func getTodoById(context *gin.Context) {
	id := context.Param("id")
	flag := false

	for _, todo := range todos {
		if todo.ID == id {
			context.IndentedJSON(http.StatusFound, todo)
			flag = true
		}
	}

	if !flag {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
	}
}

func addTodos(context *gin.Context) {
	var newTodo todo

	err := context.BindJSON(&newTodo)
	if err != nil {
		return
	}

	todos = append(todos, newTodo)
	context.IndentedJSON(http.StatusCreated, todos)
}

func updateTodo(context *gin.Context) {
	var updatedTodo todo
	err := context.BindJSON(&updatedTodo)
	if err != nil {
		return
	}

	flag := false

	for id, todo := range todos {
		if todo.ID == updatedTodo.ID {
			todos[id] = updatedTodo
			context.IndentedJSON(http.StatusAccepted, todos[id])
			flag = true
		}
	}

	if !flag {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
	}
}

func main() {
	router := gin.Default()

	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodoById)
	router.POST("/add-todos", addTodos)
	router.PUT("/update-todo", updateTodo)

	router.Run("localhost:8000")
}