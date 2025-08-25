package tablewriter

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
	"testing"

	"github.com/DONAR-0/go-workspace/assertions/pkg/utils"
	"github.com/olekukonko/tablewriter"
)

var check = utils.CheckError
var checkAppend = utils.CheckAppendError

func AssertStringGotWant(t testing.TB, got, want string) {
	t.Helper()

	table := tablewriter.NewWriter(os.Stdout)
	table.Header([]string{"Got", "Want"})
	utils.CheckAppendError(table.Append, []string{got, want})

	if got != want {
		check(table.Render)
		t.Errorf("TEST FAILED")
	} else {
		check(table.Render)
	}
}

func AssertStringSliceGotWant(t testing.TB, got, want []string) {
	t.Helper()

	table := tablewriter.NewWriter(os.Stdout)
	table.Header([]string{"Field", "Got", "Want"})

	checkAppend(table.Append, []string{"Lenght", strconv.Itoa(len(got)), strconv.Itoa(len(want))})
	checkAppend(table.Append, []string{"Values", fmt.Sprintf("%v", got), fmt.Sprintf("%v", want)})

	if !reflect.DeepEqual(got, want) {
		check(table.Render)
		t.Errorf("TEST FAILED")
	} else {
		check(table.Render)
	}
}

func AssertFloatGotWant(t testing.TB, got, want float64) {
	t.Helper()

	table := tablewriter.NewWriter(os.Stdout)
	table.Header([]string{"Got", "Want"})

	checkAppend(table.Append, []string{fmt.Sprintf("%.2f", got), fmt.Sprintf("%.2f", want)})

	if got != want {
		check(table.Render)
		t.Errorf("TEST FAILED")
	} else {
		check(table.Render)
	}
}

func AssertIntGotWant(t testing.TB, got, want int) {
	t.Helper()

	table := tablewriter.NewWriter(os.Stdout)
	table.Header([]string{"Got", "Want"})
	checkAppend(table.Append, []string{fmt.Sprintf("%d", got), fmt.Sprintf("%d", want)})

	if got != want {
		check(table.Render)
		t.Errorf("TEST FAILED")
	} else {
		check(table.Render)
	}
}

func AssertStructGotWant(t testing.TB, got, want any) {
	t.Helper()

	table := tablewriter.NewWriter(os.Stdout)
	table.Header([]string{"Got", "Want"})
	checkAppend(table.Append, []string{fmt.Sprintf("%v", got), fmt.Sprintf("%v", want)})

	if got != want {
		check(table.Render)
		t.Errorf("TEST FAILED")
	} else {
		check(table.Render)
	}
}

func AssertError(t testing.TB, got error, want string) {
	t.Helper()

	table := tablewriter.NewWriter(os.Stdout)
	table.Header([]string{"Got Error Message", "Want Error Message"})

	if got == nil {
		checkAppend(table.Append, []string{fmt.Sprintf("%v", "ERROR IS NIL"), fmt.Sprintf("%v", want)})
		t.Errorf("Expected not nil but found nil")
	} else {
		if strings.Contains(got.Error(), want) {
			checkAppend(table.Append, []string{fmt.Sprintf("%v", got.Error()), fmt.Sprintf("%v", want)})
			check(table.Render)
		} else {
			checkAppend(table.Append, []string{fmt.Sprintf("%v", got.Error()), fmt.Sprintf("%v", want)})
			check(table.Render)
			t.Errorf("TEST FAILED")
		}
	}
}

func AssertErrorNil(t testing.TB, got error) {
	t.Helper()

	table := tablewriter.NewWriter(os.Stdout)
	table.Header([]string{"Error Message"})

	if got == nil {
		checkAppend(table.Append, []string{fmt.Sprintf("%v", "ERROR IS NIL")})
	}
}

func AssertErrorType(t testing.TB, got, want error) {
	t.Helper()

	table := tablewriter.NewWriter(os.Stdout)
	table.Header([]string{"Got", "Want"})

	if got != want {
		checkAppend(table.Append, []string{fmt.Sprintf("%v", got), fmt.Sprintf("%v", want)})
		t.Errorf("TEST FAILED")
	}
}
