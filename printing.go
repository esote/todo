// Copyright 2018 Esote. All rights reserved. Use of this source code is
// governed by an MIT license that can be found in the LICENSE file.

package todo

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
)

//PrintItems dumps a slice of items, optionally being a bit more verbose
func PrintItems(items []Item, verbose bool) error {
	t := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)

	var header string

	if verbose {
		header = "ID\t| Message\t| Details\t| Category\t| Priority\n" +
			"--\t---------\t---------\t----------\t----------"
	} else {
		header = "ID\t| Message\t| Category\t| Priority\n" +
			"--\t---------\t----------\t----------"
	}

	fmt.Fprintln(t, header)

	for _, i := range items {
		i.Message = strings.Replace(i.Message, "\n", "\\n", -1)

		if verbose {
			i.Details = strings.Replace(i.Details, "\n", "\\n", -1)
		}

		i.Category = strings.Replace(i.Category, "\n", "\\n", -1)

		if verbose {
			fmt.Fprintf(t, "%d\t| %s\t| %s\t| %s\t| %d\n",
				i.ID, i.Message, i.Details, i.Category, i.Priority)
		} else {
			fmt.Fprintf(t, "%d\t| %s\t| %s\t| %d\n",
				i.ID, i.Message, i.Category, i.Priority)
		}
	}

	return t.Flush()
}

//PrintDetailed dumps info from item
func PrintDetailed(i Item) error {
	t := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

	i.Message = strings.Replace(i.Message, "\n", "\n\t| ", -1)
	i.Details = strings.Replace(i.Details, "\n", "\n\t| ", -1)
	i.Category = strings.Replace(i.Category, "\n", "\n\t| ", -1)

	fmt.Fprintln(t, "ID\t:", i.ID)
	fmt.Fprintln(t, "Message\t:", i.Message)
	fmt.Fprintln(t, "Details\t:", i.Details)
	fmt.Fprintln(t, "Category\t:", i.Category)
	fmt.Fprintln(t, "Priority\t:", i.Priority)

	return t.Flush()
}
