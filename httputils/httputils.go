package httputils

import (
        "net/http"
        "bytes"
        "strings"
        "io/ioutil"
        "fmt"
        "github.com/jshort/gocurl/cliutils"
)

var client *http.Client = http.DefaultClient

func SubmitRequest(cliInputs *cliutils.GoCurlCli) int {
        var retval int

        switch cliInputs.HttpVerb() {
        case "GET":
                retval = get(cliInputs)
        case "POST":
                retval = post(cliInputs)
        case "PUT":
                retval = put(cliInputs)
        case "DELETE":
                retval = delete(cliInputs)
        case "PATCH":
                retval = patch(cliInputs)
        }

        return retval
}

func get(cliInputs *cliutils.GoCurlCli) int {

        req, err := http.NewRequest("GET", cliInputs.Url(), nil)
        for key, value := range parseHeaderString(cliInputs.HttpHeaders()) {
                req.Header.Set(key, value)
        }

        fmt.Printf("Request Header:\n%v\n", req.Header)

        resp, err := client.Do(req)
        if err != nil {
                fmt.Printf("error in get\n")
                return 1
        }
        defer resp.Body.Close()
        respBody, err := ioutil.ReadAll(resp.Body)
        if err != nil {
                return 1
        }
        fmt.Printf("Response Header:\n%v\n", resp.Header)
        fmt.Printf("Response Status:\n%s\n", resp.Status)
        fmt.Printf("Response Body:\n%s\n", respBody)
        return 0
}

func post(cliInputs *cliutils.GoCurlCli) int {
        var postBody = []byte(cliInputs.PostData())
        req, err := http.NewRequest("POST", cliInputs.Url(), bytes.NewBuffer(postBody))
        for key, value := range parseHeaderString(cliInputs.HttpHeaders()) {
                req.Header.Set(key, value)
        }

        fmt.Printf("Request Header:\n%v\n", req.Header)

        resp, err := client.Do(req)
        if err != nil {
                fmt.Printf("error in post\n")
                return 1
        }
        defer resp.Body.Close()
        respBody, err := ioutil.ReadAll(resp.Body)
        if err != nil {
                return 1
        }
        fmt.Printf("Response Header:\n%v\n", resp.Header)
        fmt.Printf("Response Status:\n%s\n", resp.Status)
        fmt.Printf("Response Body:\n%s\n", respBody)
        return 0
}

func put(cliInputs *cliutils.GoCurlCli) int {
        return dummyReturn()
}

func delete(cliInputs *cliutils.GoCurlCli) int {
        return dummyReturn()
}

func patch(cliInputs *cliutils.GoCurlCli) int {
        return dummyReturn()
}

func dummyReturn() int {
        fmt.Println("Not implemented, exiting...")        
        return 255
}

func parseHeaderString(headers []string) map[string]string {
        var headerMap = make(map[string]string)
        for _, header := range headers {
                tokens := strings.Split(header, ":")
                // Should have validation
                headerMap[tokens[0]] = tokens[1]
        }
        return headerMap
}