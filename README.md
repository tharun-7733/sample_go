# Go Cryptocurrency API

A **Go REST API** for managing user cryptocurrency coin balances with authentication.

---

## Project Structure

```
Sample_proj/
├── cmd/api/main.go          # Entry point - starts the server
├── api/                      # Legacy files (api.go, api.yaml)
├── internal/                 # Core application logic
│   ├── handlers/            # HTTP request handlers
│   │   ├── api.go          # Route setup & middleware
│   │   └── get_coin_balance.go
│   ├── middleware/          # Custom middleware
│   │   └── authorization.go
│   ├── shared/              # Shared types & error handlers
│   │   ├── types.go
│   │   └── errors.go
│   └── tools/               # Database layer
│       ├── database.go      # Interface definitions
│       └── mockDB.go        # Mock database implementation
├── go.mod                    # Module: github.com/avukadin/goapi
└── go.sum
```

---

## How It Works

### 1. Server Entry Point (`cmd/api/main.go`)
- Starts HTTP server on `localhost:8000`
- Uses **Chi router** for routing
- Sets up error logging with **Logrus**
- Registers routes through the handlers package

### 2. Routing (`internal/handlers/api.go`)
```
/account              (Protected route)
  └── /coins         (GET) - Returns user's coin balance
```

**Route Protection:**
- Requires `Authorization` middleware (token validation)
- Query parameters: `username`
- Header: `Authorization: <token>`

### 3. Request Flow

```
GET /account/coins?username=alex
Header: Authorization: 123ABC
    ↓
[Authorization Middleware] - Validates username & token
    ↓
[GetCoinBalance Handler] - Fetches coin balance
    ↓
[Mock Database] - Returns coin data
    ↓
JSON Response: {"balance": 100, "code": 200}
```

### 4. Authentication (`internal/middleware/authorization.go`)
- Checks `username` query parameter
- Validates `Authorization` header token
- Compares against mock login details
- Returns 401 Unauthorized if invalid

### 5. Database Layer (`internal/tools/`)
- **Interface-based design** - easy to swap implementations
- **Mock Database** with hardcoded users:

| Username | Token | Coins |
|----------|-------|-------|
| Tharun | 123ABC | 100 |
| Tej | 456DEF | 200 |
| Aksh | 789GHI | 300 |

### 6. Shared Types (`internal/shared/`)
- `CoinBalanceParams` - Request parameters
- `CoinBalanceResponse` - Success response
- `ErrorResponse` - Error handling

---

## Dependencies

- `github.com/go-chi/chi` - Fast HTTP router
- `github.com/sirupsen/logrus` - Structured logging
- `github.com/gorilla/schema` - Query parameter parsing

---

## Running the Server

```bash
go run ./cmd/api/main.go
```

Server starts on `http://localhost:8000`

---

## Testing with Postman

### ✅ Success Request
```
GET http://localhost:8000/account/coins?username=alex
Header: Authorization: 123ABC
```

**Response:**
```json
{
  "balance": 100,
  "code": 200
}
```

### ❌ Failed Auth (Missing Token)
```
GET http://localhost:8000/account/coins?username=alex
```

**Response:**
```json
{
  "error": "Invalid username or token.",
  "code": 401
}
```

### ❌ Failed Auth (Invalid Token)
```
GET http://localhost:8000/account/coins?username=alex
Header: Authorization: WRONG_TOKEN
```

**Response:**
```json
{
  "error": "Invalid username or token.",
  "code": 401
}
```

---

## Available Users for Testing

- **Username:** alex | **Token:** 123ABC | **Balance:** 100 coins
- **Username:** jason | **Token:** 456DEF | **Balance:** 200 coins
- **Username:** marie | **Token:** 789GHI | **Balance:** 300 coins
