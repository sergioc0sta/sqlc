createMigrations:
	migrate create -ext=sql -dir=sql/migrations -seq init

migrate-up:
	migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/courses" -verbose up
migrate-down:
	migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/courses" -verbose down
generate: 
	sqlc generate

.PHONY: createMigrations migrate-down migrate-up

