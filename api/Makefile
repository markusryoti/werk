run:
	export GOOGLE_APPLICATION_CREDENTIALS="./secrets/serviceAccountKey.json" && \
		go run -race cmd/rest/main.go

docker-run:
	sudo docker-compose up --build

migrate-up:
	migrate -path db/migrations/ -database postgres://devuser:dev123@localhost:5432/werk?sslmode=disable up

migrate-down:
	migrate -path db/migrations/ -database postgres://devuser:dev123@localhost:5432/werk?sslmode=disable down
