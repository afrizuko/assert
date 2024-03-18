package assert

import (
	"fmt"
	"testing"
)

func TestError(t *testing.T) {
	tests := []struct {
		name string
		args error
	}{
		{
			name: "expected an error",
			args: fmt.Errorf("this is an error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Error(t, tt.args)
		})
	}
}

func TestNil(t *testing.T) {
	tests := []struct {
		name string
		args *any
	}{
		{
			name: "expected nil",
			args: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Nil(t, tt.args)
		})
	}
}

func TestNoError(t *testing.T) {
	tests := []struct {
		name string
		args error
	}{
		{
			name: "expected no error",
			args: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			NoError(t, tt.args)
		})
	}
}

func TestNotNil(t *testing.T) {
	tests := []struct {
		name string
		args *any
	}{
		{
			name: "expects item to not be nil",
			args: ptr(1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			NotNil(t, tt.args)
		})
	}
}

func ptr(a any) *any {
	return &a
}
