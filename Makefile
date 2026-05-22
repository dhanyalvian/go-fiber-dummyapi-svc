# Makefile
migrate:
	go run cmd/api/main.go --migrate

run:
	go run cmd/api/main.go

build:
	go build -o monitoring-proyek-svc ./cmd/api/main.go

dev:
	air # Jika menggunakan live reload Air

docker-build:
	docker build --no-cache -t monitoring-proyek-svc .

docker-run:
	docker run -d \
	--name monitoring-proyek-svc \
	--network dna_private_network \
	--restart on-failure \
	--env-file .env \
	-p 8082:8082 \
	monitoring-proyek-svc

docker-stop:
	docker stop monitoring-proyek-svc

docker-rm:
	docker rm monitoring-proyek-svc
	
docker-create-network:
	docker network create --driver bridge dna_private_network