package cliutils

import (
        "testing"
        "strings"
        "github.com/stretchr/testify/assert"
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
                "verbose"     : false,
                "color"       : false,
                "arguments"   : []string{"http://localhost:8080/something"},
        }
        cliInputs := InititializeCli(cliMap)

        assert.Equal(t, cliInputs.Url(), "http://localhost:8080/something", "Url should be first argument")
        assert.Equal(t, cliInputs.HttpVerb(), "POST", "HTTP Verb should be POST")
        assert.Equal(t, cliInputs.TimeOut(), 60, "Timeout should be 60")
        assert.Equal(t, len(cliInputs.HttpHeaders()), 1, "Should be 1 Header")
        assert.Equal(t, cliInputs.HttpHeaders()[0], "Content-Type: application/json", "Header should have application/json")
}

func TestValidateOptions(t *testing.T) {
        // Happy Path
        var cliMap = map[string]interface{}{
                "httpVerb"    : "GET",
                "httpHeaders" : []string{"Key1:Value1", "Key2:Value2"},
                "postData"    : "",
                "timeOut"     : 60,
                "verbose"     : false,
                "color"       : false,
                "arguments"   : []string{},
        }
        cliInputs := InititializeCli(cliMap)
        code, _ := ValidateOptions(cliInputs)

        assert.Zero(t, code, "Return code should be 0")

        // Error Path
        cliMap["httpVerb"] = "GOT"
        cliInputs = InititializeCli(cliMap)
        code, message := ValidateOptions(cliInputs)

        assert.Equal(t, code, 1, "Return code should be 1")
        assert.True(t, strings.Contains(message, "Unexpected HTTP Verb"),
                "Error message should contain \"Unexpected HTTP Verb\"")

        cliMap["httpVerb"] = "GET"
        cliMap["httpHeaders"] = []string{"Key1 Value1"}
        cliInputs = InititializeCli(cliMap)
        code, message = ValidateOptions(cliInputs)

        assert.Equal(t, code, 1, "Return code should be 1")
        assert.True(t, strings.Contains(message, "Unexpected HTTP Header"),
                "Error message should contain \"Unexpected HTTP Header\"")

}

func TestValidateArguments(t *testing.T) {
        // Happy Path
        var cliMap = map[string]interface{}{
                "httpVerb"    : "GET",
                "httpHeaders" : []string{},
                "postData"    : "",
                "timeOut"     : 60,
                "verbose"     : false,
                "color"       : false,
                "arguments"   : []string{"http://google.com"},
        }
        cliInputs := InititializeCli(cliMap)
        code, _ := ValidateArguments(cliInputs)

        assert.Zero(t, code, "Return code should be 0")

        // Error Path
        cliMap["arguments"] = []string{"http://google.com", "http://localhost:8080"}
        cliInputs = InititializeCli(cliMap)
        code, message := ValidateArguments(cliInputs)

        assert.Equal(t, code, 2, "Return code should be 2")
        assert.True(t, strings.Contains(message, "Expecting 1 positional argument"),
                "Error message should contain \"Expecting 1 positional argument\"")
}
