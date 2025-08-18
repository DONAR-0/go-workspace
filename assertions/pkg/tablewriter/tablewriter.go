package tablewriter

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/olekukonko/tablewriter"
)

func AssertStringGotWant(t testing.TB, got, want string) {
	t.Helper()

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Got", "Want"})
	table.Append([]string{got, want})

	if got != want {
		table.Render()
		t.Errorf("TEST FAILED")
	} else {
		table.Render()
	}
}

func AssertFloatGotWant(t testing.TB, got, want float64) {
	t.Helper()

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Got", "Want"})
	table.Append([]string{fmt.Sprintf("%.2f", got), fmt.Sprintf("%.2f", want)})

	if got != want {
		table.Render()
		t.Errorf("TEST FAILED")
	} else {
		table.Render()
	}
}

func AssertStructGotWant(t testing.TB, got, want any) {
	t.Helper()

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Got", "Want"})
	table.Append([]string{fmt.Sprintf("%v", got), fmt.Sprintf("%v", want)})

	if got != want {
		table.Render()
		t.Errorf("TEST FAILED")
	} else {
		table.Render()
	}
}

func AssertError(t testing.TB, got error, want string) {
	t.Helper()

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Got Error Message", "Want Error Message"})

	if got == nil {
		table.Append([]string{fmt.Sprintf("%v", "ERROR IS NIL"), fmt.Sprintf("%v", want)})
		t.Errorf("Expected not nil but found nil")
	} else {
		if strings.Contains(got.Error(), want) {
			table.Append([]string{fmt.Sprintf("%v", got.Error()), fmt.Sprintf("%v", want)})
			table.Render()
		} else {
			table.Append([]string{fmt.Sprintf("%v", got.Error()), fmt.Sprintf("%v", want)})
			table.Render()
			t.Errorf("TEST FAILED")
		}
	}
}

func AssertErrorNil(t testing.TB, got error) {
	t.Helper()

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Error Message"})

	if got == nil {
		table.Append([]string{fmt.Sprintf("%v", "ERROR IS NIL")})
	}
}

func AssertErrorType(t testing.TB, got, want error) {
	t.Helper()

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Got", "Want"})

	if got != want {
		table.Append([]string{fmt.Sprintf("%v", got), fmt.Sprintf("%v", want)})
		t.Errorf("TEST FAILED")
	}
}
