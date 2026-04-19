package main

import (
	"errors"
	"os"
	"testing"
)

func TestGetFlagFromEnv(t *testing.T) {
	var tests = []struct {
		name  string
		key   string
		value string
		err   error
	}{
		{
			name:  "Var not set",
			key:   "",
			value: "",
			err:   KeyNotFoundErr,
		},
		{
			name:  "var set but empty",
			key:   "FOO",
			value: "",
			err:   KeyPresentButEmptyErr,
		},
		{
			name:  "var set and filled",
			key:   "FOO",
			value: "BAR",
			err:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Setenv(tt.key, tt.value)

			_, err := getFlagFromEnv(tt.key)
			if !errors.Is(err, tt.err) {
				t.Errorf("got %s, wanted %s", err, tt.err)
			}
		})
	}
}
