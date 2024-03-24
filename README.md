# Go Finance API

This is a simple API for managing personal finances, focusing on expense tracking.

## Getting Started

### Prerequisites
Ensure you have the following installed:
- Go (version 1.22.1 or later)
- Any necessary environment variables or additional setup specified in `go.mod`.

### Installation

1. Clone the repository to your local machine:
```sh
git clone https://github.com/yourusername/go-finance.git
```

2. Navigate to the cloned repository:
```sh
cd go-finance
```

3. To start the server, run the `main.go` file located in the `cmd/api-server` directory:
```sh
go run cmd/api-server/main.go
```

The server will start on port 3000 by default.

## Endpoints

### Expense Management

- **Create a new expense:**

  `POST /v1/expenses`
  
  Request body should include:
  - `amount`: The amount of the expense (a number).
  - `description`: A brief description of the expense (a string).
  - `category`: The category for the expense (a string).
  
  Example request in `create_expense.http`.

- **Update an existing expense:**

  `PUT /v1/expenses/{id}`
  
  Request body should include the same properties as the create endpoint.
  
  Example request in `update_expense.http`.

- **Delete an expense:**

  `DELETE /v1/expenses/{id}`
  
  Example request in `delete_expense.http`.

- **List all expenses:**

  `GET /v1/expenses`
  
  Example request in `list_expense.http`.

- **Get a specific expense:**

  `GET /v1/expenses/{id}`
  
  Example request in `get_expense.http`.

## Project Structure

- `cmd/api-server/main.go`: The entry point of the application.
- `cmd/api-server/handler/expense.go`: The HTTP handlers for the expense-related endpoints.
- `internal/model/expense.go`: The data model for an expense entity.
- `script/`: Directory containing example HTTP requests for API endpoints.

## Dependencies

This project uses Go modules for dependency management. The list of dependencies is defined in the `go.mod` file.
