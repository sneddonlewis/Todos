### Todo App

## Description

Web Application to store and retrieve todo items

## Dependencies

Docker & Docker Compose
Go @ 1.18
Nodejs @ 16.15
npm

PostgreSQL @ 9.6.21 (image pulled from docker)

## Get started

run:  

`docker-compose -f docker-compose.dev.yml up --build`

Run the web API:

```
cd server
go run .
```

Run the client app:

```
cd webapp
npm started
```

The development client app will be served at `localhost:4200`

## Build for production

`docker-compose up`

