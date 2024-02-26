## Swagger
http://localhost:3001/docs/index.html

## Features
 - Register
 - Login
 - Auth
 - CRUD with GORM user_profiles, user_address

## Tech stack
 - docker compose
 - flat architecture
 - golang echo
 - go viper
 - go validator
 - gorm
 - HTTP
 

## localhost run
**run service**:
``` bash
go run main.go
```

**env**:
```bash
DATABASE_USER=postgres
DATABASE_HOST=localhost
DATABASE_PASSWORD=1234
DATABASE_PORT=5432
DATABASE_NAME=homework1
HTTP_PORT=":3001"
JWT_SECRET=UucwjDH7AY40XLDyWpBUalCB151WgAfF
SECRET_KEY=L1K0zInpkIYzVXqUQdvnOc7FtbKOvpsJ
```

**docker**:
```bash
docker-compose build
```
```bash
docker-compose up
```