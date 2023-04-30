package main

import (
    "database/sql"
    "fmt"
    "log"
    "net/http"

    _ "github.com/mattn/go-sqlite3"
)

func main() {
    // Create a new database file and connect to it
    db, err := sql.Open("sqlite3", "./example.db")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Create a table for form submissions
    _, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS form_submissions (
            id INTEGER PRIMARY KEY,
            name TEXT,
            email TEXT,
            phone TEXT,
            message TEXT
        )
    `)
    if err != nil {
        log.Fatal(err)
    }

    // Register the form submission handler
    http.HandleFunc("/", handleFormSubmission)
tmpl := template.Must(template.ParseFiles("contact.html"))
    // Start the server
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleFormSubmission(w http.ResponseWriter, r *http.Request) {
    // Handle the form submission here
    name := r.FormValue("name")
    email := r.FormValue("email")
    phone := r.FormValue("phone")
    message := r.FormValue("message")

    // Insert the form data into the database
    db, err := sql.Open("sqlite3", "./example.db")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    _, err = db.Exec(`
        INSERT INTO form_submissions (name, email, phone, message)
        VALUES (?, ?, ?, ?)
    `, name, email, phone, message)
    if err != nil {
        log.Fatal(err)
    }

    // Print the form data to the console
    fmt.Println("Name:", name)
    fmt.Println("Email:", email)
    fmt.Println("Phone:", phone)
    fmt.Println("Message:", message)

    // Return a response to the client
    fmt.Fprintln(w, "Form submitted successfully!")
}
