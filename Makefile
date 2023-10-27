build:
	docker-compose build 

run:
	docker-compose up -d --build

stop:
	docker-compose stop 
test:
	go test -v ./...
