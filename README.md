## Description

brief description of the folder structure:

1. `cmd/`: This folder contains the application's entry point(s) or executable(s).

2. `internal/`: This folder holds the core application code. It is not accessible from outside the module/package.

   - `app/`: This folder contains the application-specific logic.

     - `usecase/`: Contains the application's use cases or business logic.

     - `repository/`: This folder contains interfaces or contracts that define how to interact with external dependencies, such as databases.

   - `entity/`: This folder defines the application's entities or domain models.

   - `delivery/`: Contains the delivery mechanisms, such as HTTP handlers, used to interact with the outside world.

3. `pkg/`: This folder contains shared packages or utilities that can be used by different parts of the application.

   - `config/`: This folder holds configuration-related code. F

4. `migrations/`: This folder may contain database migration scripts or related files.

5. `database/`: This folder may contain database-specific code or configurations.

6. `.github/workflows/`: Script deployment using github action

7. `Jenkinsfile`: Script deployment using jenkins

## ERD Diagram

![alt text](https://github.com/dedihartono801/dating-service/blob/master/erd.png)

## Sequence Diagram

![alt text](https://github.com/dedihartono801/dating-service/blob/master/sequence-diagram.png)

## Install Migration

go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

## Install mock generator (mockgen)

go install github.com/golang/mock/mockgen@v1.6.0

## Create .env file

```bash
$ ./entrypoint.sh
```

## Run Service With Docker

```bash
$ docker-compose up -d
```

## Run Service Without Docker

```bash
$ go run cmd/main.go
```

## Run Migration UP

```bash
$ make migration-up
```

## Run Migration Down

```bash
$ make migration-down
```

## Create Migration

```bash
$ make migration
#type your migration name example: create_create_table_users
```

## Generate Mock repository

- Open Makefile
- Add this code section mock:

```
mockgen -source="YOUR_GO_INTERFACE" -destination="YOUR_MOCK_DESTINATION"
```

Example:

```
mockgen -source="./internal/app/repository/user/user.go" -destination="./internal/app/repository/user/mocks/user_mock.go"
```

```bash
$ make mock
```

## Run Unit Test and Test Coverage

```bash
$ make test-cov
```

## Check Code Smell

```bash
$ make lint
```

## Register User

```bash
$ curl --location 'http://localhost:5004/register' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email":"abc@gmail.com",
    "password":"123",
    "first_name":"dedi",
    "last_name":"hartono",
    "gender":"male",
    "age":23,
    "date_of_birth":"1996-01-01",
    "bio":"bio",
    "location":"jakarta selatan"
}'
```

## Login

```bash
$ curl --location 'http://localhost:5004/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email":"abc@gmail.com",
    "password":"123"
}'
```

## Get List Candidates

```bash
$ curl --location 'http://localhost:5004/users' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjEiLCJlbWFpbCI6ImFiY0BnbWFpbC5jb20iLCJleHAiOjE3MzEzMDEyNzEsImdlbmRlciI6Im1hbGUiLCJpc192ZXJpZmllZCI6dHJ1ZSwiaXNfcHJlbWl1bSI6dHJ1ZX0.-KcCeeg-8lM8T-2Sd7vSvG4jDM9nyfxaYTvuyFDx2xw'
```

## List Payment Methods

```bash
$ curl --location 'http://localhost:5004/payment-method' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjEiLCJlbWFpbCI6ImFiY0BnbWFpbC5jb20iLCJleHAiOjE3MzEzMDEyNzEsImdlbmRlciI6Im1hbGUiLCJpc192ZXJpZmllZCI6dHJ1ZSwiaXNfcHJlbWl1bSI6dHJ1ZX0.-KcCeeg-8lM8T-2Sd7vSvG4jDM9nyfxaYTvuyFDx2xw'
```

## List Packages

```bash
$ curl --location 'http://localhost:5004/package' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjEiLCJlbWFpbCI6ImFiY0BnbWFpbC5jb20iLCJleHAiOjE3MzEzMDEyNzEsImdlbmRlciI6Im1hbGUiLCJpc192ZXJpZmllZCI6dHJ1ZSwiaXNfcHJlbWl1bSI6dHJ1ZX0.-KcCeeg-8lM8T-2Sd7vSvG4jDM9nyfxaYTvuyFDx2xw'
```

## Swipe

```bash
$ curl --location 'http://localhost:5004/swipe' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjEiLCJlbWFpbCI6ImFiY0BnbWFpbC5jb20iLCJleHAiOjE3MzE0MTQ4MDUsImdlbmRlciI6Im1hbGUiLCJpc192ZXJpZmllZCI6dHJ1ZSwiaXNfcHJlbWl1bSI6ZmFsc2V9.aSZ1fTZu_rH9ZH6_V4E9zTffVoAg6fNJ1BVUEkLvFWg' \
--header 'Content-Type: application/json' \
--data '{
    "target_user_id":3,
    "swipe_type":"like"
}'
```

## Purchase

```bash
$ curl --location 'http://localhost:5004/transaction' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjEiLCJlbWFpbCI6ImFiY0BnbWFpbC5jb20iLCJleHAiOjE3MzEzMDEyNzEsImdlbmRlciI6Im1hbGUiLCJpc192ZXJpZmllZCI6dHJ1ZSwiaXNfcHJlbWl1bSI6dHJ1ZX0.-KcCeeg-8lM8T-2Sd7vSvG4jDM9nyfxaYTvuyFDx2xw' \
--header 'Content-Type: application/json' \
--data '{
    "payment_method_id":1,
    "package_type_id":2,
    "amount":500000,
    "currency":"IDR"
}'
```
