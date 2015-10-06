package main

import (
        "fmt"
        "net/http"
        "encoding/json"
)

func redirect302(w http.ResponseWriter, r *http.Request) {
        http.Redirect(w, r, "http://www.golang.org", 302)
}

func redirect301(w http.ResponseWriter, r *http.Request) {
        http.Redirect(w, r, "http://www.golang.org", 301)
}

func jsonHandler(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json; charset=utf-8")

        myItems := []string{"item1", "item2", "item3"}

        jsonMap := map[string]interface{}{"key1": 5, "key2": 7, "list": myItems}
        a, _ := json.MarshalIndent(jsonMap, "", "    ")

        w.Write(a)
        return
}

func startHttpServer() {
        err := http.ListenAndServe(":8000", nil)
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
        http.HandleFunc("/json",    jsonHandler)
        go startHttpServer()
        go startHttpsServer()
        fmt.Println("HTTP/HTTPS servers started...")
        for {}
}
