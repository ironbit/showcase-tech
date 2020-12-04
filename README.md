# Showcase Tech

## Technologies

* [Golang - Beego - Postgres](./golang-beego-postgres)
* [Golang - React](./golang-react)

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

## Golang - React
It's a simple application that uses **Golang** (Programming Language) on the Server side, and **React** (Javascript Library) on the Client side.
The application stores and retrieves items (entities - identity and name as properties) using a client-server architecture. Using Webpack, it's possible to generate the client files in the *build* folder or execute them in the webpack server in both modes (develop and release).

The communication is using **Promise** (asynchronous operation) and **CORS** (Cross-Origin Resource Sharing).
The configuration of the Client side is from scratch using **Webpack** (Javascript Module Bundler).
The communication is using Web Services (REST). The **Axios** library is used on the Client side, and the **Gorilla/Mux** library is used for the Server Side.
As a graphical user interface for the Client side, the **Material-UI** is used.

The technologies used are:
* Webpack
* Babel
* React
* Axios
* Material-UI
* Golang
* Gorilla/Mux
