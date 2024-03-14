postgres:
    docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres
createdb:
    docker exec -it postgres createdb --username=root --owner=root assignment2
dropdb:
    docker exec -it postgres dropdb assignment2
migrateup:
    migrate -path database/migration -database "postgresql://root:secret@localhost:5432/assignment2?sslmode=disable" -verbose up
migratedown:
    migrate -path database/migration -database "postgresql://root:secret@localhost:5432/assignment2?sslmode=disable" -verbose down   
entimport:
    go run ariga.io/entimport/cmd/entimport -dsn "postgres://root:secret@localhost:5432/assignment2?sslmode=disable"
generate:
    go generate ./ent
initent:
    go run -mod=mod entgo.io/ent/cmd/ent new
.PHONY: postgres createdb dropdb migrateup migratedown