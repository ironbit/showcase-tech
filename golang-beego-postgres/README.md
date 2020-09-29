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
