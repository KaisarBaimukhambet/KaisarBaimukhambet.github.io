package main

import (
    "fmt"
    "net/http"
)//Импортируем пакет net/http для создания сервера

func rootHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!")
}//http.ResponseWriter используется для отправки ответа клиенту

func main() {
    mux := http.NewServeMux()//содает новый роутер
    mux.HandleFunc("/", rootHandler)

    fmt.Println("Server listening on port 8000")
    http.ListenAndServe(":8000", mux)// запускаем сервер на порту 8000 и используем мультиплексный маршрутизатор для обработки входящих запросов
}
