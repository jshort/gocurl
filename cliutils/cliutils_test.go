package cliutils

import (
        "testing"
        "strings"
)

func TestContains(t *testing.T) {
        cases := []struct {
                in string
                want bool
        }{
                {"GET", true},
                {"PUT", true},
                {"POST", true},
                {"DELETE", true},
                {"PATCH", true},
                {"POKE", false},
                {"", false},
                {"get", false},
        }

        for _, c := range cases {
                got := contains(HttpVerbs, c.in)
                if got != c.want {
                        t.Errorf("contains(%q, %q) == %v, want %v", HttpVerbs, c.in, got, c.want)
                }
        }
}

func TestValidateOptions(t *testing.T) {
        // Happy Path
        var options = map[string]string{
                "httpVerb"   : "GET",
                "httpHeader" : "",
                "postData"   : "",
                "timeOut"    : "60",
        }
        code, _ := ValidateOptions(options)

        if code != 0 {
                t.Errorf("ValidateOptions(%v) returned %v, want %v", options, code, 0)
        }


        // Error Path
        options = map[string]string{
                "httpVerb"   : "GOT",
                "httpHeader" : "",
                "postData"   : "",
                "timeOut"    : "60",
        }
        code, message := ValidateOptions(options)

        if code != 1 {
                t.Errorf("ValidateOptions(%v) returned %v, want %v", options, code, "non zero")
        }
        if ! strings.Contains(message, "Unexpected HTTP Verb") {
                t.Error("Error message should contain \"Unexpected HTTP Verb\"")
        }
}

func TestValidateArguments(t *testing.T) {
        // Happy Path
        var args = []string{"http://google.com"}
        code, _ := ValidateArguments(args)

        if code != 0 {
                t.Errorf("ValidateArguments(%v) returned %v, want %v", args, code, 0)
        }

        // Error Path
        args = []string{"http://google.com", "http://localhost:8080"}
        code, message := ValidateArguments(args)

        if code != 2 {
                t.Errorf("ValidateArguments(%v) returned %v, want %v", args, code, "non zero")
        }
        if ! strings.Contains(message, "Expecting 1 positional argument") {
                t.Error("Error message should contain \"Expecting 1 positional argument\"")
        }
}
