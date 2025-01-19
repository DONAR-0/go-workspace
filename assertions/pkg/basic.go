package pkg

import (
	"testing"
)

func AssertStringGotWant(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
