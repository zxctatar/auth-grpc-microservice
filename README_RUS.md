# Auth gRPC Microservice

Микросервис аутентификации и авторизации, реализованный на Go с использованием gRPC и JWT. Сервис отвечает за регистрацию пользователей, вход в систему и проверку валидности токенов доступа.

## Основные возможности

* Регистрация пользователей
* Аутентификация (логин)
* Генерация и валидация JWT-токенов

## Архитектура

Проект реализован с использованием подхода близкого к Clean Architecture, где бизнес-логика изолирована от транспортного слоя, все зависимости направлены внутрь (handlers -> usecase -> infrastructure).

## Слои сервиса

* **gRPC транспорт**
  Реализация gRPC сервера и обработчик RPC методов.

* **Usecase**
  Содержит бизнес-логику регистрации, логина, валидации токенов.

* **Infrastructure**
  Хэширование паролей (bcrypt), работа с JWT и работа с хранилищем данных (Postgres) через интерфейсы репозиториев.

## Структура проекта

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

Для генерации и валидации токенов используется JSON Web Token (JWT):

* используется подпись HMAC;
* срок жизни токена задается через конфигурацию (.env).

## gRPC API

### Register

Регистрирует нового пользователя.

**Вход:**

* first_name
* middle_name
* last_name
* password
* email

**Выход:**

* user_id

---

### Login

Аутентифицирует пользователя.

**Вход:**

* email
* password

**Выход:**

* token

---

### ValidateToken

Проверяет валидность JWT.

**Вход:**

* token

**Выход:**

* user_id
* valid_status

## Тестирование

Проект содержит unit-тесты для ключевых компонентов.

Для тестирования используется `testing`

Запуск тестов:
```bash
go test ./...
```

## Миграции

Для управления схемой базы данных используются миграции.

Конфигурация подключения хранится в `.env` файле:

* DB_HOST - адрес базы данных
* DB_PORT - порт базы данных
* DB_NAME - название базы данных
* MIG_NAME - имя пользователя для управления миграциями
* MIG_PASS - пароль пользователя для управления миграциями
* SSLMODE - sslmode

Для запуска миграций можно использовать Makefile:
```bash
make migrate_all_up
```

## Конфигурация сервиса

Настройка сервиса осуществляется через конфигурационный (.yaml) файл:

**gRPC:**
  - `port` - порт gRPC сервиса
  - `timeout` - максимум времени на один запрос
  - `jwtSecretKey` - ключ для генерации JWT-токенов
  - `jwtTimeOut` - время валидности JWT-токена

**logger:**
  - `level` - уровень логирования (debug, info)

**postgres:**
  - `host` - адрес базы данных
  - `port` - порт базы данных
  - `db_name` - название базы данных
  - `user` - имя пользователя
  - `password` - пароль пользователя
  - `sslmode` - sslmode

## Запуск

Для запуска можно использовать Makefile:
```bash
make local
```
Или
```bash
go run ./cmd/auth/main.go
```

## Используемые технологии

* gRPC
* Protocol Buffers
* JWT
* bcrypt
* PostgreSQL

## Лицензия

Проект распространяется под лицензией MIT.