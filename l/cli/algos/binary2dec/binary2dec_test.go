package binary2dec

import (
	"fmt"
	"testing"

	"github.com/DONAR-0/go-workspace/assertions/pkg/tablewriter"
)

func TestBinary2Dec(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  int
	}{
		{name: "Zero Test", input: 0, want: 0},
		{name: "One Test", input: 1, want: 1},
		{name: "Zero One Test", input: 01, want: 1},
		{name: "Zero Zero Test", input: 0000, want: 0},
		{name: "Zero Zero Test", input: 0000, want: 0},
		{name: "Two", input: 10, want: 2},
		{name: "Three", input: 11, want: 3},
		{name: "Seven", input: 111, want: 7},
		{name: "Five", input: 101, want: 5},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d. %s", i, test.name), func(t *testing.T) {
			got := Convert(test.input)
			tablewriter.AssertIntGotWant(t, got, test.want)
		})
	}
}
