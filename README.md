# Tlab Technical Test

## System Requirement

- Go 1.21.4+
- Docker Compose

Following requirement provided in docker compose :

- PostgreSQL

## How to Install

1. Clone project
2. Copy `.env.example` to `.env`
3. Adjust configuration in `.env`
4. Install the package, simply run this command
    ```
    $ go mod tidy
    ```

## Database Migration

Once application build successfully you can run database migration with :

Migration up

```bash
$ make migration up
```

Migration down

```bash
$ make migration down
```

### Generate Migration File

Requirement : [sql-migrate](https://github.com/rubenv/sql-migrate)

```bash
$ sql-migrate new <migration_name>
```

## Unit Test

```bash
$ go test ./... --cover -coverprofile=coverage.out
```

### Directory Structure

Berikut struktur direktori yang digunakan

```
├── src
│   ├── application
│   │   ├── client_1
│   │   ├── client_2
│   │   └── client_3
│   ├── domain
│   ├── infrastrucure
│   │   ├── http
│   │   ├── repository
│   │   └── pkg
│   ├── ...
└── ...
```

The directory structure above explains that the main code is located in the `src` directory. This directory contains at least three subdirectories.

#### Domain

This directory contains the business logic of the application. All process logic is recommended to be placed only in this directory.

#### Infrastructure

This directory contains adapters and abstraction layers for interactions with all processes outside the application ecosystem. These include: database persistence, HTTP service and router, SMTP adapter, etc.

#### Application

This directory contains the operational behavior of the domain. This directory is responsible for and facilitates the execution of the logic within the domain.


## Endpoints
---
You can check list of endpoint on the [Postman]([https://blue-moon-221631.postman.co/workspace/My-Workspace~7ab23d94-3d8d-46f2-8be3-c4f74dcf0d25/collection/29535426-6aa0b350-f8e9-49a7-8ef8-4b0bf104797b?action=share&creator=29535426](https://github.com/rilgilang/tlab-technical)) 
