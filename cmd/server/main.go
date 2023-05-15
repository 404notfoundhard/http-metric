package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func HttpError(statusCode int, w http.ResponseWriter) {
	switch statusCode {
	case http.StatusForbidden:
		b, err := os.ReadFile("./custom_responce/forbidden/forbidden.html")
		if err == nil {
			w.WriteHeader(http.StatusForbidden)
			io.WriteString(w, string(b))
		} else {
			panic("Custom Page Not Found!")
		}
	case http.StatusNotFound:
		b, err := os.ReadFile("./custom_responce/notfound/notfound.html")
		if err == nil {
			w.WriteHeader(http.StatusNotFound)
			io.WriteString(w, string(b))
		} else {
			panic("Custom Page Not Found!")
		}
	default:
		b, err := os.ReadFile("./custom_responce/notfound/notfound.html")
		if err == nil {
			w.WriteHeader(http.StatusNotFound)
			io.WriteString(w, string(b))
		} else {
			panic("Custom Page Not Found!")
		}
	}
}

func update(responce http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(responce, "Only POST requests are allowed!", http.StatusMethodNotAllowed)
	}
	fmt.Println(request.URL)

}

func main() {
	mux := http.NewServeMux()
	ForbiddenfileServer := http.FileServer(http.Dir("./custom_responce/"))
	mux.Handle(`/custom_responce/`, http.StripPrefix("/custom_responce/", ForbiddenfileServer))
	mux.HandleFunc(`/update/`, update)
	err := http.ListenAndServe(`:8080`, mux)
	if err != nil {
		panic(err)
	}
}
