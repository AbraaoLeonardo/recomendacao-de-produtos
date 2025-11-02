stop-docker:
	docker container kill recomendacao
	docker container rm recomendacao
	docker image rm recomendacao

start-docker:
	docker container run \
			-p $(DB_PORT):5432 \
			-v data:/var/lib/postgresql \
			-e POSTGRES_PASSWORD=$(DB_PASSWORD) \
			-e POSTGRES_USER==$(DB_USER) \
			-e POSTGRES_DB==$(DB_NAME) \
			--name postgresql \
			-d postgres:14.19-alpine3.21

build-docker:
	docker build -t abraao/recomendacao:0.0.1 .
	docker container run --name recomendacao -p 8080:8080 -d abraao/recomendacao:0.0.1