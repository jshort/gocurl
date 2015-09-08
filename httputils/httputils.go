package httputils

import (
        "net/http"
        "io/ioutil"
        "fmt"
)

func Get(url string, options map[string]string) int {

        resp, err := http.Get(url)
        if err != nil {
                fmt.Printf("error in Get\n")
                return 1
        }
        defer resp.Body.Close()
        respBody, err := ioutil.ReadAll(resp.Body)
        if err != nil {
                return 1
        }
        fmt.Printf("Response Body:\n\n%s\n", respBody)
        return 0
}

func Post(url string, options map[string]string) int {
        return dummyReturn()
}

func Put(url string, options map[string]string) int {
        return dummyReturn()
}

func Delete(url string, options map[string]string) int {
        return dummyReturn()
}

func Patch(url string, options map[string]string) int {
        return dummyReturn()
}

func dummyReturn() int {
        fmt.Println("Not implemented, exiting...")        
        return 255
}