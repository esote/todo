// Copyright 2018 Esote. All rights reserved. Use of this source code is
// governed by an MIT license that can be found in the LICENSE file.

// Todo is a list management tool.
package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/esote/todo"
	"log"
	"strings"
)

var (
	errMissingArguments = errors.New("missing command or filename")
	errMissingID        = errors.New("missing item ID")
)

func main() {
	flag.Usage = usage
	flag.Parse()

	if flag.NArg() < 1 {
		usage()
		return
	} else if flag.NArg() < 2 {
		log.Fatal(errMissingArguments)
	}

	cmd := strings.ToLower(flag.Args()[0])
	name := flag.Args()[1]

	if err := parseCommand(cmd, name); err != nil {
		log.Fatal(err)
	}
}

func parseCommand(cmd, name string) (err error) {
	var items []todo.Item

	if items, err = todo.ReadJSON(name); err != nil {
		return
	}

	switch cmd {
	case "append":
		var i todo.Item

		if fl.item == dfl.item {
			if i, err = todo.ReadItem(); err != nil {
				return
			}
		} else {
			i = fl.item
		}

		i.ID = todo.NextID(items)

		items = append(items, i)

		if err = todo.WriteJSON(items, name); err != nil {
			return
		}

	case "edit":
		if fl.ID == dfl.ID {
			return errMissingID
		}

		var index int

		if index, err = todo.FindItem(items, fl.ID); err != nil {
			return
		}

		todo.PrintDetailed(items[index])

		msg := "Are you sure you want to overwrite this item?"
		if err = todo.Confirm(msg); err != nil {
			return
		}

		ID := items[index].ID

		if fl.item == dfl.item {
			if items[index], err = todo.ReadItem(); err != nil {
				return
			}
		} else {
			items[index] = fl.item
		}

		items[index].ID = ID

		if err = todo.WriteJSON(items, name); err != nil {
			return
		}

	case "delete":
		var index int

		if index, err = todo.FindItem(items, fl.ID); err != nil {
			return
		}

		todo.PrintDetailed(items[index])

		msg := "Are you sure you want to delete this item?"
		if err = todo.Confirm(msg); err != nil {
			return
		}

		items = append(items[:index], items[index+1:]...)

		if err = todo.WriteJSON(items, name); err != nil {
			return
		}

	case "view":
		if fl.ID == dfl.ID {
			todo.PrintItems(items, fl.verbose)
		} else {
			var index int

			if index, err = todo.FindItem(items, fl.ID); err != nil {
				return
			}

			todo.PrintDetailed(items[index])
		}

	default:
		return fmt.Errorf("incorrect command '%s'", cmd)
	}

	return
}
