# Database Migration

Install go migrate

```
    brew install golang-migrate
```

Document : https://github.com/golang-migrate/migrate/blob/master/README.md

### Command example

create new database sequence

```
    migrate create -ext sql -dir . -seq file_name
```

up version database

```
    migrate -source file://. -database "postgres://migrator:$PASSWORD@localhost:5432/shipment?sslmode=disable" up
```

down version database 1 version

```
    migrate -source file://. -database "postgres://migrator:$PASSWORD@localhost:5432/shipment?sslmode=disable" down 1
```

Refer : https://www.connectionstrings.com/postgresql/
