.PHONY: run
run:
	docker-compose -f ./build/docker-compose.yml rm && \
	docker-compose -f ./build/docker-compose.yml build --no-cache && \
 	docker-compose -f ./build/docker-compose.yml up