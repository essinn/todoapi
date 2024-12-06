package controllers

import (
	"errors"
	"net/http"
	"todoapi/pkg/models"

	"github.com/gin-gonic/gin"
)

var todos = []models.Todo{
	{ID: "1", Item: "Learn Go", Completed: true},
	{ID: "2", Item: "Create Project", Completed: true},
	{ID: "3", Item: "Finish Assignment", Completed: false},
}

func GetTodos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, todos)
}

// NewTodo handles creating a new todo
func NewTodo(c *gin.Context) {
	var newTodo models.Todo

	if err := c.BindJSON(&newTodo); err != nil {
		return
	}

	todos = append(todos, newTodo)
	c.IndentedJSON(http.StatusCreated, newTodo)
}

// getTodoById finds a todo by ID
func getTodoById(id string) (*models.Todo, error) {
	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}
	return nil, errors.New("todo not found")
}

// GetTodo handles fetching a todo by ID
func GetTodo(c *gin.Context) {
	id := c.Params.ByName("todoId")
	todo, err := getTodoById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, todo)
}

// UpdateTodo toggles the completion status of a todo
func UpdateTodo(c *gin.Context) {
	id := c.Params.ByName("todoId")
	todo, err := getTodoById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}

	todo.Completed = !todo.Completed
	c.IndentedJSON(http.StatusOK, todo)
}
