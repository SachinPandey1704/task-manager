package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}

var tasks = make(map[int]Task)
var idCounter = 1

func main() {
	e := echo.New()

	e.GET("/tasks", getTasks)
	e.GET("/tasks/:id", getTask)
	e.POST("/tasks", createTask)
	e.PUT("/tasks/:id", updateTask)
	e.DELETE("/tasks/:id", deleteTask)

	e.Logger.Fatal(e.Start(":8080"))
}

// GET all tasks
func getTasks(c echo.Context) error {
	return c.JSON(http.StatusOK, tasks)
}

// GET one task
func getTask(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	task, exists := tasks[id]
	if !exists {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Task not found"})
	}

	return c.JSON(http.StatusOK, task)
}

// CREATE task
func createTask(c echo.Context) error {
	task := new(Task)

	if err := c.Bind(task); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	if task.Title == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Title is required"})
	}

	task.ID = idCounter
	task.Status = "pending"
	task.CreatedAt = time.Now()

	tasks[idCounter] = *task
	idCounter++

	return c.JSON(http.StatusCreated, task)
}

// UPDATE task
func updateTask(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	task, exists := tasks[id]
	if !exists {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Task not found"})
	}

	updated := new(Task)
	if err := c.Bind(updated); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	if updated.Title != "" {
		task.Title = updated.Title
	}
	if updated.Description != "" {
		task.Description = updated.Description
	}
	if updated.Status != "" {
		task.Status = updated.Status
	}

	tasks[id] = task

	return c.JSON(http.StatusOK, task)
}

// DELETE task
func deleteTask(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	_, exists := tasks[id]
	if !exists {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Task not found"})
	}

	delete(tasks, id)

	return c.JSON(http.StatusOK, map[string]string{"message": "Task deleted"})
}
