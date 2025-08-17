# json-go

JSON Parser ported from TypeScript to Go 1.25.0

## License

MIT

## Reference

[json.org](http://json.org)

## Common Go Commands

### Build the project

```sh
go build ./...
```

## Lint code

To lint your Go code for common issues, use [golangci-lint](https://github.com/golangci/golangci-lint):

```sh
golangci-lint run ./...
```

### Format code

```sh
go fmt ./...
```

### Run tests

```sh
go test ./...
```

### Test Coverage

To measure test coverage:

```sh
go test -coverprofile="coverage.out" ./...
go tool cover -html="coverage.out" -o="coverage.html"
```

### Run the CLI

```sh
go run cmd/cli/main.go
```

### Run the Echo API server

To start the Echo API server:

```sh
go run cmd/api-echo/main.go
```

### Run the Gin API server

To start the Gin API server:

```sh
go run cmd/api-gin/main.go
```

## Test API with .http Requests

You can test either Gin and Echo API endpoints using the [REST Client](https://marketplace.visualstudio.com/items?itemName=humao.rest-client) extension for VSCode.

1. Install the REST Client extension.
2. Open any `.rest` file in `testdata/`.
3. Click "Send Request" above the desired request to test the API.
