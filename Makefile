.PHONY: test start stop build run migrateup migratedown sqlc 

test:
	go test -v -cover ./...

start:
	docker-compose up -d	

stop:
	docker-compose down

build:
	docker build --tag instituto-maternar .	

run: 
	docker run instituto-maternar

migrateup:
	migrate -path db/migration -database "postgresql://admin:admin@0.0.0.0:5432/appdb?sslmode=disable" -verbose up	

migratedown:
	migrate -path db/migration -database "postgresql://admin:admin@0.0.0.0:5432/appdb?sslmode=disable" -verbose down 

sqlc:
	sqlc generate

