# friend-management
This project implements a simple Friend Management system, structured as a RESTful API

# Table of Contents

1. [Introduction](#introduction)
2. [Prerequisites](#prerequisites)
3. [How to run?](#how-to-run)
4. [Project Structure](#project-structure)
5. [Database Schema](#database-schema)
5. [APIs](#apis)
6. [Technologies](#technologies)

## Introduction

`friend-management` is a simple friend management system structured as a RESTful API. This project allows users to manage friend relationships and interactions through key features such as creating friend connections, listing friends, listing mutual friends, subscribing to receive updates from an email, blocking updates from an email, and retrieving the list of people who are eligible to receive updates from an email's activities.

The system is designed to provide an easy-to-use and scalable API, enabling other applications or services to integrate friend management capabilities seamlessly. Developed by Mai NguyenT.

## Prerequisites

Before you continue, ensure you meet the following requirements:

* You have installed Docker.
* Your OS has `make tool`
* You have a basic understanding of API.

## How to run?

Quick run: use the `make start` command to quick run the application in dev mode.

Build: use `make setup` then use `make build` to run the application in built mode.

You can also use the `make` or `make help` command to see usage guidelines.

## Project Structure

```bsh
friend-management/
┣ cmd/
┃ ┣ router/
┃ ┃ ┗ router.go
┃ ┗ main.go
┣ data/
┃ ┗ migrations/
┃   ┣ 000001_init_db.down.sql
┃   ┗ 000001_init_db.up.sql
┣ internal/
┃ ┣ controller/
┃ ┃ ┣ relationship/
┃ ┃ ┗ user/
┃ ┣ handler/
┃ ┃ ┣ rest/
┃ ┃ ┃ ┗ v1/
┃ ┃ ┗ handler.go
┃ ┣ model/
┃ ┗ repository/
┃   ┣ relationship/
┃   ┗ user/
┣ pkg/
┣ .env.example
┣ .gitignore
┣ docker-compose.yaml
┣ Dockerfile
┣ Dockerfile.release
┣ go.mod
┣ go.sum
┣ makefile
┣ README.md
┗ sqlboiler.yaml
```

1. **/cmd** contains the main application entry point files for the project.
2. **/data** contains data files such as migration files or seeding files.
3. **/internal** contains the private library code, it is specific to the function of the service and not shared with other services. It includes handler, service, repository, and model packages.
    1. _/handler_ contains all the handler functions that will process requests, call functions from service and return responses.
    2. _/controller_ contains all the domain functions that will take parameters, perform operations, and return results.
    3. _/repository_ contains all the functions for performing database queries.
    4. _/model_ contains all the ORM models.
4. **vendor** contains the vendor libraries that are used by the application.
5. **/pkg** contains code that is OK for other services to consume, this includes utility functions.
6. **Makefile** is a special file containing shell commands, such as “start” commands.
7. **Dockerfile** is a text document that contains all the commands a user could call on the command line to assemble an image which provide dev tools.
8. **Dockerfile.release** a instruction to build multi stages image for building and running binaries.
9. **docker-compose.yml** provides a way to document and configure all the application's service dependencies (databases, queues, caches, web service APIs, etc).
10 **go.mod** is the root of dependency management in the Go language.
11 **.env.example** is a simple text configuration file storing key-value pairs as environment variables. Reference for .env

## Database Schema

### `users` table:
| Fields     | Datatype      | Description                  | Refer to |
|------------|---------------|------------------------------|----------|
| id         | SERIAL        | PRIMARY KEY (AUTO_INCREMENT) |          |
| email      | TEXT NOT NULL |                              |          |
| created_at | TIMESTAMP Z   | DEFAULT NOW()                |          |
| updated_at | TIMESTAMP Z   | DEFAULT NOW()                |          |

### `relationship_type` type enum (`friend`, `subscribe`, `block`)

### `relationships` table:
| Fields        | Datatype                    | Description                  | Refer to  |
|---------------|-----------------------------|------------------------------|-----------|
| id            | SERIAL                      | PRIMARY KEY (AUTO_INCREMENT) |           |
| requester_id  | INT NOT NULL                | FOREIGN KEY                  | users(id) |
| target_id     | INT NOT NULL                | FOREIGN KEY                  | users(id) |
| type          | relationships_type NOT NULL |                              |           |
| created_at    | TIMESTAMP Z                 | DEFAULT NOW()                |           |
| updated_at    | TIMESTAMP Z                 | DEFAULT NOW()                |           |

## APIs
### 1. Create friend connection

**Endpoints: `POST /api/v1/relationships/friends`**

1.1. Request body

- `friends`: array of two email addresses that will be friends.

Example:
```json
{
  "friends":[
    "user1@example.com",
    "user2@example.com"
  ]
}
```

1.2. Response body

- Successful:
    - `success`: the success of the request.

```json
{
 "success": true
}
```
- Failure: invalid_request_body

```json
{"message": "Invalid request body","code": "invalid_request_body","status": 400}
```

- Failure: invalid_given_email

```json
{"message":"Invalid given email","code":"invalid_given_email","status":400}
```

- Failure: cannot_be_friend_with_self

```json
{"desc":"Cannot be friend with self","code":"cannot_be_be_friend_with_self","status":400}
```

- Failure: already_friends

```json
{"message":"Already friend","code":"already_friends","status":400}
```

- Failure: already_blocked

```json
{"message":"Already blocked","code":"already_blocked","status":400}
```

- Failure: user_not_found

```json
{"message":"User not found with given email","code":"user_not_found_with_given_email","status":400}
```

### 2. Retrieve friends by email

*Endpoints: `POST /api/v1/relationships/friends/list`*

2.1. Request body

- `email`: the email address of the user who needs to get the friends.

```json
{
  "email": "user1@example.com"
}
```

2.2. Response body

- Successful:

    - `success`: the success of the request.
    - `friends`: list of friends.

```json
{
 "success": true,
 "friends":[
  "user2@example.com",
  "user3@example.com"
 ]
}
```

- Failure: invalid_given_email

```json
{"message":"Invalid given email","code":"invalid_given_email","status":400}
```

### 3. Get common friends:

Endpoints: `POST /api/v1/relationships/friends/list-common`

3.1. Request body

- `friends`: the emails to get common friends

```json
{
  "friends":[
    "user1@example.com",
    "user2@example.com"
  ]
}
```

3.2. Response body

- Successful:

    - `success`: the success of the request.
    -  `friends`: list of friends.

```json
{
 "success": true,
 "friends":[
  "user3@example.com"
 ]
}
```

- Failure: invalid_given_email

```json
{"message":"invalid_given_email","code":"invalid_given_email","status":400}
```

- Failure: cannot_list_common_friends_with_self

```json
{"message":"Cannot list common friends with self","code":"cannot_list_common_friends_with_self","status":400}
```

### 4. Subscribe to updates:

*Endpoints: `POST /api/v1/relationships/subscribes`*

4.1. Request body

- `requestor`: user needs to subscribe.
- `target`: user will be subscribed.

```json
{
  "requestor":"user1@example.com",
  "target":"user2@example.com"
}
```

4.2. Response body

- Successful:

    - `success`: the success of the request

```json
{
 "success": true
}
```

- Failure: invalid_given_email

```json
{"message":"invalid_given_email","code":"invalid_given_email","status":400}
```

- Failure: cannot_self_subscribe

```json
{"message":"Cannot self subscribe","code":"cannot_self_subscribe","status":400}
```

- Failure: user_not_found_with_given_email

```json
{"message":"User not found with given email","code":"user_not_found_with_given_email","status":400}
```

- Failure: already_subscribe

```json
{"message":"Already subscribe","code":"already_subscribe","status":400}
```

### 5. Block updates:
*Endpoints: `POST /api/v1/relationships/blocks`*

5.1. Request body

- `requestor`: user needs to block.
- `target`: user will be blocked.

```json
{
  "requestor":"user1@example.com",
  "target":"user2@example.com"
}
```

5.2. Response body

- Successful:

    - `success`: the success of the request

```json
{
  "success": true
}
```

- Failure: invalid_given_email

```json
{"message":"invalid_given_email","code":"invalid_given_email","status":400}
```

- Failure: cannot_self_subscribe

```json
{"message":"Cannot self subscribe","code":"cannot_self_subscribe","status":400}
```

- Failure: user_not_found_with_given_email

```json
{"message":"User not found with given email","code":"user_not_found_with_given_email","status":400}
```

- Failure: already_subscribe

```json
{"message":"Already blocked","code":"already_blocked","status":400}
```

### 6. Get recipients:

*Endpoints: `POST /api/v1/updates/recipients`*

6.1. Request body

- `sender`: the author of the update.
- `text`: the content of the update.

```json
{
  "sender":"user1@example.com",
  "text": "Hello world! user2@example.com"  
}
```

6.2. Response body

- Successful:

    - `success`: the success of the request.
    - `recipients`: the recipients.

```json
{
 "success": true,
 "recipients":[
  "user2@example.com",
  "user3@example.com"
 ]
}
```

- Failure: invalid_given_email

```json
{"message":"Invalid given email","code":"invalid_given_email","status":400}
```

### 7. Create user:

*Endpoints: `POST /api/v1/users`*

6.1. Request body

- `email`: user unique email.

```json
{
  "email":"user1@example.com"
}
```

6.2. Response body

- Successful:

    - `success`: the success of the request.

```json
{
 "success": true
}
```

- Failure: invalid_given_email

```json
{"message":"Invalid_given_email","code":"invalid_given_email","status":400}
```

- Failure: already_exists

```json
{"message":"User already exists","code":"already_exists","status":400}
```

## Technologies
- **Golang 1.23.2**: Main programming language used for the API development.
- **Docker**: Containerization tool for application deployment and environment consistency.
- **Postgres 15**: Database used to store and manage application data.
- **SQLBoiler**: ORM for Go, used for generating strongly-typed models.
- **Mockery**: Mocking library to create mocks for unit testing.
- **Go-migrate**: Database migration tool for handling schema changes.
