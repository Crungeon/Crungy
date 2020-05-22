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

pi-deploy:
	git checkout master
	git pull
	COMPOSE_DOCKER_CLI_BUILD=1 DOCKER_BUILDKIT=1 docker-compose build
	docker-compose run -d crungy

deploy: test container
	docker push crungeon/crungy