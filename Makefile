build_docker:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ./geo-tracking-redis
	docker build -t geo-tracking-redis .

build:
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w"

run:
	go build && ./geo-tracking-redis

init:
	go mod init github.com/ragul28/geo-tracking-redis
	go get -u

mod:
	go mod tidy
	go mod verify
	go mod vendor

redis-docker:
	docker run --rm -d -p 6379:6379 redis:alpine
