include .env

gen:
	sqlc generate

sqlrun:
	sudo service postgresql start

sqlcon:
	PGPASSWORD=$(DB_PASS) psql -U postgres -d $(DB_NAME)

mig:
	goose -dir ./sql/schema postgres $(DB_URL) up

run:
	go run .