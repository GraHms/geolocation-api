postgres:
	docker run --name pg_container -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it pg_container createdb --username=postgres --owner=postgres geoloationsdb

dropdb:
	docker exec -it pg_container dropdb --username=postgres  geoloationsdb

migrate:
	migrate -path db/migrations/ -database "postgresql://postgres:west04@localhost:5432/geoloationsdb?sslmode=disable" -verbose up

sqlc:
	sqlc generate

test:
	go test -v -cover ./...


.PHONY: createdb dropdb postgres migrate sqlc test