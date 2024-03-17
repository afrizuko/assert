package assert

import (
	"testing"
)

func TestContains(t *testing.T) {
	Contains(t, []string{"Google"}, "Google")
}
