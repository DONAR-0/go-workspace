package utils

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

var check = CheckError
var checkAppend = CheckAppendError

func RenderIOTable(input, output string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.Header([]string{"Input", "Output"})
	checkAppend(table.Append, []string{fmt.Sprintf("%v", input), fmt.Sprintf("%v", output)})
	check(table.Render)
}

func RenderAppStatusTable(data map[string]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.Header([]string{"App", "Status"})

	for k, v := range data {
		checkAppend(table.Append, []string{fmt.Sprintf("%v", k), fmt.Sprintf("%v", v)})
	}

	check(table.Render)
}
