// Copyright 2018 Esote. All rights reserved. Use of this source code is
// governed by an MIT license that can be found in the LICENSE file.

package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// Item represents an individual item in the todo list.
type Item struct {
	ID       int    `json:"id"`
	Message  string `json:"message"`
	Details  string `json:"detail"`
	Category string `json:"category"`
	Priority int    `json:"priority"`
}

func readJSON(name string) ([]Item, error) {
	b, err := ioutil.ReadFile(name)

	if err != nil {
		return nil, err
	}

	var items []Item
	err = json.Unmarshal(b, &items)

	return items, err
}

func writeJSON(items []Item, name string) (err error) {
	var b []byte

	b, err = json.Marshal(items)

	if err != nil {
		return
	}

	return ioutil.WriteFile(name, b, 0600)
}

func readItem() (i Item, err error) {
	r := bufio.NewReader(os.Stdin)

	// Message.
	fmt.Print("Message: ")
	i.Message, err = r.ReadString('\n')

	if err != nil {
		return
	}

	i.Message = strings.Trim(i.Message, "\n")

	// Details.
	fmt.Println("Details (reads until 'END'):")
	i.Details, err = readDetails(r)

	if err != nil {
		return
	}

	i.Details = strings.Trim(i.Details, "\n")

	// Category.
	fmt.Print("Category: ")
	i.Category, err = r.ReadString('\n')

	if err != nil {
		return
	}

	i.Category = strings.Trim(i.Category, "\n")

	// Priority.
	fmt.Print("Priority: ")
	tmp, err := r.ReadString('\n')

	if err != nil {
		return
	}

	tmp = strings.Trim(tmp, "\n")

	// Use Atoi over Scanf allows easier file input through standard input.
	i.Priority, err = strconv.Atoi(tmp)

	return
}

func readDetails(r *bufio.Reader) (string, error) {
	var b bytes.Buffer
	var l string
	var err error

	for {
		l, err = r.ReadString('\n')

		if err != nil || l == "END\n" {
			break
		}

		_, err = b.WriteString(l)

		if err != nil {
			break
		}
	}

	return b.String(), err
}
