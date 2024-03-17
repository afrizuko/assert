package assert

import (
	"strings"
	"testing"
)

func Equal[T comparable](t *testing.T, expected, actual T) {
	if expected != actual {
		t.Errorf("expected: %v differs from actual: %v", expected, actual)
	}
}

func Len[S ~[]E, E comparable](t *testing.T, object S, expected int) {
	size := len(object)
	if size != expected {
		t.Errorf("expected object size to be: %v actual size: %v", expected, size)
	}
}

func ElementsMatch[S ~[]E, E comparable](t *testing.T, expected S, actual S) {
	if len(expected) != len(actual) {
		t.Errorf("elements mismatch: expected: %v actual: %v", expected, actual)
		return
	}

	mapped := make(map[E]struct{})
	for _, e := range expected {
		mapped[e] = struct{}{}
	}

	for _, a := range actual {
		if _, found := mapped[a]; !found {
			t.Errorf("expected elements to match:%v elements mismatch: %v", expected, actual)
			return
		}
	}
}

func Greater[E Ordered](t *testing.T, expected E, actual E) {
	if expected <= actual {
		t.Errorf("expected: %v to be greater than: %v", expected, actual)
	}
}

func GreaterOrEqual[E Ordered](t *testing.T, expected E, actual E) {
	if expected < actual {
		t.Errorf("expected: %v to be greater than or equal to: %v", expected, actual)
	}
}

func LessOrEqual[E Ordered](t *testing.T, expected E, actual E) {
	if expected > actual {
		t.Errorf("expected: %v to be less than or equal to: %v", expected, actual)
	}
}

func Less[E Ordered](t *testing.T, expected E, actual E) {
	if expected >= actual {
		t.Errorf("expected: %v to be less than: %v", expected, actual)
	}
}

func Contains[S ~[]E, E Ordered](t *testing.T, haystack S, needle E) {
	found := false
	for _, item := range haystack {
		if item == needle {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("expected object: %v to be in : %v", needle, haystack)
	}
}

func ContainsString(t *testing.T, haystack string, needle string) {
	if !strings.Contains(haystack, needle) {
		t.Errorf("expected object: %v to contain: %v", needle, haystack)
	}
}
