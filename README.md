## Go-book-store-api
Go Sample project to implement MySQL CRUD operation for a book store management 

A production ready sample Book store RESTful API with Go using gorilla/mux with Mysql (relational Database). This Project contains a golang implementation of basic API endpoints that can be tested using POSTMAN

Installation & Run
## Download this project
git clone git@github.com:darkRose173/goREST.git

The following database environment variables are set in the docker-compose.yml file
``` 
{
     MYSQL_RANDOM_ROOT_PASSWORD: "secret"
     MYSQL_DATABASE: "book-store"
     MYSQL_USER: "tester"
     MYSQL_PASSWORD: "secret"
} 
```

## Run the application via Docker
Requirements
Step 1: Create the Docker image according to Dockerfile. Ensure docker-compose is installed on your build system. For details on how to do this, see:                [Docker compose](https://docs.docker.com/compose/install/)

This may take a few minutes. In the root directory /darkRose173/book-store-api run the command:
``` 
$ docker compose build --no-cache
```
This will create two containers in background for Go and Mysql respectively
The verification of installation of the same can be verified in the docker Desktop application

Step 2:. Run the application, starting both the MySQL database and the API containers
# $ docker compose up 

Step 3: Open a webpage and access the following API endpoint.
# http://localhost:8080/api/books/

Response: 
API endpoint above returns the collection of books stored in the database 
(200 OK)

# POSTMAN API:
GET, PUT, POST, DELETE API's can be tested using POSTMAN API

API
/api/books/ GET         : Retrieves all the books stored in the database 
/api/books/ POST        : Creates a new entry of a book
/api/books/{id} DELETE  : Deletes an entry of a book with ID
/api/books/{id} PUT     : To update a book entry with ID

Routing
gorilla/mux is being used to setup for routing. It provides some powerful feature like grouping/middleware/handler etc.

	r := mux.NewRouter()

	r.HandleFunc("/api/books/", func(w http.ResponseWriter, r *http.Request) 
	r.HandleFunc("/api/books/{id}", func(w http.ResponseWriter, r *http.Request)

DB
Mysql is being used as database database/sql and github.com/go-sql-driver/mysql module to create and manage database connection

wait-for
This extension waits for the database port to be set-up
https://raw.githubusercontent.com/eficode/wait-for/v2.1.0/wait-for /usr/local/bin/wait-for

# API Endpoint : http://127.0.0.1:8080/api/books/
#                http://127.0.0.1:8080/api/books/{id}
                 
Structure
main.go -> Main Go application with API endpoints
book-store.sql -> MySQL database for the book-store (ID, Title, ISBN & Author)

//Docker files
api.Dockerfile -> Commands to instal the required dependencies and run the GO application
db.Dockerfile  -> .sql file migration to the database
docker-compose.yml -> Database and API container specification, MySQL database authentication and command to run the mysql database
entrypoint.sh -> bash file running wait-for extension that waits for the database port tobe set up before running the API
