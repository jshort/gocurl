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
                got := contains(HttpVerbs[:], c.in)
                if got != c.want {
                        t.Errorf("contains(%q, %q) == %v, want %v", HttpVerbs, c.in, got, c.want)
                }
        }
}

func TestInitializeCli(t *testing.T) {
        // Happy Path
        var cliMap = map[string]interface{}{
                "httpVerb"    : "POST",
                "httpHeaders" : []string{"Content-Type: application/json"},
                "postData"    : "",
                "timeOut"     : 60,
                "arguments"   : []string{"http://localhost:8080/something"},
        }
        cliInputs := InititializeCli(cliMap)
        if cliInputs.Url() != "http://localhost:8080/something" {
                t.Errorf("Url() returned %v, want %v", cliInputs.Url(), "http://localhost:8080/something")
        }
        if cliInputs.HttpVerb() != "POST" {
                t.Errorf("HttpVerb() returned %v, want %v", cliInputs.HttpVerb(), "POST")
        }
        if cliInputs.TimeOut() != 60 {
                t.Errorf("TimeOut() returned %v, want %v", cliInputs.TimeOut(), 60)
        }
        if len(cliInputs.HttpHeaders()) != 1 && cliInputs.HttpHeaders()[0] != "Content-Type: application/json" {
                t.Errorf("HttpHeaders() returned %v, want %v", cliInputs.HttpHeaders(), "Content-Type: application/json")
        }



}

func TestValidateOptions(t *testing.T) {
        // Happy Path
        var cliMap = map[string]interface{}{
                "httpVerb"    : "GET",
                "httpHeaders" : []string{},
                "postData"    : "",
                "timeOut"     : 60,
                "arguments"   : []string{},
        }
        cliInputs := InititializeCli(cliMap)
        code, _ := ValidateOptions(cliInputs)

        if code != 0 {
                t.Errorf("ValidateOptions(%v) returned %v, want %v", cliMap, code, 0)
        }


        // Error Path
        cliMap["httpVerb"] = "GOT"
        cliInputs = InititializeCli(cliMap)
        code, message := ValidateOptions(cliInputs)

        if code != 1 {
                t.Errorf("ValidateOptions(%v) returned %v, want %v", cliMap, code, "non zero")
        }
        if ! strings.Contains(message, "Unexpected HTTP Verb") {
                t.Error("Error message should contain \"Unexpected HTTP Verb\"")
        }
}

func TestValidateArguments(t *testing.T) {
        // Happy Path
        var cliMap = map[string]interface{}{
                "httpVerb"    : "GET",
                "httpHeaders" : []string{},
                "postData"    : "",
                "timeOut"     : 60,
                "arguments"   : []string{"http://google.com"},
        }
        cliInputs := InititializeCli(cliMap)
        code, _ := ValidateArguments(cliInputs)

        if code != 0 {
                t.Errorf("ValidateArguments(%v) returned %v, want %v", cliMap["arguments"], code, 0)
        }

        // Error Path
        cliMap["arguments"] = []string{"http://google.com", "http://localhost:8080"}
        cliInputs = InititializeCli(cliMap)
        code, message := ValidateArguments(cliInputs)

        if code != 2 {
                t.Errorf("ValidateArguments(%v) returned %v, want %v", cliMap["arguments"], code, "non zero")
        }
        if ! strings.Contains(message, "Expecting 1 positional argument") {
                t.Error("Error message should contain \"Expecting 1 positional argument\"")
        }
}
