# TicketCLI Documentation

## Overview
TicketCLI is a command-line client for interacting with the Ticket API. It allows you to create, retrieve, update, delete, and list tickets directly from your terminal.

## Installation
1. Build the CLI:
   ```sh
   go build -o ticketcli main.go
   ```
2. (Optional) Move the binary to a directory in your PATH for global use.

## Configuration
- By default, TicketCLI connects to the API at `http://localhost:8080`.
- You can specify a different API base URL using the `--api` flag:
  ```sh
  ./ticketcli --api http://your-api-host:8080 <command>
  ```

## Commands

### Create a Ticket
```sh
./ticketcli create --name "My Ticket" --type kindA
```

### Get a Ticket by ID
```sh
./ticketcli get 1
```

### Update a Ticket
```sh
./ticketcli update 1 --name "Updated Name" --type kindB
```

### Delete a Ticket
```sh
./ticketcli delete 1
```

### List All Tickets
```sh
./ticketcli list
```

## Ticket Types
- Valid types: `kindA`, `kindB`, `kindC`

## Examples
- Create and then get a ticket:
  ```sh
  ./ticketcli create --name "Test" --type kindA
  ./ticketcli get 1
  ```

## Notes
- All commands return JSON responses from the API.
- Ensure the Ticket API server is running and accessible before using the CLI.
