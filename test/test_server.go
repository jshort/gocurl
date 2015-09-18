package main

import (
        "fmt"
        "net/http"
)

func redirect302(w http.ResponseWriter, r *http.Request) {
        http.Redirect(w, r, "http://www.golang.org", 302)
}

func redirect301(w http.ResponseWriter, r *http.Request) {
        http.Redirect(w, r, "http://www.golang.org", 301)
}

func startHttpServer() {
        err := http.ListenAndServe(":8080", nil)
        if err != nil {
                fmt.Printf("ListenAndServe: %v", err)
        }
}

func startHttpsServer() {
        err := http.ListenAndServeTLS(":8443", "cert.pem", "key.pem", nil)
        if err != nil {
                fmt.Printf("ListenAndServeTLS: %v", err)
        }
}

func main() {
        http.HandleFunc("/test302", redirect302)
        http.HandleFunc("/test301", redirect301)
        go startHttpServer()
        go startHttpsServer()
        fmt.Println("HTTP/HTTPS servers started...")
        for {}
}
