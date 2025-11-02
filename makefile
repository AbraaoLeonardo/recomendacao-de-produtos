reset-docker:
	docker container kill postgresql
	docker container rm postgresql

start-docker:
	docker container run \
			-p $(DB_PORT):5432 \
			-v data:/var/lib/postgresql \
			-e POSTGRES_PASSWORD=$(DB_PASSWORD) \
			-e POSTGRES_USER==$(DB_USER) \
			-e POSTGRES_DB==$(DB_NAME) \
			--name postgresql \
			-d postgres:14.19-alpine3.21