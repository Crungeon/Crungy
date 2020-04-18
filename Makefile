VER=$(shell cat VERSION)

dev:
	GO111MODULE=on go run main.go

test:
	go test ./... -v

build:
	go build -o bin/crungy

container:
	docker build -t crungeon/crungy:$(VER) .

run-container:
	docker run --env-file .env crungeon/crungy:$(VER)

deploy: test container
	docker push crungeon/crungy