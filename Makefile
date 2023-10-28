build:
	docker-compose build 

run:
	sudo docker-compose up --build

stop:
	docker-compose stop 
test:
	go test -v ./...
