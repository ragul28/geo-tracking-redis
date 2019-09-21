build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ./cmd/geo-tracking-redis
	docker build -t geo-tracking-redis .

run:
	go build && ./geo-tracking-redis

init:
	GO111MODULE=on go mod init github.com/ragul28/geo-tracking-redis
	GO111MODULE=on go get -u

mod:
	GO111MODULE=on go mod tidy
	GO111MODULE=on go mod verify
	GO111MODULE=on go mod vendor

redis-docker:
	docker run --rm -d -p 6379:6379 redis:alpine
