# Showcase Tech

## Technologies

* [Golang - Beego - Postgres](./golang-beego-postgres)

## Golang - Beego - Postgres
It's a Rest API implementation using Golang (language programming), Beego (Framework), and Postgres (Database). It uses a Rest API as the interface, and the information is stored in a database using Postgres.

The entities, Books and Authors, are chosen for the implementation because it's many to many relationships.

The requests are processed using logging and metrics. For that purpose, it is used the decorator pattern (in the golang community, it's known as middleware).

The packages used are:
* github.com/astaxie/beego
* github.com/go-kit/kit
* github.com/lib/pq
* github.com/satori/go.uuid

Note: It's still under development.
