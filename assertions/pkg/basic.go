package pkg

import (
	"reflect"
	"testing"
)

func AssertStringGotWant(t testing.TB, got, want string) {
	assertAny(t, got, want)
}

func assertAny(t testing.TB, got, want any) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, want %q", got, want)
	}
}
