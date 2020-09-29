SERVER := app_srv_1
DATABASE := app_db_1
NETWORK := app_net
VOLUME := app_vol

.PHONY: clean
clean: stop
	docker rm $(SERVER)
	docker rm $(DATABASE)
	docker volume rm $(VOLUME)
	docker network rm $(NETWORK)

.PHONY: build
build: _load-env
	env APP_PORT=${APP_PORT} DATABASE_PORT=${DATABASE_PORT} \
	docker-compose -p app -f docker/compose.yaml build

.PHONY: up
up: _load-env
	env APP_PORT=${APP_PORT} DATABASE_PORT=${DATABASE_PORT} \
	docker-compose -p app -f docker/compose.yaml up

.PHONY: down
down: _load-env
	env APP_PORT=${APP_PORT} DATABASE_PORT=${DATABASE_PORT} \
	docker-compose -p app -f docker/compose.yaml down

.PHONY: start
start: start-server start-database

.PHONY: stop
stop: stop-server stop-database

.PHONY: start-server
start-server: _load-env
	env APP_PORT=${APP_PORT} DATABASE_PORT=${DATABASE_PORT} \
	docker-compose -p app -f docker/compose.yaml start srv

.PHONY: start-database
start-database: _load-env
	env APP_PORT=${APP_PORT} DATABASE_PORT=${DATABASE_PORT} \
	docker-compose -p app -f docker/compose.yaml start db

.PHONY: restart-server
restart-server: _load-env
	env APP_PORT=${APP_PORT} DATABASE_PORT=${DATABASE_PORT} \
	docker-compose -p app -f docker/compose.yaml restart srv

.PHONY: restart-database
restart-database: _load-env
	env APP_PORT=${APP_PORT} DATABASE_PORT=${DATABASE_PORT} \
	docker-compose -p app -f docker/compose.yaml restart db

.PHONY: stop-server
stop-server: _load-env
	env APP_PORT=${APP_PORT} DATABASE_PORT=${DATABASE_PORT} \
	docker-compose -p app -f docker/compose.yaml stop srv

.PHONY: stop-database
stop-database: _load-env
	env APP_PORT=${APP_PORT} DATABASE_PORT=${DATABASE_PORT} \
	docker-compose -p app -f docker/compose.yaml stop db

_load-env:
	$(eval include env/app-db.env)
	$(eval include env/app-srv.env)

default: up
