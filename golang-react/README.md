# Golang - React

## Description
It's a simple application that uses **Golang** (Programming Language) on the Server side, and **React** (Javascript Library) on the Client side.
The application stores and retrieves items (entities - identity and name as properties) using a client-server architecture.

The communication is using **Promise** (asynchronous operation) and **CORS** (Cross-Origin Resource Sharing).
The configuration of the Client side is from scratch using **Webpack** (Javascript Module Bundler).
The communication is using Web Services (REST). The **Axios** library is used on the Client side, and the **Gorilla/Mux*** library is used for the Server Side.
As a graphical user interface for the Client side, the **Material-UI** is used.

---

## Execution
### Client
* Install package dependency

```bash
# in 'showcase-tech/golang-react/client' folder
$ npm i
```

* Execute the application as 'develop'

```bash
# in 'showcase-tech/golang-react/client' folder
$ npm run develop
```

### Server
* Execute Server
```bash
# in 'showcase-tech/golang-react/server' folder
$ go run main.go
```

---

## Packages

### Client
* Webpack
* Babel
* React
* Axios
* Material-UI

### Server
* Golang
* Gorilla/Mux

---

## Configuration
### Create 'package.json' File

* Create base 'package.json' file.

```bash
$ npm init
```

* Install webpack libraries.
```bash
$ npm install --save-dev webpack webpack-dev-server webpack-cli html-webpack-plugin webpack-merge
```

* Install babel libraries.
```bash
$ npm install --save-dev babel-loader @babel/core @babel/preset-env @babel/preset-react
```

* Install react
```bash
$ npm install react react-dom
```

* Install axios (communication library)
```bash
$ npm install axios
```

* Install material-ui (user interface library)
```bash
$ npm install @material-ui/core
```

---

## Application

### Initial View

#### Client
![alt text][img-client-0]

#### Server
![alt text][img-server-0]

### Update Data From the Server (Press 'Update Products' Button)

#### Client
![alt text][img-client-1]

#### Server
![alt text][img-server-1]

### Insert Data To the Server (Insert data and press 'Send' Button)

#### Client
![alt text][img-client-2]

#### Server
![alt text][img-server-2]

### Update Data From the Server

#### Client
![alt text][img-client-3]

#### Server
![alt text][img-server-3]

[img-client-0]: docs/0-client-init.png
[img-client-1]: docs/1-client-update.png
[img-client-2]: docs/2-client-insert.png
[img-client-3]: docs/3-client-update.png
[img-server-0]: docs/0-server-init.png
[img-server-1]: docs/1-server-update.png
[img-server-2]: docs/2-server-insert.png
[img-server-3]: docs/3-server-update.png
