package main

import (
        "fmt"
        "os"
        goopt "github.com/droundy/goopt"
        cliutils "github.com/jshort/gocurl/cliutils"
        httputils "github.com/jshort/gocurl/httputils"
)

var amVerbose *bool

func main() {
        cliInputs := cliSetup()
        os.Exit(run(cliInputs))
}

func cliSetup() *cliutils.GoCurlCli {
        var cliMap = map[string]interface{}{
                "httpVerb"    : "GET",
                "httpHeaders" : []string{},
                "postData"    : "",
                "timeOut"     : 60,
                "arguments"   : []string{},
        }


        var httpVerb   = goopt.StringWithLabel([]string{"-X", "--command"}, cliMap["httpVerb"].(string),
                "COMMAND", fmt.Sprintf("HTTP verb for request: %s", cliutils.HttpVerbs))
        var httpHeaders = goopt.Strings([]string{"-H", "--header"},
                "KEY:VALUE", "Custom HTTP Headers to be sent with request (can pass multiple times)")
        var postData   = goopt.StringWithLabel([]string{"-d", "--data"}, cliMap["postData"].(string),
                "DATA", "HTTP Data for POST")
        var timeOut    = goopt.IntWithLabel([]string{"-t", "--timeout"}, cliMap["timeOut"].(int),
                "TIMEOUT", "Timeout in seconds for request")

        amVerbose      = goopt.Flag([]string{"-v", "--verbose"}, []string{}, "Verbose output", "")

        goopt.Summary = "Golang based http client program"
        goopt.Parse(nil)

        cliMap["httpVerb"]    = *httpVerb
        cliMap["httpHeaders"] = *httpHeaders
        cliMap["postData"]    = *postData
        cliMap["timeOut"]     = *timeOut
        cliMap["arguments"]   = goopt.Args

        cliInputs := cliutils.InititializeCli(cliMap)
        
        exitWithMessageIfNonZero(cliutils.ValidateOptions(cliInputs))
        exitWithMessageIfNonZero(cliutils.ValidateArguments(cliInputs))

        return cliInputs
}

func run(cliInputs *cliutils.GoCurlCli) int {
        printOptionsAndArgs(cliInputs)

        return httputils.SubmitRequest(cliInputs)
}

func exitWithMessageIfNonZero(code int, message string) {
        if code != 0 {
                fmt.Println(message)
                os.Exit(code)
        }
}

func printOptionsAndArgs(cliInputs *cliutils.GoCurlCli) {
        if *amVerbose {
                cliInputs.Print()
        }
}