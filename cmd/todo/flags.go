// Copyright 2018 Esote. All rights reserved. Use of this source code is
// governed by an MIT license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"github.com/esote/todo"
	"os"
)

// Flags represents the possible command line flags.
type flags struct {
	ID int // -n

	item todo.Item // -amess -adet -acat -apr

	verbose bool // -v
}

var (
	fl flags

	dfl flags // default flag values
)

func init() {
	flag.IntVar(&fl.ID, "n", dfl.ID,
		`the ID of the item to be operated on. When used with {edit} and
{delete} it specifies the item which will be respectively overwritten or
deleted. When used with {view} it instructs todo to show all an item's data.`)

	flag.StringVar(&fl.item.Message, "amess", dfl.item.Message,
		`append message. Used with {append} to specify the new item's message.
Used with {edit} to overwrite an item's message.`)

	flag.StringVar(&fl.item.Details, "adet", dfl.item.Details,
		`append details. Used with {append} to specify the new item's details.
Used with {edit} to overwrite an item's details.`)

	flag.StringVar(&fl.item.Category, "acat", dfl.item.Category,
		`append category. Used with {append} to specify the new item's category.
Used with {edit} to overwrite an item's category.`)

	flag.IntVar(&fl.item.Priority, "apr", dfl.item.Priority,
		`append priority. Used with {append} to specify the new item's priority.
Used with {edit} to overwrite an item's priority.`)

	flag.BoolVar(&fl.verbose, "v", dfl.verbose,
		`verbose output. When used with {view} todo will show all columns.`)
}

func usage() {
	fmt.Fprintf(os.Stderr,
		`Todo is a list management tool.

Usage:

	todo [arguments] command filename

The commands are:

	append    append an item to the list
	delete    delete an item in the list
	edit      edit an existing list item
	init      initialize a new todo list
	view      view item(s) from the list

The arguments are:

`)

	flag.PrintDefaults()
}
