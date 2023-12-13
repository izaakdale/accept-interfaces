### Repo for article on accepting interfaces and returning structs

To run:

1. go test ./...
2. docker run -d -p 6379:6379 --name redis redis:latest
3. HOST=localhost PORT=8080 REDIS_URL=redis://localhost:6379 go run .
4. curl localhost:8080/ping
