# ESB Technical Test

## Description

This is a project built with Golang Clean architecture.

## Tech Stack

- Golang v1.20.13
- MySQL (Database) v5.7

## Framework & Library

- GoFiber (HTTP Framework) : https://github.com/gofiber/fiber
- GORM (ORM) : https://github.com/go-gorm/gorm
- Viper (Configuration) : https://github.com/spf13/viper
- Golang Migrate (Database Migration) : https://github.com/golang-migrate/migrate
- Go Playground Validator (Validation) : https://github.com/go-playground/validator
- Logrus (Logger) : https://github.com/sirupsen/logrus

## Configuration

All configuration is in `config.json` file.

## API Spec

All API collection is in `api` folder.

## Database

### Install Mysql v5.7

via docker
```shell
docker-compose up -d
```

## Database Migration

All database migration is in `db` folder.

## Run Application

### Run web server

```bash
go run cmd/web/main.go
```