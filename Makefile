# Makefile
migrate:
	go run cmd/api/main.go --migrate

run:
	go run cmd/api/main.go

build:
	go build -o dummyapi-svc ./cmd/api/main.go

dev:
	air # Jika menggunakan live reload Air

docker-build:
	docker build --no-cache -t dummyapi-svc .

docker-run:
	docker run -d \
	--name dummyapi-svc \
	--network dna_private_network \
	--restart on-failure \
	--env-file .env \
	-p 8084:8084 \
	dummyapi-svc

docker-stop:
	docker stop dummyapi-svc

docker-rm:
	docker rm dummyapi-svc
	
docker-create-network:
	docker network create --driver bridge dna_private_network