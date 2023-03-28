/*-----------------------------------------------------------
File name     : main.go
Author	      : NRaghuraj
Date Created  : 2023 MAR 24
Date Modified : 2023 MAR 27
Description   : GO REST API to perform simple CRUD
				operations from a MySQL database
				{book-store.sql}
				APIs tested  : {GET, PUT, POST, DELETE}
				Functionality: GETALLBOOKS, UPDATEBOOKBYID
							   CREATEBOOK, DELETEBOOK
-----------------------------------------------------------*/

package main

//Required packages
import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// Book represents a book in the book store.
type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	ISBN   string `json:"isbn"`
	Author string `json:"author"`
}

func main() {

	fmt.Println("Trying to establish connection to database..")

	// Connect to the database root password is null for testing purposes
	db, err := sql.Open("mysql", "tester:secret@tcp(db:3306)/book-store")

	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Connected to MySQL database!")
	defer db.Close()

	//Instantiate a new router
	r := mux.NewRouter()

	/*-----------------------------------------------------------------------------
		Switch Case to get/retrieve all books (GET) and
		insert or create a book into the database (POST)
	-----------------------------------------------------------------------------*/

	// Define the HTTP handler for the "get books" endpoint.
	r.HandleFunc("/api/books/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		/*-------------------------------------------------------------------------------------------
				API: Get all the books
				Method: GET
				Desc: Get all the books stored in the database
				Access Type: Private
		--------------------------------------------------------------------------------------------*/
		case "GET":
			// Retrieve the list of books from the database.
			rows, err := db.Query("SELECT id, title, isbn, author FROM book_table")
			if err != nil {
				log.Println(err)
				http.Error(w, "Failed to retrieve books", http.StatusInternalServerError)
				return
			}
			defer rows.Close()

			// Iterate over the rows and create a slice of Book structs.
			books := []Book{}
			for rows.Next() {
				book := Book{}
				err := rows.Scan(&book.ID, &book.Title, &book.ISBN, &book.Author)
				if err != nil {
					log.Println(err)
					http.Error(w, "Failed to retrieve books", http.StatusInternalServerError)
					return
				}
				books = append(books, book)
			}

			// Marshal the slice of Book structs to JSON and write it to the response.
			w.Header().Set("Content-Type", "application/json")
			err = json.NewEncoder(w).Encode(books)
			if err != nil {
				log.Println(err)
				http.Error(w, "Failed to encode books", http.StatusInternalServerError)
				return
			}

			/*--------------------------------------------------------------------------------------------
					API: Create a new book
					Method: POST
					Desc: Create/Add a book into the database
					Access Type: Private
			--------------------------------------------------------------------------------------------*/
		case "POST":
			// Decode the JSON request body into a new Book struct.
			book := Book{}
			err := json.NewDecoder(r.Body).Decode(&book)
			if err != nil {
				log.Println(err)
				http.Error(w, "Failed to decode request body", http.StatusBadRequest)
				return
			}

			// Insert the new book into the database.
			result, err := db.Exec("INSERT INTO book_table (ID, title, isbn, author) VALUES (?, ?, ?, ?)", book.ID, book.Title, book.ISBN, book.Author)
			if err != nil {
				log.Println(err)
				http.Error(w, "Failed to create book", http.StatusInternalServerError)
				return
			}

			// Retrieve the ID of the new book.
			id, err := result.LastInsertId()
			if err != nil {
				log.Println(err)
				http.Error(w, "Failed to retrieve book ID", http.StatusInternalServerError)
				return
			}

			// Set the ID of the new book and marshal it to JSON.
			book.ID = strconv.Itoa(int(id))
			w.Header().Set("Content-Type", "application/json")
			err = json.NewEncoder(w).Encode(book)
			if err != nil {
				log.Println(err)
				http.Error(w, "Failed to encode book", http.StatusInternalServerError)
				return
			}

			// Return success message
			w.WriteHeader(http.StatusCreated)
			fmt.Fprintln(w, "Book successfully added")
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
	})

	// /*--------------------------------------------------------------------------------------
	// 	   Switch Case to delete a book by ID (DELETE) and
	// 	   update a book by ID from the database (PUT)
	// --------------------------------------------------------------------------------------*/

	r.HandleFunc("/api/books/{id}", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		// 	/*--------------------------------------------------------------------------------------------
		// 			API: Delete a Book
		// 			Method: DELETE
		// 			Desc: Delete a book from the store by the book ID
		// 			Access Type: Private
		// 	--------------------------------------------------------------------------------------------*/
		case "DELETE":
			// Extract the book ID from the request URL.
			vars := mux.Vars(r)
			id, err := strconv.Atoi(vars["id"])
			if err != nil {
				log.Println(err)
				http.Error(w, "Invalid book ID", http.StatusBadRequest)
				return
			}

			// Delete the book from the database.
			result, err := db.Exec("DELETE FROM book_table WHERE id = ?", id)
			if err != nil {
				log.Println(err)
				http.Error(w, "Failed to delete book", http.StatusInternalServerError)
				return
			}

			// Check if the book was actually deleted.
			rowsAffected, err := result.RowsAffected()
			if err != nil {
				log.Println(err)
				http.Error(w, "Failed to delete book", http.StatusInternalServerError)
				return
			}
			if rowsAffected == 0 {
				http.Error(w, "Book not found", http.StatusNotFound)
				return
			}

			// Return a success message.
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "text/plain")
			_, err = w.Write([]byte("Book successfully deleted"))
			if err != nil {
				log.Println(err)
				return
			}
			// 		/*--------------------------------------------------------------------------------------------
			// 				API: Update an existing book
			// 				Method: PUT
			// 				Desc: Edit or alter an existing book in the store by book ID
			// 				Access Type: Private
			// 		--------------------------------------------------------------------------------------------*/
		case "PUT":
			// Get the book ID from the request parameters.
			vars := mux.Vars(r)
			id := vars["id"]

			// Decode the JSON request body into a new Book struct.
			book := Book{}
			err := json.NewDecoder(r.Body).Decode(&book)
			if err != nil {
				log.Println(err)
				http.Error(w, "Failed to decode request body", http.StatusBadRequest)
				return
			}

			// Update the book in the database.
			result, err := db.Exec("UPDATE book_table SET title = ?, isbn = ?, author = ? WHERE id = ?", book.Title, book.ISBN, book.Author, id)
			if err != nil {
				log.Println(err)
				http.Error(w, "Failed to update book", http.StatusInternalServerError)
				return
			}

			// Check if the book was found and updated.
			rowsAffected, err := result.RowsAffected()
			if err != nil {
				log.Println(err)
				http.Error(w, "Failed to retrieve rows affected", http.StatusInternalServerError)
				return
			}
			if rowsAffected == 0 {
				http.Error(w, "Book not found", http.StatusNotFound)
				return
			}

			// Write a success response with the updated book.
			w.Header().Set("Content-Type", "application/json")
			book.ID = id
			err = json.NewEncoder(w).Encode(book)
			if err != nil {
				log.Println(err)
				http.Error(w, "Failed to encode book", http.StatusInternalServerError)
				return
			}

			// Return success message
			w.WriteHeader(http.StatusCreated)
			fmt.Fprintln(w, "Book entry successfully updated")
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
	})
	// Start the HTTP server.
	log.Println("Starting server on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}

//--------------------------------------EOF---------------------------------------------//
