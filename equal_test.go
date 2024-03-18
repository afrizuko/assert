package assert

import (
	"testing"
)

func TestContains(t *testing.T) {
	tests := []struct {
		name     string
		haystack []int
		needle   int
	}{
		{
			name:     "should contain 1",
			haystack: []int{1, 2, 3},
			needle:   1,
		},
		{
			name:     "should contain 3",
			haystack: []int{1, 2, 3},
			needle:   3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Contains(t, tt.haystack, tt.needle)
		})
	}
}

func TestContainsString(t *testing.T) {
	tests := []struct {
		name     string
		haystack string
		needle   string
	}{
		{
			name:     "should contain substring",
			haystack: "Assertions",
			needle:   "tions",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ContainsString(t, tt.haystack, tt.needle)
		})
	}
}

func TestElementsMatch(t *testing.T) {
	tests := []struct {
		name     string
		expected []string
		actual   []string
	}{
		{
			name:     "elements should match",
			expected: []string{"A", "B", "C"},
			actual:   []string{"B", "C", "A"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ElementsMatch(t, tt.expected, tt.actual)
		})
	}
}

func TestEqual(t *testing.T) {
	tests := []struct {
		name     string
		expected int
		actual   int
	}{
		{
			name:     "should be equal",
			expected: 10,
			actual:   10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Equal(t, tt.expected, tt.actual)
		})
	}
}

func TestGreater(t *testing.T) {
	tests := []struct {
		name     string
		expected int
		actual   int
	}{
		{
			name:     "should be greater than",
			expected: 10,
			actual:   11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Greater(t, tt.actual, tt.expected)
		})
	}
}

func TestGreaterOrEqual(t *testing.T) {
	tests := []struct {
		name     string
		expected int
		actual   int
	}{
		{
			name:     "should be greater than or equal",
			expected: 11,
			actual:   12,
		},
		{
			name:     "should be greater than ",
			expected: 10,
			actual:   10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GreaterOrEqual(t, tt.actual, tt.expected)
		})
	}
}

func TestLen(t *testing.T) {
	tests := []struct {
		name     string
		object   []int
		expected int
	}{
		{
			name:     "len should be equal to 3",
			object:   []int{2, 8, 0},
			expected: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Len(t, tt.object, tt.expected)
		})
	}
}

func TestLessThan(t *testing.T) {
	tests := []struct {
		name     string
		expected int32
		actual   int32
	}{
		{
			name:     "expected should be less than actual",
			expected: 10,
			actual:   11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			LessThan(t, tt.expected, tt.actual)
		})
	}
}

func TestLessOrEqual(t *testing.T) {
	tests := []struct {
		name     string
		expected int
		actual   int
	}{
		{
			name:     "expected should be less or equal to actual",
			expected: 10,
			actual:   11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			LessOrEqual(t, tt.expected, tt.actual)
		})
	}
}
