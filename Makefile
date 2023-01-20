.PHONY: test start stop 

test:
	go test -v -cover ./...

start:
	docker-compose up -d	

stop:
	docker-compose down 