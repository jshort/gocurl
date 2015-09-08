package cliutils

import (
        "fmt"
        "os"
)

var HttpVerbs = []string{"GET", "PUT", "POST", "DELETE", "PATCH"}

func ValidateOptions(m map[string]string) {
        if ! contains(HttpVerbs, m["httpVerb"]) {
                exitWithMessage(1, fmt.Sprintf("Unnexpected HTTP Verb: %s", m["httpVerb"]))
        }
}

func ValidateArguments(args []string) {
        if len(args) != 1 {
                exitWithMessage(1, "Expecting 1 positional argument.")
        }
}

func exitWithMessage(code int, message string) {
        fmt.Println(message)
        os.Exit(code)
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