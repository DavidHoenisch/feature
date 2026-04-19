package main_test

import (
	"testing"

	main "github.com/DavidHoenisch/feature"
)

func TestCheckerImpl_CheckStateModel(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		v       string
		want    *main.Flag
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := main.NewFeater()
			got, gotErr := c.CheckStateModel(tt.v)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("CheckStateModel() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("CheckStateModel() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("CheckStateModel() = %v, want %v", got, tt.want)
			}
		})
	}
}
