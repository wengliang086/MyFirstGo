package main

import (
	"GoLearnNotes/errhandling/filelistingserver/filelisting"
	"log"
	"net/http"
	"os"
)

type appHandler func(writer http.ResponseWriter, request *http.Request) error

func errorWrapper(handler appHandler) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Panic: %v", r)
				http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()

		err := handler(writer, request)
		if err == nil {
			return
		}

		log.Printf("Error occured! handling request: %s", err.Error())

		if u, ok := err.(userError); ok {
			http.Error(writer, u.Message(), http.StatusBadRequest)
			return
		}

		code := http.StatusOK
		switch {
		case os.IsPermission(err):
			code = http.StatusForbidden
		case os.IsNotExist(err):
			code = http.StatusNotFound
		default:
			code = http.StatusInternalServerError
		}
		http.Error(writer, http.StatusText(code), code)
	}
}

type userError interface {
	error
	Message() string
}

func main() {
	http.HandleFunc("/list/", errorWrapper(filelisting.HandleFileList))
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		panic(err)
	}
}
