# Golang - Beego - Postgres

## Develop
1. Generate the Rest API application called **app** using the **bee** script.
```bash
# in 'golang-beego-postgres' folder
$ bee api app
```

2. Generate the module files for the application (GOPATH and GO111MODULE are configured).
```bash
# in 'app' folder
$ go mod init
$ go build
```

## Structure
* **app**

  Contains the web rest api in golang using beego framework.

* **docker**

  Contains the docker files.
  The compose.yaml file for docker-compose.
  The srv.dockerfile for the web rest api application.

* **env**

  Contains the environmental variables for the server and the database.
  app-db.env is for the database (docker deploy).
  app-srv.env is for the application (docker deploy).
  app-db-dev.env is for the database (develop - local).
  app-srv-dev.env is for the application (develop - local).

* **sql**

  Contains the files for the initial configuration of the database.
  10-database.sh is for the creation of the user and the database; and the permissions.
  20-tables.sh is for the creation of the tables and the sequences in the database.

## Deploy
### Using Docker
```bash
# in 'golang-beego-postgres' folder
$ make up
```

The docker compose will create two containers (app_srv_1 & app_db_1), a network (app_net) and a volume (app_vol).

The Makefile targets:
* **clean**: stop the containers, erase containers, volume and network.
* **build**: build the application (only the srv service).
* **up**: (default target) build the application, create network and volume resources, and start the services.
* **up-server**: build the application, create network and volume resources, and start the both containers.
* **up-database**: create network and volume resources, and start the database container.
* **down**: stop the services, and remove services and network resources.
* **start**: start both services.
* **stop**: stop both services.
* **start-server**: start only the server container.
* **start-database**: start only the database container.
* **restart-server**: restart only the server container.
* **restart-database**: restart only the database container.
* **stop-server**: stop only the server container.
* **stop-database**: stop only the database container.

#### Note: Clear Resouces in Docker
Erase the image created by the composer and erase all resources used by this application. Be sure to use these commands due to it can erase other resources.
```bash
# in 'golang-beego-postgres' folder
$ make clean
$ docker rmi app-srv
$ docker image prune
```

### Using Local
#### Execute Server Application
1. Load environment variables
```bash
# in 'golang-beego-postgres' folder
$ export $(xargs < env/app-db-dev.env)
$ export $(xargs < env/app-srv-dev.env)
```

2. Build and exec the application
```bash
# in 'app' folder
$ go build
$ ./app
```
#### Execute Database Application
Only used in container.
```bash
# in 'golang-beego-postgres' folder
$ make up-database
```

## Web API
### Summary
* **POST /v1/author**
* **POST /v1/book**
* **GET /v1/books**
* **GET /v1/books?genre="genre"**
* **GET /v1/book/{book_id}**
* **GET /v1/authors**
* **GET /v1/author/{author_id}**

### API
#### **POST /v1/author**
```bash
curl -X POST -H "Content-Type: application/json" -d '{"name":"Isaac Asimov"}' http://localhost:8000/v1/author
```
###### Request
```json
{
  "name": "Isaac Asimov"
}
```

###### Response
```json
{
  "id": 1,
  "name": "Issac Asimov"
}
```

#### **POST /v1/book**
```bash
curl -X POST -H "Content-Type: application/json" -d '{"name":"The End of Eternity","genre":"Science Fiction","author":["Isaac Asimov"]}' http://localhost:8000/v1/book
```
###### Request
```json
{
  "name": "The End of Eternity",
  "genre": "Science Fiction",
  "author": [
    "Isaac Asimov"
  ]
}
```

###### Response
```json
{
  "id": 1
}
```

#### **GET /v1/books**
```bash
curl -X GET http://localhost:8000/v1/books
```
###### Response
```json
[
  {
    "id": 1,
    "name": "The End of Eternity",
    "genre": "Science Fiction",
    "author": [
      {
        "id": 1,
        "name": "Isaac Asimov"
      }
    ]
  },{
    "id": 2,
    "name": "The Talisman",
    "genre": "Dark Fantasy",
    "author": [
      {
        "id": 2,
        "name": "Stephen King"
      },{
        "id": 3,
        "name": "Peter Straub"
      }
    ]
  }
]
```

#### **GET /v1/books?genre="genre"**
```bash
curl -X GET http://localhost:8000/v1/books?genre="Science Fiction"
```
###### Response
```json
[
  {
    "id": 1,
    "name": "The End of Eternity",
    "genre": "Science Fiction",
    "author": [
      {
        "id": 1,
        "name": "Isaac Asimov"
      }
    ]
  },{
    "id": 3,
    "name": "Robot Visions",
    "genre": "Science Fiction",
    "author": [
      {
        "id": 1,
        "name": "Isaac Asimov"
      }
    ]
  }
]
```

#### **GET /v1/book/{book_id}**
```bash
curl -X GET http://localhost:8000/v1/book/1
```
###### Response
```json
{
  "id": 1,
  "name": "The End of Eternity",
  "genre": "Science Fiction",
  "author": [
    {
      "id": 1,
      "name": "Isaac Asimov"
    }
  ]
}
```

#### **GET /v1/authors**
```bash
curl -X GET http://localhost:8000/v1/authors
```
###### Response
```json
[
  {
    "id": 1,
    "name": "Isaac Asimov",
    "books": [
      {
        "id": 1,
        "name": "The End of Eternity"
      },{
        "id": 3,
        "name": "Robot Visions",
      }
    ]
  },{
    "id": 2,
    "name": "Stephen King",
    "books": [
      {
        "id": 2,
        "name": "The Talisman"
      }
    ]
  }
]
```

#### **GET /v1/author/{author_id}**
```bash
curl -X GET http://localhost:8000/v1/author/1
```
###### Response
```json
{
  "id": 1,
  "name": "Isaac Asimov",
  "books": [
    {
      "id": 1,
      "name": "The End of Eternity"
    },{
      "id": 3,
      "name": "Robot Visions",
    }
  ]
}
```
