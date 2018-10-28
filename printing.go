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

	if _, err := fmt.Fprintln(t, header); err != nil {
		return err
	}

	for _, i := range items {
		i.Message = strings.Replace(i.Message, "\n", "\\n", -1)

		if verbose {
			i.Details = strings.Replace(i.Details, "\n", "\\n", -1)
		}

		i.Category = strings.Replace(i.Category, "\n", "\\n", -1)

		if verbose {
			if _, err := fmt.Fprintf(t, "%d\t| %s\t| %s\t| %s\t| %d\n", i.ID,
				i.Message, i.Details, i.Category, i.Priority); err != nil {
				return err
			}
		} else {
			if _, err := fmt.Fprintf(t, "%d\t| %s\t| %s\t| %d\n",
				i.ID, i.Message, i.Category, i.Priority); err != nil {
				return err
			}
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

	if _, err := fmt.Fprintln(t, "ID\t:", i.ID); err != nil {
		return err
	}

	if _, err := fmt.Fprintln(t, "Message\t:", i.Message); err != nil {
		return err
	}

	if _, err := fmt.Fprintln(t, "Details\t:", i.Details); err != nil {
		return err
	}

	if _, err := fmt.Fprintln(t, "Category\t:", i.Category); err != nil {
		return err
	}

	if _, err := fmt.Fprintln(t, "Priority\t:", i.Priority); err != nil {
		return err
	}

	return t.Flush()
}
