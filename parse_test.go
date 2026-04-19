package main

import (
	"fmt"
	"testing"
)

func TestParseFlagToModel(t *testing.T) {
	var tests = []struct {
		name string
		in   string
		want *Flag
	}{
		{
			name: "",
			in:   "op:something:true",
			want: &Flag{
				Category: "op",
				Feature:  "something",
				State:    true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			m, _ := parseFlagToModel(tt.in)
			if m != tt.want {
				fmt.Errorf("wanted %v, got %v", tt.want, m)
			}
		})
	}
}
