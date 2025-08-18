package debugWriter

import (
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
)

func T_map(data map[string]int) {
	table := tablewriter.NewTable(os.Stdout)
	table.Header([]string{"Key", "Value"})
	for k, v := range data {
		table.Append([]string{k, strconv.Itoa(v)})
	}
	table.Render()
}
