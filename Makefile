.PHONY: run
run:
	docker compose -f ./docker-compose.yml rm && \
	docker compose -f ./docker-compose.yml build --no-cache && \
 	docker compose -f ./docker-compose.yml up -d


.PHONY: stop
stop:
	docker compose -f ./docker-compose.yml down


.PHONY: kafka
kafka:
	docker exec -it kafka-server bash


.PHONY: logs-api
logs-api:
	docker logs api


.PHONY: logs-init
logs-init:
	docker logs init-kafka


.PHONY: logs-consumer
logs-consumer:
	docker logs consumer


.PHONY: logs-ch
logs-ch:
	docker logs ch-db

