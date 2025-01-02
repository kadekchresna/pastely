

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
	docker run -p 6379:6379 \
	--name redis-pastely -d \
	-v redis_pastely:/data \
	redis

cassandra:
	docker volume create cassandra_pastely && \
	docker run -p 7000:7000 \
	-p 9042:9042 \
  	-p 7001:7001 \
	--name cassandra-pastely -d \
	-v cassandra_pastely:/var/lib/cassandra \
  	-v ./deploy/cassandra.yaml:/etc/cassandra/cassandra.yaml \
	cassandra:latest

timescale:
	docker volume create timescale_pastely && \
	docker run -d --name timescaledb-pastely -p 6543:5432 \
	-e POSTGRES_PASSWORD=admin \
    -e POSTGRES_USER=postgres \
	-v timescale_pastely:/var/lib/postgresql/data \
	timescale/timescaledb:latest-pg16

init: pg redis
	@echo "init...."

docker-tag:
	docker build -t  kadekchresna/pastely:latest .

docker-push: docker-tag
	docker push kadekchresna/pastely:latest

minio-upload:
	mc alias set minio-docker http://localhost:9000 root rootatleast8 
	mc cp env.example minio-docker/pastely-bucket/env.file

# Migrate
# db (analytic, main)
# name (migration name) ex: make migrate-create db=main name=create_table_user
migrate-create:
	@if [ "$(db)" = "analytic" ]; then \
		migrate create -ext sql -dir ./migration/analytic $(name); \
	else \
		migrate create -ext sql -dir ./migration/main $(name); \
	fi

migrate-up:
	go run migration/migrate.go $(db) up

migrate-down:	
	go run migration/migrate.go $(db) down


pastely-up:
	helm upgrade --install postgres-operator ./deploy/chart/postgres-operator
	sleep 20
	helm upgrade --install postgres-pastely ./deploy/chart/postgres-pastely
	helm upgrade --install minio-pastely ./deploy/chart/minio-pastely
	helm upgrade --install pastely ./deploy/chart/pastely
