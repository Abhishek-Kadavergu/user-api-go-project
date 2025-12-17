# Go Fiber User Management API

A production-grade RESTful API built with **Go (Golang)** and **Fiber**, featuring **PostgreSQL** for persistence and **SQLC** for type-safe database interactions.

The core logic involves managing user data (`Name`, `DOB`) while **dynamically calculating the user's Age** at runtime, ensuring data consistency without storing redundant fields.

---

## 🚀 Tech Stack

- **Language:** [Go](https://go.dev/) (1.23+)
- **Framework:** [Fiber v2](https://gofiber.io/) (Fastest HTTP engine for Go)
- **Database:** [PostgreSQL](https://www.postgresql.org/)
- **Database Access:** [SQLC](https://sqlc.dev/) + [pgx/v5](https://github.com/jackc/pgx)
- **Validation:** [go-playground/validator](https://github.com/go-playground/validator)
- **Logging:** [Uber Zap](https://github.com/uber-go/zap)

---

## 🏗️ Project Architecture

This project follows **Clean Architecture** principles to ensure separation of concerns and testability.

```text
/cmd
  /server         # Entry point (main.go)
/config           # Database connection configuration
/db
  /migrations     # SQL schema definitions
  /sqlc           # Auto-generated Go code from SQLC
/internal
  /handler        # HTTP Layer (Parses requests, Validates input)
  /service        # Business Logic (Calculates Age, handles data rules)
  /repository     # Data Access Layer (Interacts with DB via SQLC)
  /routes         # Route definitions
  /models         # Request/Response structs

```
## Setup & Installation

### Prerequisites

- Go installed on your machine
- PostgreSQL installed and running

### Database Setup

1. Open pgAdmin or your terminal.
2. Create a new database named `userdb`.
3. Run the following SQL query to create the table:

```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    dob DATE NOT NULL
);
```
### Environment Configuration
```You need to set the DATABASE_URL environment variable to point to your local PostgreSQL instance.```

For Windows (PowerShell):
```powershell
$env:DATABASE_URL="postgres://postgres:YOUR_PASSWORD@localhost:5432/userdb"
```

## Run the Application

Run the following commands in the project root:

```bash
# Download dependencies
go mod tidy

# Start the server
go run cmd/server/main.go
```
### 🔌 API Endpoints
1. Create User

POST /users

Request Body
```json
{
  "name": "Alice",
  "dob": "1990-05-10"
}
```
Response
```json
{
  "id": 1,
  "name": "Alice",
  "dob": "1990-05-10"
}
```

### 2. Get User (With Age Calculation)

`GET /users/:id`

#### The API fetches the Date of Birth (DOB) from the database and dynamically calculates the age in the Service layer.

Response
```json

{
  "id": 1,
  "name": "Alice",
  "dob": "1990-05-10",
  "age": 35
}
```

### 3. List All Users

`GET /users`

Returns a list of all users with their calculated ages.

### 4. Update User

`PUT /users/:id`

Request Body
```json
{
  "name": "Alice Updated",
  "dob": "1995-01-01"
}
```

### 5. Delete User

`DELETE /users/:id`
