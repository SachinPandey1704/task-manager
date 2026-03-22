# Task Manager API

## Overview
This is a simple Task Manager backend built using Go and the Echo framework.  
It supports basic CRUD operations (Create, Read, Update, Delete) for managing tasks.

---

## How to Clone and Run the Project

### Clone the repository
git clone https://github.com/SachinPandey1704/task-manager.git
cd task-manager

---

### Run the project
go mod tidy
go run main.go

---

### Server runs on:
http://localhost:8080

---

## API Endpoints

### 1. Get all tasks
GET /tasks
curl http://localhost:8080/tasks

---

### 2. Get task by ID
GET /tasks/:id
curl http://localhost:8080/tasks/1

---

### 3. Create a task
POST /tasks
curl -X POST http://localhost:8080/tasks \
-H "Content-Type: application/json" \
-d '{"title":"My task","description":"Test task"}'

---

### 4. Update a task
PUT /tasks/:id
curl -X PUT http://localhost:8080/tasks/1 \
-H "Content-Type: application/json" \
-d '{"status":"done"}'

---

### 5. Delete a task
DELETE /tasks/:id
curl -X DELETE http://localhost:8080/tasks/1

---

## Storage Used
In-memory storage (Go map)  
Data will be lost when the server restarts.

---

## Framework Used
Echo framework in Go

---

## Limitations
- Data is not stored permanently
- No authentication
- Basic validation only

---

## Future Improvements
- Add database (SQLite)
- Better validation
- Improved structure
