.PHONY: run
run:
	docker compose -f ./build/docker-compose.yml rm && \
	docker compose -f ./build/docker-compose.yml build --no-cache && \
 	docker compose -f ./build/docker-compose.yml up -d


.PHONY: stop
stop:
	docker compose -f ./build/docker-compose.yml down


.PHONY: kafka
kafka:
	docker exec -it kafka-server bash


.PHONY: logs
logs:
	docker logs init-kafka

