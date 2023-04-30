package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/", handleFormSubmission)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleFormSubmission(w http.ResponseWriter, r *http.Request) {
    // Handle the form submission here
    name := r.FormValue("name")
    email := r.FormValue("email")
    phone := r.FormValue("phone")
    message := r.FormValue("message")

    // Print the form data to the console
    fmt.Println("Name:", name)
    fmt.Println("Email:", email)
    fmt.Println("Phone:", phone)
    fmt.Println("Message:", message)

    // Return a response to the client
    fmt.Fprintln(w, "Form submitted successfully!")
}
