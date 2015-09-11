package cliutils

import (
        "fmt"
)

var HttpVerbs = []string{"GET", "PUT", "POST", "DELETE", "PATCH"}

func ValidateOptions(m map[string]string) (int, string) {
        if ! contains(HttpVerbs, m["httpVerb"]) {
                return 1, fmt.Sprintf("Unexpected HTTP Verb: %s", m["httpVerb"])
        }
        return 0, ""
}

func ValidateArguments(args []string) (int, string) {
        if len(args) != 1 {
                return 2, "Expecting 1 positional argument."
        }
        return 0, ""
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