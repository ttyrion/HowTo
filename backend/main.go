package main

import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
    "HowTo/backend/handler"
    "HowTo/backend/env"
)

func main() {
    router := mux.NewRouter()
    router.HandleFunc("/", handler.HandleDefault).Methods("GET")
    host, port := env.GetCmdlineHostAndPort()
    server := fmt.Sprintf("%s:%d", host, port)
    fmt.Println("starting service on", server)
    err := http.ListenAndServe(server, router)
    if err != nil {
        fmt.Println("start service faild.", err)
        return
    }
}
