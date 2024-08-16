## BBB API

api to count votes, simulating bbb architecture.

![alt text](<public/api big brother.png>)

## Install Go

https://go.dev/doc/install

## Running

1. If you don't have it on your machine, run the postgres and rabbitmq containers

```bash
docker run -d --name rabbitmq -p 5672:5672 -p 15672:15672 -e RABBITMQ_DEFAULT_USER=guest -e RABBITMQ_DEFAULT_PASS=guest rabbitmq:3-management

docker run --name postgres -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -p 5432:5432 -d postgres:latest
```

2. Create the database by accessing the postgres container

```bash
docker exec -it postgres bash

psql -U postgres

CREATE DATABASE big_brother;
```

3. Create a network and add the postgres and rabbitmq containers

```bash
docker network create bbb

docker network connect bbb postgres

docker network connect bbb rabbitmq
```

to find out the IP of each container use

```bash
docker inspect -f '{{.NetworkSettings.IPAddress}}' postgres

docker inspect -f '{{.NetworkSettings.IPAddress}}' rabbitmq
```
This information will be important to put in the .env in the HOST variable (postgres) and RABBITMQ_URL

4. Create .env file (use .env.example as a reference)

5. Start api

```bash
go run bbb/server
```

6. Request example

```bash
curl -d '{"candidate": "Linus Torvalds"}' -X POST http://localhost:8080/api/votos

curl -d '{"candidate": "Steve Jobs"}' -X POST http://localhost:8080/api/votos
```

## [RabbitMQ](http://localhost:15672)
