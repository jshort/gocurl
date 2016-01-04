package main

import (
        "fmt"
        "net/http"
        "encoding/json"
        "compress/gzip"
        "io"
        "strings"
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

func gzipHandler(w http.ResponseWriter, r *http.Request) {
        if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
                handleGzip(w, r)
                return
        }
        w.Header().Set("Content-Encoding", "gzip")
        gz := gzip.NewWriter(w)
        defer gz.Close()
        gzr := gzipResponseWriter{Writer: gz, ResponseWriter: w}
        handleGzip(gzr, r)
}

type gzipResponseWriter struct {
        io.Writer
        http.ResponseWriter
}

func (w gzipResponseWriter) Write(b []byte) (int, error) {
        return w.Writer.Write(b)
}

func handleGzip(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "text/plain")
        w.Write([]byte("This is a test."))
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
        http.HandleFunc("/gzip",    gzipHandler)
        go startHttpServer()
        go startHttpsServer()
        fmt.Println("HTTP/HTTPS servers started...")
        ch := make(chan bool)
        <-ch
}
