package handler

import "net/http"

func HandleRecover(h http.Handler) http.Handler {
    return &recoverHandler{
        handler: h,
    }
}

type recoverHandler struct {
    handler http.Handler
}

func(rc *recoverHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    defer func() {
        if err := recover(); err != nil {
            http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
        }
    }()

    rc.handler.ServeHTTP(w,r)
}