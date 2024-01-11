createdb:
	docker exec -it postgres16 createdb --username=root --owner=root email_otp

dropdb:
	docker exec -it postgres16 dropdb email_otp

migratecreate:
	migrate create -ext sql -dir db/migrations -seq $(name)

migrateup:
	migrate -path db/migrations -database "postgresql://root:root@localhost:5432/email_otp?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migrations -database "postgresql://root:root@localhost:5432/email_otp?sslmode=disable" -verbose down 1

migratefix:
	migrate -path db/migrations -database "postgresql://root:root@localhost:5432/email_otp?sslmode=disable" -verbose force $(version)

test:
	go test -v -cover -short ./...

PHONY: postgres createdb dropdb migratecreate migrateup migratedown migratefix test