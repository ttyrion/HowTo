package handler

import (
    "net/http"
    "fmt"
    "io/ioutil"
)

func HandleUpload(w http.ResponseWriter, r *http.Request) {
    r.ParseMultipartForm(1024)
    fHeader := r.MultipartForm.File["upload"][0]
    file, err := fHeader.Open()
    if err != nil {
        fmt.Println("fileheader Open error.")
    }

    var data []byte
    data, err = ioutil.ReadAll(file)
    if err != nil {
        fmt.Println("ReadAll faild:", fHeader)
    }

    w.Write(data)
}