### Repo for article on accepting interfaces and returning structs

To run:
docker run -d -p 6379:6379 --name redis redis:latest
HOST=localhost PORT=8080 REDIS_URL=redis://localhost:6379 go run .
curl localhost:8080/ping
