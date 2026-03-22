# Task Manager API (Go + Echo)

## How to Run

go mod tidy  
go run main.go  

## Endpoints

GET /tasks  
GET /tasks/:id  
POST /tasks  
PUT /tasks/:id  
DELETE /tasks/:id  

## Example

curl -X POST http://localhost:8080/tasks \
-H "Content-Type: application/json" \
-d '{"title":"Learn Go"}'

## Storage
In-memory (data resets on restart)

## Framework
Echo (Golang)

## Limitations
- No database
- No authentication
