package main

import (
        "fmt"
        "os"
        goopt "github.com/jshort/goopt"
        cliutils "github.com/jshort/gocurl/cliutils"
        httputils "github.com/jshort/gocurl/httputils"
)

func main() {
        args, options := cliSetup()
        os.Exit(run(args, options))
}

func cliSetup() ([]string, map[string]string) {
        var options = map[string]string{
                "httpVerb"   : "GET",
                "httpHeader" : "",
                "postData"   : "",
                "timeOut"    : "60",
        }


        var httpVerb   = goopt.StringWithLabel([]string{"-X", "--command"}, options["httpVerb"],
                "COMMAND", fmt.Sprintf("HTTP verb for request: %s", cliutils.HttpVerbs))
        var httpHeader = goopt.StringWithLabel([]string{"-H", "--header"}, options["httpHeader"], "HEADER", "Custom HTTP Header to be sent with request")
        var postData   = goopt.StringWithLabel([]string{"-d", "--data"}, options["postData"], "DATA", "HTTP Data for POST")
        var timeOut    = goopt.StringWithLabel([]string{"-t", "--timeout"}, options["timeOut"], "TIMEOUT", "Timeout in seconds for request")

        goopt.Summary = "Golang based http client program"
        goopt.Parse(nil)

        options["httpVerb"]   = *httpVerb
        options["httpHeader"] = *httpHeader
        options["postData"]   = *postData
        options["timeOut"]    = *timeOut
        
        cliutils.ValidateOptions(options)
        cliutils.ValidateArguments(goopt.Args)

        return goopt.Args, options
}

func run(args []string, options map[string]string) int {
        for key, value := range options {
                fmt.Printf("%s is: %s\n", key, value)
        }
        fmt.Println("Args are: ", args)

        var retval int

        switch options["httpVerb"] {
        case "GET":
                fmt.Println("Hi James")
                retval = httputils.Get(args[0], options)
        case "POST":
                retval = httputils.Post(args[0], options)
        case "PUT":
                retval = httputils.Put(args[0], options)
        case "DELETE":
                retval = httputils.Delete(args[0], options)
        case "PATCH":
                retval = httputils.Patch(args[0], options)
        }

        return retval
}

// Helpers to be moved