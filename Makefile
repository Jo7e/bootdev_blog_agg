include .env

gen:
	sqlc generate

sqlrun:
	sudo service postgresql start

sqlcon:
	sudo -u postgres psql

mig:
	goose -dir ./sql/schema postgres $(DB_URL) up

run:
	go run .