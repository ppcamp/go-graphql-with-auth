# go-graphql-with-auth


This repository will contain a simple (and yet maintanable) project structure
for using in a graphql server.

This repository also have a full tutorial in my portfolio. Take a look at 
https://ppcamp.github.io/post/graphql-go/


## How to run
go


## Migrations

Install the [migration cli](https://github.com/golang-migrate/migrate/releases/tag/v4.15.1)

```bash
# Create
migrate create -ext sql -dir src/internal/repository/migrations -seq create_users_table

# Export pgurl
export POSTGRES_URL=postgres://gouser:gopsswd@localhost:5432/gousers?sslmode=disable

# Migrate
migrate -database $POSTGRES_URL -path src/internal/repository/migrations/ up

# Fallback
migrate -database $POSTGRES_URL -path src/internal/repository/migrations/ up
```

## Packages

- [sqlx](https://jmoiron.github.io/sqlx/)
