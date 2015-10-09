package main

import (
        "fmt"
        "strings"
)

var HttpVerbs = [6]string{"GET", "HEAD", "PUT", "POST", "DELETE", "PATCH"}

type cliOptions struct {
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

func newCliOptions() *cliOptions {
        options := &cliOptions{
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
        return options
}

func validateOptions(options *cliOptions) (int, string) {
        if ! contains(HttpVerbs[:], options.httpVerb) {
                return 1, fmt.Sprintf("Unexpected HTTP Verb: %s", options.httpVerb)
        }
        for _, header := range options.httpHeaders {
                if strings.Count(header, ":") != 1 {
                        return 1, fmt.Sprintf("Unexpected HTTP Header: %s", header)
                }
        }
        return 0, ""
}

func validateArguments(options *cliOptions) (int, string) {
        if len(options.arguments) != 1 {
                return 2, "Expecting 1 positional argument."
        }
        return 0, ""
}

func (options *cliOptions) url() string {
        return parseUrl(options.arguments[0])
}

func parseUrl(rawurl string) string {
        parsedurl := rawurl
        if ! (strings.HasPrefix(rawurl, "http://") || strings.HasPrefix(rawurl, "https://")) {
                parsedurl = "http://" + rawurl
        }
        return parsedurl
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