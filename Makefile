init:
	docker volume create postgres_pastely && \
	docker run --name postgres-pastely -e POSTGRES_PASSWORD=admin -d \
	-v postgres_pastely:/var/lib/postgresql/data -p 5432:5432 \
	postgres:16
