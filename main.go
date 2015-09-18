package main

import (
        "fmt"
        "os"
        goopt "github.com/droundy/goopt"
        cliutils "github.com/jshort/gocurl/cliutils"
        httputils "github.com/jshort/gocurl/httputils"
)

func main() {
        cliInputs := cliSetup()
        os.Exit(run(cliInputs))
}

func cliSetup() *cliutils.GoCurlCli {

        cliInputs := cliutils.NewGoCurlCli()

        var httpVerb   = goopt.StringWithLabel([]string{"-X", "--command"}, cliInputs.HttpVerb(),
                "COMMAND", fmt.Sprintf("HTTP verb for request: %s", cliutils.HttpVerbs))
        var httpHeaders = goopt.Strings([]string{"-H", "--header"},
                "KEY:VALUE", "Custom HTTP Headers to be sent with request (can pass multiple times)")
        var postData   = goopt.StringWithLabel([]string{"-d", "--data"}, cliInputs.PostData(),
                "DATA", "HTTP Data for POST")
        var timeOut    = goopt.IntWithLabel([]string{"-t", "--timeout"}, cliInputs.Timeout(),
                "TIMEOUT", "Timeout in seconds for request")
        var isVerbose  = goopt.Flag([]string{"-v", "--verbose"}, []string{}, "Verbose output", "")
        var hasColor   = goopt.Flag([]string{"-c", "--color"}, []string{}, "Colored output", "")


        goopt.Summary = "Golang based http client program"
        goopt.Parse(nil)

        cliInputs.SetHttpVerb(*httpVerb)
        cliInputs.SetHttpHeaders(*httpHeaders)
        cliInputs.SetPostData(*postData)
        cliInputs.SetTimeout(*timeOut)
        cliInputs.SetVerbose(*isVerbose)
        cliInputs.SetColor(*hasColor)
        cliInputs.SetArgs(goopt.Args)
        
        exitWithMessageIfNonZero(cliutils.ValidateOptions(cliInputs))
        exitWithMessageIfNonZero(cliutils.ValidateArguments(cliInputs))

        return cliInputs
}

func run(cliInputs *cliutils.GoCurlCli) int {
        return httputils.SubmitRequest(cliInputs)
}

func exitWithMessageIfNonZero(code int, message string) {
        if code != 0 {
                fmt.Println(message)
                os.Exit(code)
        }
}
