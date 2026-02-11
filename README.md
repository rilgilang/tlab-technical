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


# TLab API Documentation

## Collection Overview
Postman collection for TLab API endpoints. Base URL is configured using `{{host}}` variable.

## Authentication
Most endpoints require Bearer token authentication. The token is automatically captured from the Login response and stored in the `{{access_token}}` variable.


## Endpoints
---
You can check list of endpoint on the Postman
https://dsi-aicare-marketplace.postman.co/workspace/My-Workspace~c2d2de3d-2007-443f-81d2-b45e28972599/collection/31714573-155a48d5-079b-42d9-b015-66a073682ebd?action=share&creator=31714573&active-environment=31714573-521dacd5-27bd-48b2-a366-03bab710da3e
