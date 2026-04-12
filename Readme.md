# 🚀 Go Task Management API (Dockerized with SQL Server)

A backend service built with **Go** for managing tasks, projects, and users.
The application is fully containerized using **Docker** and uses **Microsoft SQL Server** as the database.

---

## 🧱 Tech Stack

* **Backend:** Go (Golang)
* **Database:** Microsoft SQL Server
* **Containerization:** Docker & Docker Compose
* **Architecture:** Modular (cmd, service, types, utils)

---

## 📁 Project Structure

```
.
├── cmd/
│   ├── api/        # Entry point (main.go)
│   ├── config/     # Configuration handling
│   ├── db/         # Database connection logic
│   └── migrate/    # Migration scripts
├── service/
│   ├── auth/
│   ├── middleware/
│   ├── project/
│   ├── task/
│   └── user/
├── types/          # Structs and models
├── utils/          # Utility functions
├── Dockerfile
├── docker-compose.yml
├── go.mod
└── go.sum
```

---

## ⚙️ Features

* Task CRUD operations
* Project and user management
* Filter tasks by:

  * Status
  * Assignee
  * Project
* SQL Server integration
* Dockerized environment
* Environment-based configuration

---

## 🐳 Running the Application (Docker)

### 1. Clone the repository

```
git clone <your-repo-url>
cd <your-project>
```

---

### 2. Start services

```
docker-compose up --build
```

---

### 3. Services

| Service     | URL                   |
| ----------- | --------------------- |
| Backend API | http://localhost:9020 |
| SQL Server  | localhost:9010        |

---

## 🔌 API Endpoints

### Health Check

```
GET /health
```

Response:

```
OK
```


---

## 🗄️ Database Configuration

The application connects using environment variables see .example.env:

```
DB_HOST=Your-sqlserver
DB_PORT=1433
DB_USER=Ryuk-Bhai
DB_PASSWORD=YourPassw0rd
DB_NAME=master
```


## ⚠️ Important Notes

* Do NOT use `localhost` as DB host inside Docker
* Use `sqlserver` (service name) instead
* SQL Server may take a few seconds to start — restart Go backend logic if required
---

## 🧪 Development (Without Docker)

```
make run 
```

Make sure SQL Server is running locally and update `.env` accordingly.

---

## 🛠️ Common Issues

### 1. Connection refused

* Ensure SQL Server container is running
* Check port mappings

### 2. Socket hang up

* Ensure server binds to `:8080`
* Check container logs:

  ```
  docker-compose logs -f backend
  ```

### 3. Environment variables not loading

* Use `docker-compose.yml` `environment` section
* Avoid relying solely on `.env` inside containers

---

## 📌 Future Improvements

* Swagger/OpenAPI documentation
* Pagination & sorting
* Unit and integration tests

---

## 👨‍💻 Author

---
