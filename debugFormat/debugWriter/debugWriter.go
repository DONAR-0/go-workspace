package debugWriter

import (
	"os"
	"strconv"

	"github.com/DONAR-0/go-workspace/assertions/pkg/utils"
	"github.com/olekukonko/tablewriter"
)

var checkAppend = utils.CheckAppendError
var check = utils.DeferCheck

func T_map(data map[string]int) {
	table := tablewriter.NewTable(os.Stdout)
	table.Header([]string{"Key", "Value"})

	for k, v := range data {
		checkAppend(table.Append, []string{k, strconv.Itoa(v)})
	}

	check(table.Render)
}
