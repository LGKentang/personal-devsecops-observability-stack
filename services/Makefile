.PHONY: build build-all run docker-build docker-up
.PHONY: migrate-up migrate-down

build:
	cd coffee-catalog && go build -o bin/coffee-catalog .
	cd ../coffee-orders && go build -o bin/coffee-orders .
	cd ../gateway && go build -o bin/gateway .

build-all: build

run:
	# run services locally (three shells recommended)
	@echo "Run each service in its own terminal:"
	@echo "cd services/coffee-catalog && go run ."
	@echo "cd services/coffee-orders && COFFEE_CATALOG_URL=http://localhost:8081 go run ."
	@echo "cd services/gateway && go run ."

docker-build:
	docker-compose build

docker-up:
	docker-compose up --build

migrate-up:
	# start the DB (detached)
	docker-compose up -d db
	# run the migrate service which applies all up migrations
	docker-compose run --rm migrate

migrate-down:
	# roll back all migrations (be careful, this will drop data if down scripts remove tables)
	docker-compose run --rm migrate -command "-path=/migrations" "-database" "postgres://postgres:postgres@db:5432/coffee?sslmode=disable" down
