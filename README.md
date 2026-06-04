# E-Commerce API - Learning Project

This is a learning project created to revisit and reinforce key software engineering principles and best practices.

## Purpose

This project serves as a recap/refresher on:

- **Architecture Patterns**: Layered MVC architecture with clear separation of concerns (Controllers → Services → Repositories → Models)
- **API Design**: RESTful API design principles with proper HTTP methods and status codes
- **Authentication & Authorization**: JWT-based authentication with Bearer tokens
- **Error Handling**: Structured error responses with meaningful error codes
- **Database Design**: Relational database modeling with GORM ORM, foreign keys, and migrations
- **Documentation**: API documentation with Swagger/OpenAPI standards
- **Middleware**: Request validation and authentication middleware patterns
- **Code Organization**: Package structure, dependency injection, and clean code principles

## Tech Stack

- **Language**: Go 1.25
- **Web Framework**: Gin
- **Database**: GORM with SQLite/PostgreSQL
- **Authentication**: JWT + bcrypt
- **Documentation**: Swagger/OpenAPI 3.0
- **Server**: Port 3004

## Quick Start

### Prerequisites
- Go 1.25+
- PostgreSQL or SQLite (default)

### Setup

1. Clone the repository
2. Set required environment variables:
   ```bash
   export SECRET_KEY="your-secret-key-at-least-32-characters"
   ```

3. Run the application:
   ```bash
   go run main.go
   ```

4. Access API:
   - **API Base URL**: http://localhost:3004/api/v1
   - **Swagger UI**: http://localhost:3004/swagger/index.html

## API Endpoints

### Public (No Auth Required)
- `POST /auth/signup` - Register new user
- `POST /auth/login` - Login and get JWT token

### Protected (Require Bearer Token)
- **Products**: `GET, POST, PUT, DELETE /products`
- **Cart**: `GET, POST, PUT, DELETE /cart`
- **Addresses**: `GET, POST, PUT, DELETE /addresses`

### How to Use Protected Endpoints

1. Login to get token:
   ```bash
   curl -X POST http://localhost:3004/api/v1/auth/login \
     -H "Content-Type: application/json" \
     -d '{"email":"user@example.com","password":"password123"}'
   ```

2. Use token in requests:
   ```bash
   curl http://localhost:3004/api/v1/cart \
     -H "Authorization: Bearer YOUR_ACCESS_TOKEN"
   ```

3. Or use Swagger UI:
   - Click the 🔒 Authorize button
   - Paste your access token
   - All protected endpoints will auto-include the token

## Project Structure

```
├── controllers/       # HTTP request handlers
├── services/         # Business logic layer
├── repository/       # Data access layer (GORM)
├── models/           # Data models
├── middleware/       # Request/auth middleware
├── routes/           # Route definitions
├── database/         # Database setup
├── utils/            # Utilities (errors, tokens, config)
├── docs/             # Swagger generated docs
├── main.go           # Entry point
└── README.md         # This file
```

## Learning Topics Covered

### 1. Clean Architecture
- Layered separation with clear boundaries
- Dependency injection pattern
- Interface-based design

### 2. Error Handling
- Custom error types with error codes
- Proper HTTP status code mapping
- Structured error responses

### 3. Authentication
- JWT token generation and validation
- Bearer token middleware
- Password hashing with bcrypt

### 4. Database
- ORM usage (GORM)
- Model relationships (1:N, M:N)
- Auto-migrations
- Soft deletes

### 5. API Documentation
- Swagger annotations
- Security definitions
- Request/response schema documentation

### 6. Middleware & Routing
- Route grouping
- Middleware application
- Global middleware patterns

### 7. Configuration Management
- Environment variables
- Default values with validation
- Multi-database support

## Development Notes

- **Graceful Shutdown**: Application properly closes database connections on SIGINT/SIGTERM
- **Context Handling**: Proper context management throughout request lifecycle
- **Pagination**: Built-in pagination support for list endpoints
- **Validation**: Request validation at multiple layers (binding, service, database)
- **Database Flexibility**: Support for both SQLite and PostgreSQL

## Key Principles Demonstrated

✅ **SOLID Principles**: Single Responsibility, Open/Closed, Dependency Inversion  
✅ **DRY (Don't Repeat Yourself)**: Generic repository pattern, middleware reuse  
✅ **Error Handling**: Structured, meaningful error messages with codes  
✅ **Security**: Password hashing, JWT validation, Bearer token authentication  
✅ **Documentation**: Auto-generated API docs with Swagger  
✅ **Code Organization**: Clear package structure and separation of concerns  
✅ **Best Practices**: Context usage, graceful shutdown, environment-based config  

---

*This project is a learning exercise and not meant for production use without additional security hardening, testing, and monitoring.*
