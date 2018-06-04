# Movies Service
> A simple movies CRUD Restful API in Go that serves as a POC for a microservice

## Getting started

### Locally

**Pre-requisite**: Be sure you have an instance of MongoDB running.

```
go get -d -v ./
DB_SERVER=localhost DB_NAME=movies go run main.go
```

*Note*: Replace the `DB_` variables with your own values pointing to your MongoDB server.

### With Docker

```
docker-compose up --build
```

## Configuration

There are to mandatory environment variables:

- `DB_SERVER`: The hostname of the mongo server.
- `DB_NAME`: The name of the database.
