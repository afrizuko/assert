package assert

import (
	"strings"
	"testing"
)

func Equal[T comparable](t *testing.T, actual, expected T) {
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

func ElementsMatch[S ~[]E, E comparable](t *testing.T, actual S, expected S) {
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

func Greater[E Ordered](t *testing.T, actual, target E) {
	if actual <= target {
		t.Errorf("actual: %v should be greater than target: %v", actual, target)
	}
}

func GreaterOrEqual[E Ordered](t *testing.T, actual, target E) {
	if actual < target {
		t.Errorf("actual: %v should be greater than or equal to target: %v", actual, target)
	}
}

func LessOrEqual[E Ordered](t *testing.T, actual E, target E) {
	if actual > target {
		t.Errorf("actual: %v should be less than or equal to target: %v", actual, target)
	}
}

func LessThan[E Ordered](t *testing.T, actual E, target E) {
	if actual >= target {
		t.Errorf("actual: %v should be greater than target: %v", actual, target)
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
