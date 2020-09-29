# Golang - Beego - Potsgres

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
  The compose.yaml file for docker-compose and the app.dockerfile for the web rest api application.

* **env**

  Contains the environmental variables for the server and the database.
  app-db.env is for the database.
  app-srv.env is for the application.

* **sql**

  Contains the files for the initial configuration of the database.
  10-db.sh is for the creation of the user and the database.
  20-table.sh is for the creation of the tables in the database.

## Deploy
```bash
# in 'golang-beego-postgres' folder
$ docker-compose --env-file env/app-srv.env -p app -f docker/compose.yaml up
```

* --env-file

  export the environmental variables in 'app-srv.env'.

* -p

  project name or prefix

* -f

  docker compose file


The docker compose will create two containers (app_srv_1 & app_db_1), a network (app_net) and a volume (app_vol).

```bash
$ docker container ls -a
```

```bash
$ docker network ls
```

```bash
$ docker volume ls
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
curl -X POST -H "Content-Type: application/json" -d '{"name":"Isaac Asimov"}' http://localhost:12345/v1/author
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
  "id": 1
}
```

#### **POST /v1/book**
```bash
curl -X POST -H "Content-Type: application/json" -d '{"name":"The End of Eternity","genre":"Science Fiction","author":["Isaac Asimov"]}' http://localhost:12345/v1/book
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
curl -X GET http://localhost:12345/v1/books
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
curl -X GET http://localhost:12345/v1/books?genre="Science Fiction"
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
curl -X GET http://localhost:12345/v1/book/1
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
curl -X GET http://localhost:12345/v1/authors
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
curl -X GET http://localhost:12345/v1/author/1
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
