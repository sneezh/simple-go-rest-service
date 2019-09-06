# Simple GO REST API service

## Installation

### Without Docker

Open service directory and install dependencies.

```
go get -d -v ./...
go get github.com/stretchr/testify/assert
go get github.com/golangci/golangci-lint/cmd/golangci-lint
```

Copy default environment file to .env.
```
cp .env.example .env
```

Build and run app
```bash
go build
./simple-go-rest-service
```

### With docker-compose

Copy default docker-compose file to docker-compose.yml.

```
cp docker-compose.example.yml docker-compose.yml
```

Change docker-compose if you need (it includes some additional services like redis, postgres etc. by default).

Copy default environment file to .env.

```
cp .env.example .env
```

Set up required env variables in .env file.

Start your application.

``` 
docker-compose up -d
```

## Create table for entity
```sql
create table entity
(
	id serial not null,
	name varchar,
	created_at timestamp default CURRENT_TIMESTAMP not null,
	updated_at timestamp default CURRENT_TIMESTAMP not null,
	deleted_at timestamp
);

create unique index entity_id_uindex
	on entity (id);

create unique index entity_name_uindex
	on entity (name);

alter table entity
	add constraint entity_pk
		primary key (id);
```