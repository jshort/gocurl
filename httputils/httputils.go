package httputils

import (
        "net/http"
        "net/http/httputil"
        "bytes"
        "strings"
        "io/ioutil"
        "bufio"
        "fmt"
        "github.com/jshort/gocurl/cliutils"
)

const userAgent string = "Gocurl-client/1.0" 

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
        req, _ := http.NewRequest("GET", cliInputs.Url(), nil)

        prepareRequest(req, cliInputs.HttpHeaders(), cliInputs.Verbose())

        return processRequest(req, cliInputs.Verbose())
}

func post(cliInputs *cliutils.GoCurlCli) int {
        var postBody = []byte(cliInputs.PostData())
        req, _ := http.NewRequest("POST", cliInputs.Url(), bytes.NewBuffer(postBody))

        prepareRequest(req, cliInputs.HttpHeaders(), cliInputs.Verbose())

        return processRequest(req, cliInputs.Verbose())
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

func prepareRequest(req *http.Request, headers []string, verbose bool) {
        headerMap := parseHeaderString(headers)
        req.Header.Set("User-Agent", userAgent)
        for key, value := range headerMap {
                req.Header.Set(key, value)
        }

        if verbose {
                printRequest(req)
        }
}

func processRequest(req *http.Request, verbose bool) int {
        resp, err := client.Do(req)
        if err != nil {
                fmt.Printf("Error in processing request.\n")
                return 1
        }
        defer resp.Body.Close()
        respBody, err := ioutil.ReadAll(resp.Body)
        if err != nil {
                return 1
        }

        if verbose {
                printResponse(resp)
        }

        fmt.Printf("%s\n", respBody)
        return 0
}

func printRequest(req *http.Request) {
        reqDump, err := httputil.DumpRequestOut(req, true)
        fmt.Println("Request:")

        r := bufio.NewReader(bytes.NewBuffer(reqDump))
        line, isPrefix, err := r.ReadLine()
        for err == nil && !isPrefix {
                s := string(line)
                fmt.Printf("> %s\n", s)
                line, isPrefix, err = r.ReadLine()
        }
        fmt.Println("")
}

func printResponse(resp *http.Response) {
        respDump, err := httputil.DumpResponse(resp, false)
        fmt.Println("Response:")

        r := bufio.NewReader(bytes.NewBuffer(respDump))
        line, isPrefix, err := r.ReadLine()
        for err == nil && !isPrefix {
                s := string(line)
                fmt.Printf("< %s\n", s)
                line, isPrefix, err = r.ReadLine()
        }
        fmt.Println("")
}

// Assumes the header slice was validated (1 colon per entry)
func parseHeaderString(headers []string) map[string]string {
        var headerMap = make(map[string]string)
        for _, header := range headers {
                tokens := strings.Split(header, ":")
                headerMap[tokens[0]] = tokens[1]
        }
        return headerMap
}
