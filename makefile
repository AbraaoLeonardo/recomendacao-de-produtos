reset-docker:
	docker container kill postgresql
	docker container rm postgresql

start-docker:
	docker container run \
			-p 9999:5432 \
			-v data:/var/lib/postgresql \
			-e POSTGRES_PASSWORD=kzzyy2 \
			-e POSTGRES_USER=kzzyy2 \
			-e POSTGRES_DB=recomendacoes \
			--name postgresql \
			-d postgres:14.19-alpine3.21