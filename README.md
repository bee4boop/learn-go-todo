---
# Go ToDo Application

**Project:** Go ToDo App
**Author:** [Your Name]
**Technologies:** Go, PostgreSQL, Docker, Docker Compose

---

## Overview

This is a simple ToDo application implemented in **Go** with **PostgreSQL** as the database. The project demonstrates basic understanding of backend development, clean architecture principles, and working with HTTP APIs and databases.

The application includes:

- Basic CRUD operations for tasks
- PostgreSQL integration
- Separation between handlers (controllers), service layer, and repository layer
- Ready structure for further learning and extension

---

## Project Structure

```
go-todo/
├── main.go               # Application entry point
├── go.mod
├── go.sum
├── todo/                 # Application logic
│   ├── model.go          # Task model
│   ├── repo.go           # Repository (DB access)
│   ├── service.go        # Business logic (CRUD)
│   └── handler.go        # HTTP handlers / controllers
└── .gitignore
```

- `repo.go` handles database interactions
- `service.go` contains business logic
- `handler.go` manages HTTP endpoints
- `main.go` initializes the application and registers routes

---

## Running with Docker / Compose

### 1. Create `.env` file

```env
POSTGRES_USER=todo
POSTGRES_PASSWORD=todo123
POSTGRES_DB=todo
```

### 2. Start services with Docker Compose

```bash
docker-compose up --build
```

- PostgreSQL database will run on port `5432`
- Go server will run on port `8080`
- The server automatically creates the `tasks` table if it doesn't exist

### 3. Dockerfile

The `Dockerfile` builds the Go binary and runs it inside an Alpine container.

```dockerfile
FROM golang:1.21-alpine
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod tidy
COPY . .
RUN go build -o server main.go
CMD ["./server"]
```

---

## API Endpoints

- `GET /tasks` - List all tasks
- `POST /tasks` - Create a new task
- `POST /tasks/toggle/{id}` - Toggle task completion
- `PUT /tasks/{id}` - Update title and status
- `DELETE /tasks/{id}` - Delete a task

All endpoints use JSON for request and response.

### Task Model

```json
{
  "id": 1,
  "title": "Buy milk",
  "done": false,
  "created_at": "2025-11-05T00:00:00Z"
}
```

---

## Development & Run Instructions

1. Clone the repository:
```bash
git clone <repo-url>
cd go-todo
```

2. Install dependencies:
```bash
go mod tidy
```

3. Set environment variables in `.env` (optional if using Docker Compose)

4. Run the server directly (without Docker):
```bash
go run main.go
```
Server runs at `http://localhost:8080`

---

## Why this project matters

- Demonstrates ability to structure a Go project
- Shows understanding of basic CRUD operations and REST API design
- Illustrates separation of concerns between handlers, service, and repository
- Acts as a starting point for learning more advanced Go and backend development concepts
- Ready for Docker and Compose usage, showing awareness of containerized environments

---

This project is a learning exercise, showing readiness to explore backend development and willingness to grow skills professionally.
