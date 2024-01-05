DB_DOCKER_CONTAINER=goapi_db
BINARY_NAME=goapi

postgres:
	docker run --name ${DB_DOCKER_CONTAINER} -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=123qwe -d postgres:12-alpine

createdb:
	docker exec -it ${DB_DOCKER_CONTAINER} createdb --username=root --owner=root go-api

build:
	@go build -o bin/go-api

run: build
	@./bin/go-api

test:
	@go test -v ./
