\*\*\*\*# GLIP

Global Logistics Integration Platform

---

This project uses **node.js** (also pnpm) and **golang**, please make sure you have those two already installed

This project also use Docker for both dev and prod environment.

Before proceeding please check your .env file

### Setup 3rd party services

Before proceeding you'll need **golang-migrate** to migrate the database and **sqlc** to generate query files. Both are available on Brew.

- run `pnpm compose`
- run `go run apps/server/registry-api/cmd/scripts/setup-etcd/main.go` to setup etcd credentials
- run `migrate -source file://. -database "postgres://USER:PASSW@localhost:DB_PORT/DB_NAME?sslmode=disable" up` on both shipment api and registry api db to migrate

### Running backends

- run `go run apps/server/registry-api/cmd/scripts/setup-etcd/main.go` to start up registry api
- run `go run apps/server/shipment-api/cmd/main.go` to start up shipment api
