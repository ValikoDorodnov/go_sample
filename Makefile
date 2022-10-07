docker-up:
	docker-compose up -d --build

docker-down:
	docker-compose down -v --remove-orphans

docker-up-db:
	docker-compose up -d --build postgres

run-test:
	go test ./...