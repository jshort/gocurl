package httputils

import (
        "testing"
        "github.com/stretchr/testify/assert"
)

func TestParseHeaderString(t *testing.T) {
        headers := []string{"Key1:Value1", "Key2:Value2"}
        parsedHeaderMap, ok := parseHeaderString(headers)

        assert.NotNil(t, parsedHeaderMap, "Parsed map should not be nil")
        assert.True(t, ok, "Parsing failed")

        for key, value := range parsedHeaderMap {

                switch key {
                case "Key1":
                        assert.Equal(t, value, "Value1", "Key1 should have Value1")  
                case "Key2":
                        assert.Equal(t, value, "Value2", "Key2 should have Value2")
                }
        }

        headers = []string{"Key1:Value1:invalid", "Key2:Value2"}
        parsedHeaderMap, ok = parseHeaderString(headers)

        assert.Nil(t, parsedHeaderMap, "Parsed map should be nil")
        assert.False(t, ok, "Parsing succeeded when it should have failed")
}