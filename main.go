package main

import (
        "fmt"
        "os"
        goopt     "github.com/droundy/goopt"
)

func main() {
        options := cliSetup()
        os.Exit(run(options))
}

func cliSetup() *cliOptions {

        options := newCliOptions()

        var httpVerb       = goopt.StringWithLabel([]string{"-X", "--command"}, options.httpVerb,
                "COMMAND", fmt.Sprintf("HTTP verb for request: %s", HttpVerbs))
        var httpHeaders    = goopt.Strings([]string{"-H", "--header"},
                "KEY:VALUE", "Custom HTTP Headers to be sent with request (can pass multiple times)")
        var postData       = goopt.StringWithLabel([]string{"-d", "--data"}, options.postData,
                "DATA", "HTTP Data for POST")
        var timeOut        = goopt.IntWithLabel([]string{"-t", "--timeout"}, options.timeout,
                "TIMEOUT", "Timeout in seconds for request")
        var shouldRedirect = goopt.Flag([]string{"-r", "--redirect"}, []string{}, "Follow redirects", "")
        var isVerbose      = goopt.Flag([]string{"-v", "--verbose"}, []string{}, "Verbose output", "")
        var hasColor       = goopt.Flag([]string{"-c", "--color"}, []string{}, "Colored output", "")
        var isInsecureSSL  = goopt.Flag([]string{"-k", "--insecure"}, []string{}, "Allow insecure https connections", "")


        goopt.Summary = "Golang based http client program"
        goopt.Parse(nil)

        options.httpVerb = *httpVerb
        options.httpHeaders = *httpHeaders
        options.postData = *postData
        options.timeout = *timeOut
        options.verbose = *isVerbose
        options.redirect = *shouldRedirect
        options.color = *hasColor
        options.sslSecure = ! *isInsecureSSL
        options.arguments = goopt.Args
        
        exitWithMessageIfNonZero(validateOptions(options))
        exitWithMessageIfNonZero(validateArguments(options))

        return options
}

func run(options *cliOptions) int {
        return submitRequest(options)
}

func exitWithMessageIfNonZero(code int, message string) {
        if code != 0 {
                fmt.Println(message)
                os.Exit(code)
        }
}
