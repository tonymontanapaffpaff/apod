.PHONY: mod run up build_up down create build start stop logs

mod:
	GO111MODULE=on go mod tidy

run: build create start

up:
	sudo docker-compose -f docker-compose.yml up

build_up:
	sudo docker-compose -f docker-compose.yml up --build

down:
	sudo docker-compose -f docker-compose.yml down

create:
	sudo docker-compose -f docker-compose.yml create

build:
	sudo docker-compose -f docker-compose.yml build

start:
	sudo docker-compose -f docker-compose.yml start

stop:
	sudo docker-compose -f docker-compose.yml stop

logs:
	sudo docker-compose -f docker-compose.yml logs
