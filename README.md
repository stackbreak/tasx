# Todo

## env file example

`configs/local.env`

```
DB_USER=secret_user
DB_PASS=secret_pass
DB_NAME=secret_dbname
DB_HOST=localhost
DB_PORT=5432
DB_SSLMODE=disable

PGADMIN_EMAIL=secret@email.org
PGADMIN_PASS=secret_pass
PGADMIN_PORT=4444

TOKEN_SECRET=super@HYPER!secret
```


## REST API Endpoints

### Authentication

- **POST** **`/auth/signup`**

  **Body:**
  ```json
  {
    "name": "Sector Wins",
    "username": "shellslayer",
    "password": "123qwe"
  }
  ```
  **Response:** `200 OK`

  ---
- **POST** **`/auth/login`**

  **Body:**
  ```json
  {
    "username": "shellslayer",
    "password": "123qwe"
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
