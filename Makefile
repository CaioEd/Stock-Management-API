install:
	go mod tidy

database:
	docker-compose up -d

migrations:
	migrate create -ext sql -dir migrations -seq create_users_table

run:
	go run main.go