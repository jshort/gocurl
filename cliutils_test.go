package main

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

func TestNewCliOptions(t *testing.T) {
        // Happy Path
        options := newCliOptions()
        options.httpVerb = "POST"
        options.httpHeaders = []string{"Content-Type: application/json"}
        options.arguments = []string{"http://localhost:8080/something"}

        assert.Equal(t, options.url(), "http://localhost:8080/something", "Url should be first argument")
        assert.Equal(t, options.httpVerb, "POST", "HTTP Verb should be POST")
        assert.Equal(t, options.timeout, 60, "Timeout should be 60")
        assert.Equal(t, len(options.httpHeaders), 1, "Should be 1 Header")
        assert.Equal(t, options.httpHeaders[0], "Content-Type: application/json", "Header should have application/json")
}

func TestValidateOptions(t *testing.T) {
        // Happy Path
        options := newCliOptions()

        options.httpHeaders = []string{"Key1:Value1", "Key2:Value2"}

        code, _ := validateOptions(options)

        assert.Zero(t, code, "Return code should be 0")

        // Error Path
        options.httpVerb = "GOT"

        code, message := validateOptions(options)

        assert.Equal(t, code, 1, "Return code should be 1")
        assert.True(t, strings.Contains(message, "Unexpected HTTP Verb"),
                "Error message should contain \"Unexpected HTTP Verb\"")

        options.httpVerb = "GET"
        options.httpHeaders = []string{"Key1 Value1"}

        code, message = validateOptions(options)

        assert.Equal(t, code, 1, "Return code should be 1")
        assert.True(t, strings.Contains(message, "Unexpected HTTP Header"),
                "Error message should contain \"Unexpected HTTP Header\"")

}

func TestValidateArguments(t *testing.T) {
        // Happy Path
        options := newCliOptions()

        options.arguments = []string{"http://google.com"}

        code, _ := validateArguments(options)

        assert.Zero(t, code, "Return code should be 0")

        // Error Path
        options.arguments = []string{"http://google.com", "http://localhost:8080"}

        code, message := validateArguments(options)

        assert.Equal(t, code, 2, "Return code should be 2")
        assert.True(t, strings.Contains(message, "Expecting 1 positional argument"),
                "Error message should contain \"Expecting 1 positional argument\"")
}
