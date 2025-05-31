# Quote Manager

A simple REST API service for managing quotes written in Go.

## Features

- Create new quotes
- Get all quotes
- Get quotes by author
- Get random quote
- Delete quote by ID
- Persistent storage option
- Configurable logging levels

## Installation

```bash
git clone https://github.com/TheTeemka/Task_QuoteManager.git
cd Task_QuoteManager
go mod download
```

## Running the Service

### Using Make

Development mode with debug logging and persistence:
```bash
make dev
```

Production mode:
```bash
make run
```

### Using Go directly

```bash
go run cmd/api/main.go [options]
```

Available options:
- `-port` - Server port (default ":8080")
- `-file` - File path for persistent storage (default "temirlan_bayangazy_data.json")
- `-logLevel` - Logging level: debug, info, warn, error (default "info")
- `-bePersistent` - Enable persistent storage (default false)

## API Examples

### Create a new quote
```bash
curl -X POST http://localhost:8080/quotes \
  -H "Content-Type: application/json" \
  -d '{"author":"Confucius", "quote":"Life is simple, but we insist on making it complicated."}'
```

### Get all quotes
```bash
curl http://localhost:8080/quotes
```

### Get quotes by author
```bash
curl http://localhost:8080/quotes?author=Confucius
```

### Get random quote
```bash
curl http://localhost:8080/quotes/random
```

### Delete quote by ID
```bash
curl -X DELETE http://localhost:8080/quotes/1
```

## API Response Examples

### Successful quote creation
```json
{
  "id": 1,
  "author": "Confucius",
  "quote": "Life is simple, but we insist on making it complicated."
}
```

### Get all quotes response
```json
{
  "quotes": [
    {
      "id": 1,
      "author": "Confucius",
      "quote": "Life is simple, but we insist on making it complicated."
    }
  ]
}
```


## Development

Running tests:
```bash
make test
```

## License

MIT License