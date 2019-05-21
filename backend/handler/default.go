package handler

import (
    "net/http"
)

func HandleDefault(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("OK"))
}