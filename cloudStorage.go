package main

import (
    "fmt"
    "io"
    "net/http"
    "os"
    "path/filepath"
    "html/template"
)


func InitStorage(folder string) error {
    return os.MkdirAll(folder, os.ModePerm)
}

func UploadHandler(storageFolder string) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            w.WriteHeader(http.StatusMethodNotAllowed)
            fmt.Fprintf(w, "Only POST requests are allowed")
            return
        }

        r.ParseMultipartForm(10<<20) //10 MB max per file 

        file, handler, err := r.FormFile("file")
        if err != nil{
            http.Error(w, "Error reading file: "+err.Error(), http.StatusBadRequest)
            return
        }
        defer file.Close()

        dstPath := filepath.Join(storageFolder, handler.Filename)
        dst, err := os.Create(dstPath)
        if err != nil {
            http.Error(w, "Unable to create file: "+err.Error(), http.StatusInternalServerError)
            return
        }
        defer dst.Close()

        _, err = io.Copy(dst, file)
        if err != nil {
            http.Error(w, "Error saving file: "+err.Error(), http.StatusInternalServerError)
            return
        }

        fmt.Fprintf(w, "File uploaded successfully: %s\n", handler.Filename)
    }
}

func FilesHandler(storageFolder string) http.Handler {
    return http.StripPrefix("/files/", http.FileServer(http.Dir(storageFolder)))
}

func ListFilesHandler(storageFolder string) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        files, err := os.ReadDir(storageFolder)
        if err!= nil {
            http.Error(w, "Unable to read cloud files", http.StatusInternalServerError)
            return
        }

        var filenames []string
        for _, f:= range files {
            filenames = append(filenames, f.Name())
        }

        tmplt, err := template.ParseFiles("static/fileBrowserPage.html")
        if err != nil {
            http.Error(w, "Template Error", http.StatusInternalServerError)
            return
        }

        w.Header().Set("Content-type", "text/html")
        tmplt.Execute(w, filenames)
    }
}
