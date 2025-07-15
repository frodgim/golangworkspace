# Ticket API Documentation

## Overview
This service provides a RESTful API for managing tickets, using Go (Gin), MySQL, and Redis. It supports full CRUD operations.

## Entity: Ticket
- **ID**: int, primary key, auto-increment
- **Name**: string, max 50 chars
- **Type**: enum, one of `kindA`, `kindB`, `kindC`

## Endpoints

### Create Ticket
- **POST** `/tickets`
- **Body**: `{ "name": "string", "type": "kindA|kindB|kindC" }`
- **Response**: 201 Created, returns created ticket

### Get Ticket
- **GET** `/tickets/{id}`
- **Response**: 200 OK, returns ticket or 404 if not found

### Update Ticket
- **PUT** `/tickets/{id}`
- **Body**: `{ "name": "string", "type": "kindA|kindB|kindC" }`
- **Response**: 200 OK, returns updated ticket

### Delete Ticket
- **DELETE** `/tickets/{id}`
- **Response**: 204 No Content

### List Tickets
- **GET** `/tickets`
- **Response**: 200 OK, returns array of tickets

## Running with Docker Compose
1. Build and start all services:
   ```sh
   docker-compose up --build
   ```
2. The API will be available at `http://localhost:8080`.

## Testing
- Run unit tests:
  ```sh
  go test ./...
  ```

## Notes
- MySQL and Redis are required for full functionality.
- The API validates ticket type and name length.

## Example Usage with curl

### Create Ticket
```sh
curl -X POST http://localhost:8080/tickets \
  -H "Content-Type: application/json" \
  -d '{"name": "Sample Ticket", "type": "kindA"}'
```

### Get Ticket
```sh
curl http://localhost:8080/tickets/1
```

### Update Ticket
```sh
curl -X PUT http://localhost:8080/tickets/1 \
  -H "Content-Type: application/json" \
  -d '{"name": "Updated Ticket", "type": "kindB"}'
```

### Delete Ticket
```sh
curl -X DELETE http://localhost:8080/tickets/1
```

### List Tickets
```sh
curl http://localhost:8080/tickets
```

## Caching Strategy

The API uses Redis to cache the 50 most frequently accessed, created, or updated tickets:
- Each ticket is cached in Redis with a key like `ticket:{id}`.
- A Redis sorted set (`ticket:freq`) tracks how often each ticket is accessed or modified.
- On every GET, POST (create), or PUT (update) operation, the ticket is cached and its frequency is incremented.
- Only the top 50 most frequent ticket IDs are kept in the cache; others are evicted automatically.
- This improves performance for the most commonly used tickets while keeping memory usage predictable.
