package di

import (
	"bytes"
	"testing"

	"github.com/DONAR-0/go-workspace/assertions/pkg/tablewriter"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Chris")
	got := buffer.String()
	want := "Hello, Chris"
	tablewriter.AssertStringGotWant(t, got, want)
}
