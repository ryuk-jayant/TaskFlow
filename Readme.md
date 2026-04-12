# 📌 TaskFlow — Submission Justification & Notes

## 1. Overview

This project is my implementation of the **TaskFlow Engineering Home Assignment**, focused primarily on delivering a **complete, production-ready backend system** using Go, Docker, and SQL Server (adapted from PostgreSQL due to familiarity and setup constraints).

The backend supports:

* Authentication (JWT-based)
* Project and Task management
* Filtering and relational data handling
* Dockerized environment with persistent storage

While the assignment also includes a frontend requirement, this submission prioritizes **backend, architecture, and reliability**.

---

## 2. What Has Been Completed

### ✅ Backend (Fully Implemented)

All core backend requirements from the assignment have been completed:

#### 🔐 Authentication

* User registration and login endpoints
* Password hashing using bcrypt
* JWT-based authentication (24-hour expiry)
* Protected routes using middleware

#### 📦 Projects API

* Create, fetch, update, and delete projects
* Ownership validation (only owners can modify/delete)
* Fetch project with associated tasks

#### 📝 Tasks API

* Create, update, delete tasks
* Dynamic filtering (status, assignee)
* Partial updates supported (PATCH semantics)
* Proper relational mapping with project and user

#### 🧠 API Design

* RESTful endpoints
* Proper HTTP status codes (`200`, `201`, `400`, `401`, `403`, `404`, `500`)
* Structured error responses
* Clean separation of concerns (handler → service → store)

#### 🗄️ Data Layer

* Relational schema aligned with assignment requirements
* Parameterized queries (safe from SQL injection)
* Dynamic query building for filtering and updates

#### 🐳 Docker & Dev Environment

* Fully dockerized backend + database
* Persistent database using volumes
* Environment-driven configuration
* Single command setup (`docker-compose up --build`)

---

## 3. What Is Missing (Frontend)

### ❌ Frontend (Not Completed)

The frontend portion of the assignment has **not been implemented**.

### Reason:

Although I have backend experience, I am **very new to Go**, and a significant portion of the allotted time was spent on:

* Understanding Go Syntax and patterns
* Debugging containerization and networking issues
* Implementing a correct backend system
* Ensuring proper API behavior and database interaction

Rather than rushing a frontend implementation, I chose to **focus on delivering a solid backend**, as that aligns better with my current strengths.

---

## 4. Engineering Decisions & Tradeoffs

### 🧠 Focus on Backend Depth Over Full Stack Breadth

* Prioritized correctness, stability, and structure over feature breadth
* Ensured all backend flows are production-grade and debuggable

---

### ⚙️ Ms-SQL Server Instead of PostgreSQL

* Used SQL Server due to familiarity and faster setup in Docker
* Maintained relational integrity and schema consistency
* All queries are portable with minor syntax adjustments

---

### 🧩 No ORM — Raw SQL

* Chose raw SQL over ORM for:

  * Better control
  * Performance
  * Transparency
* Tradeoff: slightly more verbose code

---

### 🔄 Dynamic Query Building

* Implemented flexible filtering and update logic
* Avoided hardcoded queries
* Ensured safe parameter binding

---

## 5. Challenges Faced

* Docker networking (container-to-container communication)
* Environment variable handling inside containers
* SQL query debugging and scan mismatches
* Learning Go-specific patterns (error handling, struct design)

---

## 6. What I am Missing

### 🎨 Frontend

* Build a React + TypeScript UI
* Implement authentication flow
* Add project/task dashboards
* Ensure responsive and polished UX

---

<!--  -->

### 📊 Enhancements

* Pagination for list endpoints
* Task statistics endpoint (`/projects/:id/stats`)
* Role-based access improvements

---

### 🧱 Infrastructure Improvements

* Seed scripts for initial data
* Deep Health checks and readiness probes

---

## 7. Final Note

This submission reflects my approach as an engineer:

* Prioritize correctness over completeness
* Understand systems deeply rather than superficially
* Be honest about tradeoffs and limitations

While the frontend is not included, the backend is designed to **fully support it**, and extending this into a complete full-stack application would be straightforward.

---

## 🙏 Thank You
---
