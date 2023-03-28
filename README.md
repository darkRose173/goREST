## Go-book-store-api
Go Sample project to implement MySQL CRUD operation for a book store management 

A production ready sample Book store RESTful API with Go using gorilla/mux with Mysql (relational Database). This Project contains a golang implementation of basic API endpoints that can be tested using POSTMAN

## Installation & Run
### Run the application via Docker
#### Requirements
**Step 1:** Create the Docker image according to Dockerfile. Ensure docker-compose is installed on your build system. For details on how to do this, see:                [Docker compose](https://docs.docker.com/compose/install/)

This may take a few minutes. In the root directory /darkRose173/book-store-api run the command:
``` 
$ docker compose build --no-cache
```
This will create two containers in background for Go and Mysql respectively
The verification of installation of the same can be verified in the docker Desktop application

**Step 2:** Run the application, starting both the MySQL database and the API containers
```
$ docker compose up 
```

**Step 3:** Open a webpage and access the following API endpoint.
```
http://localhost:8080/api/books/
```

**Response:**
API endpoint above returns the collection of all the books stored in the database 
(200 OK)

### POSTMAN API:
GET, PUT, POST, DELETE API's can be tested using POSTMAN API

**API** <br>
/api/books/ *GET*         : Retrieves all the books stored in the database <br>
/api/books/ *POST*        : Creates a new entry of a book <br>
/api/books/{id} *DELETE*  : Deletes an entry of a book with ID <br>
/api/books/{id} *PUT*     : To update a book entry with ID <br>

### Routing
gorilla/mux is being used to setup for routing. It provides some powerful feature like grouping/middleware/handler etc. (Gorilla Mux)[github.com/gorilla/mux]

	r := mux.NewRouter()

	r.HandleFunc("/api/books/", func(w http.ResponseWriter, r *http.Request) 
	r.HandleFunc("/api/books/{id}", func(w http.ResponseWriter, r *http.Request)

### DB
Mysql is being used as database database/sql and (MySQL-driver)[github.com/go-sql-driver/mysql] module to create and manage database connection

### wait-for
This extension waits for the database port to be set-up and then runs the API 
(wait-for)[https://raw.githubusercontent.com/eficode/wait-for/v2.1.0/wait-for]

### API Endpoint (GET/POST) 
```
http://127.0.0.1:8080/api/books/
```
### API Endpoint (DELETE/PUT)                
```
http://127.0.0.1:8080/api/books/{id}
```

### Structure
*main.go* -> Main Go application with API endpoints <br>
*book-store.sql* -> MySQL database for the book-store (ID, Title, ISBN & Author) <br>

*Docker files* <br>
api.Dockerfile -> Commands to instal the required dependencies and run the GO application <br>
db.Dockerfile  -> .sql file migration to the database <br>
docker-compose.yml -> Database and API container specification, MySQL database authentication and command to run the mysql database <br>
entrypoint.sh -> bash file running wait-for extension that waits for the database port to set up before running the API <br>
----------------------------------------------------------------------------------------------------------------------------------------------------------
