package assert

import (
	"testing"
)

func True(t *testing.T, expected bool) {
	if !expected {
		t.Errorf("expected to be true but actual: %v", expected)
	}
}

func False(t *testing.T, expected bool) {
	if expected {
		t.Errorf("expected to be false but actual: %v", expected)
	}
}

func Empty[S ~[]E, E any](t *testing.T, haystack S) {
	size := len(haystack)
	if size != 0 {
		t.Errorf("expected object to be empty but got: %d", size)
	}
}

func NotEmpty[S ~[]E, E any](t *testing.T, haystack S) {
	size := len(haystack)
	if size == 0 {
		t.Errorf("expected object not to be empty but got: %d", size)
	}
}
