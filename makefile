include .env
export

MIGRATIONS_PATH=./db/migrations

migration-create:
	migrate create -ext sql -dir $(MIGRATIONS_PATH) -seq $(filename)
migration-up:
	migrate -database ${DATABASE_DSN} -path $(MIGRATIONS_PATH) up
migration-down:
	migrate -database ${DATABASE_DSN} -path $(MIGRATIONS_PATH) down
migration-force:
	migrate -path $(MIGRATIONS_PATH) -database ${DATABASE_DSN} force $(version)

