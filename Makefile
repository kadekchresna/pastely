

ENVS := CGO_ENABLED=0 GOOS=linux GOARCH=amd64
FLAGS := -a 

tidy:
	go mod tidy


build: tidy
	mkdir -p build/
	$(ENVS) go build ${FLAGS} -o $(BIN) build/.


pg:
	docker volume create postgres_pastely && \
	docker run --name postgres-pastely -e POSTGRES_PASSWORD=admin -d \
	-v postgres_pastely:/var/lib/postgresql/data -p 5432:5432 \
	postgres:16

minio:
	docker volume create minio_pastely
	docker run \
		-p 9000:9000 \
		-p 9001:9001 \
		--name minio-pastely \
		-v minio_pastely:/data \
		-e "MINIO_ROOT_USER=root" \
		-e "MINIO_ROOT_PASSWORD=rootatleast8" -d \
		quay.io/minio/minio server /data --console-address ":9001"

redis:
	docker volume create redis_pastely && \
	docker run --name redis-pastely -d \
	-v redis_pastely:/data \
	redis

init: pg redis
	@echo "init...."

docker-tag:
	docker build -t  kadekchresna/pastely:latest .

docker-push: docker-tag
	docker push kadekchresna/pastely:latest

minio-upload:
	mc alias set minio-docker http://localhost:9000 root rootatleast8 
	mc cp env.example minio-docker/pastely-bucket/env.file
