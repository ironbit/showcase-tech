# Showcase-Tech

## Technologies

* [Golang - Beego - Postgres](./golang-beego-postgres)

## Golang - Beego - Postgres
It's a Rest API implementation using Golang (language programming), Beego (Framework), and Postgres (Database).

It's implemented a simple Rest API for a Book and Author entities, known as many to many relationships, which uses as Rest API as the interface and the information is stored in a database using Postgres.

The requests are processed using logging and metrics. For that purpose, it is used the decorator pattern (in the golang community, it's known as middleware).

The packages used are:
* github.com/astaxie/beego
* github.com/go-kit/kit
* github.com/lib/pq
* github.com/satori/go.uuid

Note: It's still under development.
