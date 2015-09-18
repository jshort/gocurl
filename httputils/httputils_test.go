package httputils

import (
        "testing"
        "errors"
        "github.com/stretchr/testify/assert"
)

func TestParseHeaderString(t *testing.T) {
        headers := []string{"Key1:Value1", "Key2:Value2"}
        parsedHeaderMap := parseHeaderString(headers)

        assert.NotNil(t, parsedHeaderMap, "Parsed map should not be nil")

        for key, value := range parsedHeaderMap {

                switch key {
                case "Key1":
                        assert.Equal(t, value, "Value1", "Key1 should have Value1")  
                case "Key2":
                        assert.Equal(t, value, "Value2", "Key2 should have Value2")
                }
        }

}

func TestErrorHandling(t *testing.T) {
        err := errors.New("malformed HTTP response")

        retval, _ := handleHttpError(err)
        assert.Equal(t, retval, 1)

        err = errors.New("tls: oversized record received")

        retval, _ = handleHttpError(err)
        assert.Equal(t, retval, 2)

        err = errors.New("lookup blahblah.com: no such host")

        retval, msg := handleHttpError(err)
        assert.Equal(t, retval, 5)
        assert.Contains(t, msg, "blahblah.com")

        err = errors.New("blah blah")

        retval, _ = handleHttpError(err)
        assert.Equal(t, retval, 255)



}