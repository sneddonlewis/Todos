# Todo App

## Description

Web Application to store and retrieve todo items

## Dependencies

Docker & Docker Compose  
Go @ 1.18  
Nodejs @ 16.15  
npm  

PostgreSQL @ 9.6.21 (image pulled from docker)

## Get started developing

Pull a database container image and populate the db:  

```
docker-compose -f docker-compose.dev.yml up --build
```

Run the web API:

```
cd server
go run .
```

Run the client app:

```
cd webapp
npm start
```

The development client app will be served at `localhost:8080`

## Build for production

```
docker-compose up
```

