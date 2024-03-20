image:
	docker images

container:
	docker ps -a

compose:
	docker-compose up -d

compose-ui:
	docker-compose -f docker-compose-kafka-ui.yml up -d

exec:
	winpty docker container exec -it kafka bash

producer:
	go run main.go

consumer1:
	go run .\data\main.go

consumer2:
	go run .\process\main.go

.PHONY: image container compose compose-ui exec producer consumer1