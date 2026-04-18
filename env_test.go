package main

import (
	"errors"
	"fmt"
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

func TestCastVarToInt(t *testing.T) {
	var tests = []struct {
		name    string
		in      string
		wantInt int
		wantErr error
	}{
		{
			name:    "",
			in:      "1",
			wantInt: 1,
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			i, _ := castVarToInt(tt.in)
			if i != tt.wantInt {
				fmt.Errorf("wanted %v, got %v", tt.wantInt, i)
			}
		})
	}
}

func TestCastVarTofloat(t *testing.T) {
	var tests = []struct {
		name    string
		in      string
		wantFlt float64
		wantErr error
	}{
		{
			name:    "",
			in:      "3.14",
			wantFlt: 3.14,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			i, _ := castVarToFloat64(tt.in)
			if i != tt.wantFlt {
				fmt.Errorf("wanted %v, got %v", tt.wantFlt, i)
			}
		})
	}
}

func TestCastVarToBool(t *testing.T) {
	var tests = []struct {
		name     string
		in       string
		wantBool bool
		wantErr  error
	}{
		{
			name:     "",
			in:       "true",
			wantBool: true,
			wantErr:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			i, _ := castVarToBool(tt.in)
			if i != tt.wantBool {
				fmt.Errorf("wanted %v, got %v", tt.wantBool, i)
			}
		})
	}
}
