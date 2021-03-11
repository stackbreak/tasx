# Tasx

> **`Tasx`** is an example microservice app built with Go.

## About

This app provides **REST API** for the creation and management of the task lists linked to users.

For user access separation is used **JWT**-based authentication system.

_It was created for educational purposes._

## Requirements
- Go `1.16+`
- Go Modules
- Docker & Docker-Compose

## Getting Started

### `.env` files

Before starting, you should create two `.env` files with environment variables needed to configure the app and pass private secrets, for **dev** and **prod** mode.
- **`.env.local`**
  ```json
  APP_PORT=4000

  DB_USER=example_dbuser // replace to yours
  DB_PASS=example_dbpass // replace to yours
  DB_NAME=example_dbname // replace to yours
  DB_HOST=pgsql          // addr in docker-compose network
  DB_PORT=5432
  DB_SSLMODE=disable

  TOKEN_SECRET=super@HYPER!secret12345 // replace to yours
  ```

- **`.env.prod`**
  ```json
  APP_PORT=4000
  GIN_MODE=release

  DB_USER=prod_dbuser // replace to yours
  DB_PASS=prod_dbpass // replace to yours
  DB_NAME=prod_dbname // replace to yours
  DB_HOST=localhost   // replace to yours
  DB_PORT=5432        // replace to yours
  DB_SSLMODE=disable

  TOKEN_SECRET=super@HYPER!secret12345 // replace to yours
  ```

### dev mode

1. **install deps:**
  `go mod download`

1. **build migration tool:**
  `make install-tool-migrate`

1. **run local postgres:**
  `make compose-pgsql`

1. **update migrations:**
  `make migrate-up`

1. **build dev container:**
  `make compose-build`

1. **run app in dev container** (with hot-reload):
  `make compose-up`

> for debug with `dlv` replace in `docker-compose.yaml` command part `./entry.sh watch` to `./entry.sh debug`.

### prod mode

1. **build prod container:**
  `make prod-build` _(it creates `tasx_app` image)_

1. **update migrations** (on target database):
  `make prod-migrate-up`

1. **run app container:**
  `make prod-run`

## REST API Endpoints

### Healthcheck

- **GET** **`/healthz`**

  **Response:** `200 OK`
  ```json
  {
    "status": "available"
  }
  ```
  ---

### Authentication

- **POST** **`/auth/signup`**

  **Body:**
  ```json
  {
    "name": "John Doe",
    "username": "johndoe777",
    "password": "secret_pass"
  }
  ```
  **Response:** `200 OK`

  ---
- **POST** **`/auth/login`**

  **Body:**
  ```json
  {
    "username": "johndoe777",
    "password": "secret_pass"
  }
  ```
  **Response:** `200 OK`
  ```json
  {
    "token": <TOKEN>
  }
  ```

### Task Lists

- **POST** **`/api/lists`**

  **Authorization:** Bearer \<TOKEN>

  **Body:**
  ```json
  {
    "title": "list name",
    "description": "list description"
  }
  ```
  **Response:** `200 OK`

  ---

- **GET** **`/api/lists`**

  **Authorization:** Bearer \<TOKEN>

  **Response:** `200 OK`
  ```json
  {
    "lists": [
      {
        "id": <ID>,
        "title": string,
        "description": string
      }
    ]
  }
  ```
  ---

- **GET** **`/api/lists/:id`**

  **Authorization:** Bearer \<TOKEN>

  **Response:** `200 OK`
  ```json
  {
    "id": <ID>,
    "title": string,
    "description": string
  }
  ```
  ---

- **PUT** **`/api/lists/:id`**

  **Authorization:** Bearer \<TOKEN>

  **Body:**
  ```json
  {
    "title": "Updated Name",      // optional
    "description": "Updated Desc" // optional
  }
  ```
  **Response:** `200 OK`

  ---

- **DELETE** **`/api/lists/:id`**

  **Authorization:** Bearer \<TOKEN>

  **Response:** `200 OK`

### Tasks of the List

- **POST** **`/api/lists/:id/items`**

  **Authorization:** Bearer \<TOKEN>

  **Body:**
  ```json
  {
    "title": "task title",
    "description": "task desc"
  }
  ```
  ---

- **GET** **`/api/lists/:id/items`**

  **Authorization:** Bearer \<TOKEN>

  **Response:** `200 OK`
  ```json
  {
    "tasks": [
      {
        "id": <ID>,
        "title": string,
        "description": string,
        "is_done": boolean
      }
    ]
  }
  ```
  ---

### Tasks by ID

- **GET** **`/api/items/:id`**

  **Authorization:** Bearer \<TOKEN>

  **Response:** `200 OK`
  ```json
  {
    "task": {
      "id": <ID>,
      "title": string,
      "description": string,
      "is_done": boolean
    }
  }
  ```
  ---

- **PUT** **`/api/items/:id`**

  **Authorization:** Bearer \<TOKEN>

  **Body:**
  ```json
  {
    "title": "task title",      // optional
    "description": "task desc", // optional
    "is_done": true             // optional
  }
  ```

  **Response:** `200 OK`

  ---

- **DELETE** **`/api/items/:id`**

  **Authorization:** Bearer \<TOKEN>

  **Response:** `200 OK`

  ---
