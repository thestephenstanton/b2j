package main

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBytesFromArgs(t *testing.T) {
	testCases := []struct {
		desc          string
		osArgs        []string
		expectedBytes []byte
		shouldError   bool
		expectedError string
	}{
		{
			desc:          "standard input",
			osArgs:        []string{"program path", "123", "34", "104", "101", "108", "108", "111", "34", "58", "34", "119", "111", "114", "108", "100", "34", "125"},
			expectedBytes: json.RawMessage(`{"hello":"world"}`),
		},
		{
			desc:          "bad input",
			osArgs:        []string{"program path", "123", "foo", "125"},
			shouldError:   true,
			expectedError: "turning string 'foo' to a int: strconv.Atoi: parsing \"foo\": invalid syntax",
		},
		{
			desc:          "include []",
			osArgs:        []string{"program path", "[123", "34", "104", "101", "108", "108", "111", "34", "58", "34", "119", "111", "114", "108", "100", "34", "125]"},
			expectedBytes: json.RawMessage(`{"hello":"world"}`),
		},
		{
			desc:          "no input",
			osArgs:        []string{"program path"},
			shouldError:   true,
			expectedError: "no bytes passed in",
		},
		{
			desc:          "ignore flag",
			osArgs:        []string{"program path", "-m", "[123", "34", "104", "101", "108", "108", "111", "34", "58", "34", "119", "111", "114", "108", "100", "34", "125]"},
			expectedBytes: json.RawMessage(`{"hello":"world"}`),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			bytes, err := getBytesFromArgs(tc.osArgs)
			if tc.shouldError && err == nil {
				assert.Fail(t, "should have errored but didn't")
			}
			if err != nil {
				assert.EqualError(t, err, tc.expectedError)
			}

			assert.Equal(t, tc.expectedBytes, bytes)
		})
	}
}
