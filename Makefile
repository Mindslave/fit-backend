postgres:
	docker run --name fit-backend-db -e POSTGRES_PASSWORD=thisissupersecret -e POSTGRES_USER=fituser -e POSTGRES_DB=fit_db -p 5432:5432 -d postgres

createdb:
	docker exec -it fit-backend-db createdb --username=fituser --owner=fituser fit_db

dropdb:
	docker exec -it fit-backend-db dropdb --username=fituser fit_db

migrateup:
	migrate -path db/migration -database "postgresql://fituser:thisissupersecret@localhost:5432/fit_db?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://fituser:thisissupersecret@localhost:5432/fit_db?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown sqlc