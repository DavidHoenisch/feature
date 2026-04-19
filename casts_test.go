package main

import (
	"fmt"
	"testing"
)

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
