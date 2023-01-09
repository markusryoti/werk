Run migrations

Up
```bash
migrate -path db/migrations/ -database postgres://devuser:dev123@localhost:5432/werk?sslmode=disable up
```

Down
```bash
migrate -path db/migrations/ -database postgres://devuser:dev123@localhost:5432/werk?sslmode=disable down
```
