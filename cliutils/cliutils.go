package cliutils

import (
        "fmt"
        "strings"
)

var HttpVerbs = [5]string{"GET", "PUT", "POST", "DELETE", "PATCH"}

type GoCurlCli struct {
        httpVerb string
        httpHeaders []string
        postData string
        timeout int
        verbose bool
        arguments []string
}

func InititializeCli(m map[string]interface{}) *GoCurlCli {
        cliInputs := new(GoCurlCli)
        cliInputs.httpVerb = m["httpVerb"].(string)
        cliInputs.httpHeaders = m["httpHeaders"].([]string)
        cliInputs.postData = m["postData"].(string)
        cliInputs.timeout = m["timeOut"].(int)
        cliInputs.verbose = m["verbose"].(bool)
        cliInputs.arguments = m["arguments"].([]string)
        return cliInputs
}



func ValidateOptions(cliInputs *GoCurlCli) (int, string) {
        if ! contains(HttpVerbs[:], cliInputs.httpVerb) {
                return 1, fmt.Sprintf("Unexpected HTTP Verb: %s", cliInputs.httpVerb)
        }
        for _, header := range cliInputs.httpHeaders {
                if strings.Count(header, ":") != 1 {
                        return 1, fmt.Sprintf("Unexpected HTTP Header: %s", header)
                }
        }
        return 0, ""
}

func ValidateArguments(cliInputs *GoCurlCli) (int, string) {
        if len(cliInputs.arguments) != 1 {
                return 2, "Expecting 1 positional argument."
        }
        return 0, ""
}

func (cliInputs *GoCurlCli) HttpVerb() string {
        return cliInputs.httpVerb
}

func (cliInputs *GoCurlCli) HttpHeaders() []string {
        return cliInputs.httpHeaders
}

func (cliInputs *GoCurlCli) PostData() string {
        return cliInputs.postData
}

func (cliInputs *GoCurlCli) TimeOut() int {
        return cliInputs.timeout
}

func (cliInputs *GoCurlCli) Verbose() bool {
        return cliInputs.verbose
}

func (cliInputs *GoCurlCli) Url() string {
        return cliInputs.arguments[0]
}

func (cliInputs *GoCurlCli) Print() {
        fmt.Printf("HTTP Verb is: %v\n", cliInputs.httpVerb)
        fmt.Printf("HTTP Headers are: %v\n", cliInputs.httpHeaders)
        fmt.Printf("POST Data is: %v\n", cliInputs.postData)
        fmt.Printf("Request Timeout is: %v\n", cliInputs.timeout)
        fmt.Printf("Url is: %v\n\n", cliInputs.arguments[0])
}

// Checks if an array/slice contains a given string
func contains(s []string, str string) bool {
        for _, a := range s {
                if a == str {
                        return true
                }
        }
        return false
}