postgres:
    docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres
createdb:
    docker exec -it postgres createdb --username=root --owner=root postgres
dropdb:
    docker exec -it postgres dropdb postgres
migrateup:
    migrate -path database/migration -database "postgresql://root:secret@localhost:5432/postgres?sslmode=disable" -verbose up
migratedown:
    migrate -path database/migration -database "postgresql://root:secret@localhost:5432/postgres?sslmode=disable" -verbose down
creatschema:
    go run -mod=mod entgo.io/ent/cmd/ent new
entimport:
    go run ariga.io/entimport/cmd/entimport -dsn "postgres://root:secret@localhost:5432/postgres?sslmode=disable"
generate:
    go generate ./ent
.PHONY: postgres createdb dropdb migrateup migratedown