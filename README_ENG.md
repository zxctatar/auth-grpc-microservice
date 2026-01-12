# Auth gRPC Microservice

Authentication microservice implemented in Go using gRPC and JWT. The service handles user registration, login, and access token validation.

## Key Features

* User registration
* Authentication (login)
* JWT token generation and validation

## Architecture

The project is implemented using an approach close to Clean Architecture, where business logic is isolated from the transport layer, and all dependencies are directed inward (handlers -> usecase -> infrastructure).

## Service Layers

* **gRPC Transport**
  Implementation of the gRPC server and RPC method handler.

* **Usecase**
  Contains business logic for registration, login, and token validation.

* **Infrastructure**
  Password hashing (bcrypt), JWT handling, and working with the data storage (Postgres) via repository interfaces.

## Project Structure

```text
.
├── cmd
│   └── auth
│       └── main.go
├── config
│   └── local.yaml
├── internal
│   ├── config
│   │   ├── config_test.go
│   │   └── config.go
│   ├── domain
│   │   └── user
│   │       ├── error.go
│   │       ├── user_test.go
│   │       └── user.go
│   ├── infrastructure
│   │   ├── hasher
│   │   │   ├── hasher_test.go
│   │   │   └── hasher.go
│   │   ├── jwt
│   │   │   ├── error.go
│   │   │   ├── jwtservice_test.go
│   │   │   └── jwtservice.go
│   │   └── postgres
│   │       ├── posmodels
│   │       │   ├── user_auth_data_model_test.go
│   │       │   ├── user_auth_data_model.go
│   │       │   ├── user_model_test.go
│   │       │   └── user_model.go
│   │       ├── mapper_test.go
│   │       ├── mapper.go
│   │       ├── postgres_test.go
│   │       └── postgres.go
│   ├── repository
│   │   ├── hashservice
│   │   │   └── hash_service.go
│   │   ├── storagerepo
│   │   │   ├── error.go
│   │   │   └── repo.go
│   │   └── tokenservice
│   │       ├── error.go
│   │       └── token_service.go
│   ├── transport
│   │   └── grpc
│   │       ├── handler
│   │       │   ├── handler_test.go
│   │       │   ├── handler.go
│   │       │   ├── login_usercase_mock_test.go
│   │       │   ├── register_usecase_mock_test.go
│   │       │   └── validate_token_mock_test.go
│   │       ├── pb
│   │       │   ├── auth_grpc.pb.go
│   │       │   └── auth.pb.go
│   │       ├── proto
│   │       │   └── auth.proto
│   │       └── server.go
│   └── usecase
│       ├── errors
│       │   ├── login
│       │   │   └── error.go
│       │   └── registration
│       │       └── error.go
│       ├── implementations
│       │   ├── login
│       │   │   ├── hash_service_mock_test.go
│       │   │   ├── login_test.go
│       │   │   ├── login.go
│       │   │   ├── repo_mock_test.go
│       │   │   └── tok_mock_test.go
│       │   ├── registration
│       │   │   ├── hash_service_mock_test.go
│       │   │   ├── mapper_test.go
│       │   │   ├── mapper.go
│       │   │   ├── register_test.go
│       │   │   ├── register.go
│       │   │   └── repo_mock_test.go
│       │   └── validtoken
│       │       ├── tok_mock_test.go
│       │       ├── validate_token_test.go
│       │       └── validate_token.go
│       ├── interfaces
│       │   ├── login.go
│       │   ├── registration.go
│       │   └── validate_token.go
│       └── models
│           ├── login
│           │   ├── error.go
│           │   ├── login_input_test.go
│           │   ├── login_input.go
│           │   ├── user_auth_data_test.go
│           │   └── user_auth_data.go
│           └── registration
│               ├── error.go
│               ├── reg_input_test.go
│               └── reg_input.go
├── migrations
│   ├── 001_create_table_users.down.sql
│   └── 001_create_table_users.up.sql
├── pkg
│   └── logger
│       └── logger.go
├── .gitignore
├── go.mod
├── go.sum
├── LICENSE
├── Makefile
├── README_RUS.md
└── README.md
```

## JWT

JSON Web Token (JWT) is used for generation and validation:

* HMAC signature is used;
* token lifetime is configured via the .env file.

## gRPC API

### Register

Registers a new user.

**Input:**

* first_name
* middle_name
* last_name
* password
* email

**Output:**

* user_id

---

### Login

Authenticates a user.

**Input:**

* email
* password

**Output:**

* token

---

### ValidateToken

Validates the JWT.

**Input:**

* token

**Output:**

* user_id
* valid_status

## Testing

The project contains unit tests for key components.

`testing` package is used.

Run tests:
```bash
go test ./...
```

## Migrations

Database schema is managed via migrations.

Connection configuration is stored in the `.env` file:

* DB_HOST - database host
* DB_PORT - database port
* DB_NAME - database name
* MIG_NAME - user for managing migrations
* MIG_PASS - password for managing migrations
* SSLMODE - sslmode

Migrations can be executed using Makefile:
```bash
make migrate_all_up
```

## Service Configuration

Service configuration is done via the configuration (.yaml) file:

**gRPC:**
  - `port` - gRPC service port
  - `timeout` - maximum time per request
  - `jwtSecretKey` - key for JWT generation
  - `jwtTimeOut` - JWT token lifetime

**logger:**
  - `level` - logging level (debug, info)

**postgres:**
  - `host` - database host
  - `port` - database port
  - `db_name` - database name
  - `user` - database username
  - `password` - database password
  - `sslmode` - sslmode

## Run

Service can be started using Makefile:
```bash
make local
```
Or
```bash
go run ./cmd/auth/main.go
```

## Technologies Used

* gRPC
* Protocol Buffers
* JWT
* bcrypt
* PostgreSQL

## License

The project is licensed under the MIT License.