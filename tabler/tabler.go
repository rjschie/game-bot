package tabler

import (
	"fmt"
	"os"
	"reflect"

	"github.com/olekukonko/tablewriter"
)

// RenderColumns .
func RenderColumns(header []string, cols ...[]interface{}) {
	table := tablewriter.NewWriter(os.Stdout)

	table.SetHeader(header)
	appendColumns(
		table,
		cols...,
	)
	table.Render()
}

// RenderRows .
func RenderRows(header []string, rows ...[]string) {
	table := tablewriter.NewWriter(os.Stdout)

	table.SetHeader(header)
	table.AppendBulk(rows)
	table.Render()
}

// Columnize .
func Columnize(col interface{}) []interface{} {
	colR := reflect.ValueOf(col)

	intf := make([]interface{}, colR.Len())
	for i := 0; i < colR.Len(); i++ {
		intf[i] = colR.Index(i).Interface()
	}

	return intf
}

func appendColumns(table *tablewriter.Table, cols ...[]interface{}) {
	var ln int
	for _, col := range cols {
		if colLen := len(col); colLen > ln {
			ln = colLen
		}
	}

	for i := 0; i < ln; i++ {
		row := []string{}
		for j := 0; j < len(cols); j++ {
			if len(cols[j]) > i {
				row = append(row, fmt.Sprint(cols[j][i]))
			} else {
				row = append(row, " ")
			}
		}
		table.Append(row)
	}
}
