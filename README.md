# go-build-service

Rest service template in go

## Set up

To run postgres and adminer, use make as below:

`make dev-up` runs postgres and adminer containers. You will need to have docker and make installed and ready.
Adminer can be accessed on `localhost:8080`.

### Migrations

To run the single example migration: `make migrate-up`

Create a new migration: `make migrate-create NAME=migration_name`

Roll back migrations: `make migrate-down`

### Running the app

`go run main.go`

See the `makefile` for teardown and other commands.


## Features

- Base api using fiber
- Environment variable loading and validation
- Structured logging and request logger middleware
- Postgres database
- Database migrations
- Tooling and lint
- Authentication
  - Stateless JWT access token
  - Revokable (log out) refresh token using jti
  - Authentication middleware

## Roadmap

- Dockerizing
- Tests
- Event processing
