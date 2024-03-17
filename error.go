package assert

import (
	"testing"
)

func Error(t *testing.T, err error) {
	if err != nil {
		t.Errorf("expected no error but got %v", err)
		t.Fail()
	}
}

func NoError(t *testing.T, err error) {
	if err == nil {
		t.Error("expected error but got nil")
		t.Fail()
	}
}

func Nil[T any](t *testing.T, object *T) {
	if object != nil {
		t.Errorf("expected object to be nil but got %v", object)
		t.Fail()
	}
}

func NotNil[T any](t *testing.T, object *T) {
	if object == nil {
		t.Error("expected object not to be nil")
		t.Fail()
	}
}
