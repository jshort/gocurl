package cliutils

import (
        "fmt"
        "strings"
)

var HttpVerbs = [6]string{"GET", "HEAD", "PUT", "POST", "DELETE", "PATCH"}

type GoCurlCli struct {
        httpVerb    string
        httpHeaders []string
        postData    string
        timeout     int
        verbose     bool
        redirect    bool
        color       bool
        sslSecure   bool
        arguments   []string
}

func NewGoCurlCli() *GoCurlCli {
        cliInputs := &GoCurlCli{
                httpVerb:    "GET",
                httpHeaders: []string{},
                postData:    "",
                timeout:     60,
                verbose:     false,
                redirect:    false,
                color:       false,
                sslSecure:   true,
                arguments:   []string{},
        }
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

func (cliInputs *GoCurlCli) SetHttpVerb(httpVerb string) {
        cliInputs.httpVerb = httpVerb
}

func (cliInputs *GoCurlCli) HttpHeaders() []string {
        return cliInputs.httpHeaders
}

func (cliInputs *GoCurlCli) SetHttpHeaders(headers []string) {
        cliInputs.httpHeaders = headers
}

func (cliInputs *GoCurlCli) PostData() string {
        return cliInputs.postData
}

func (cliInputs *GoCurlCli) SetPostData(postData string) {
        cliInputs.postData = postData
}

func (cliInputs *GoCurlCli) Timeout() int {
        return cliInputs.timeout
}

func (cliInputs *GoCurlCli) SetTimeout(timeout int) {
        cliInputs.timeout = timeout
}

func (cliInputs *GoCurlCli) Verbose() bool {
        return cliInputs.verbose
}

func (cliInputs *GoCurlCli) SetVerbose(verbose bool) {
        cliInputs.verbose = verbose
}

func (cliInputs *GoCurlCli) Redirect() bool {
        return cliInputs.redirect
}

func (cliInputs *GoCurlCli) SetRedirect(redirect bool) {
        cliInputs.redirect = redirect
}

func (cliInputs *GoCurlCli) Color() bool {
        return cliInputs.color
}

func (cliInputs *GoCurlCli) SetColor(color bool) {
        cliInputs.color = color
}

func (cliInputs *GoCurlCli) SslSecure() bool {
        return cliInputs.sslSecure
}

func (cliInputs *GoCurlCli) SetSslSecure(sslSecure bool) {
        cliInputs.sslSecure = sslSecure
}

func (cliInputs *GoCurlCli) SetArgs(arguments []string) {
        cliInputs.arguments = arguments
}

func (cliInputs *GoCurlCli) Url() string {
        return cliInputs.arguments[0]
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