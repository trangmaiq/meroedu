[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=meroedu_meroedu&metric=alert_status)](https://sonarcloud.io/dashboard?id=meroedu_meroedu)
![Go](https://github.com/meroedu/meroedu/workflows/Go/badge.svg?branch=master)
[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=meroedu_meroedu&metric=coverage)](https://sonarcloud.io/dashboard?id=meroedu_meroedu)
[![Go Report Card](https://goreportcard.com/badge/github.com/meroedu/meroedu)](https://goreportcard.com/report/github.com/meroedu/meroedu)
# Introduction
Mero Edu is a software application for the administration, documentation, tracking, reporting, automation and delivery of educational courses, training programs, or learning and development programs.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

[![alt text](screenshots/meroedu.png)](https://meroedu.com)

### Prerequisites

1. NodeJS       => Frontend
2. Golang       => Backend
3. Docker(optional)
4. Make

## Development Stack
1. Golang
2. ReactJS
3. Mariadb(Bitnami galera cluster)

### Running & Initializing database
`make prepare`

## Running the tests

`make test`
`make test-coverage`

## Start/Stop Backend Application

`make run`
`make stop`

## Reset Backend Application

`make reset`

## Start/Stop Frontend Application
```
cd ui/
yarn add react-scripts
yarn start
(CTRL+C) to stop
```

## Code Folder Structure

```
domain // Entity
├── course.go
├── category.go
└── author.go 

course
├── delivery
│   └── http
│       ├── course_handler.go
│       └── course_test.go
├── mocks
│   ├── courseRepository.go
│   └── courseUsecase.go
├── repository // implementation
│   ├── mysql_course.go
│   └── mysql_course_test.go
├── repository.go 
├── usecase // Implementation
│   ├── courseu_usecase_test.go
│   └── course_usecase.go
└── usecase.go // Usecase Interface.
```


## License
[View License](https://github.com/meroedu/meroedu/blob/master/LICENSE)
