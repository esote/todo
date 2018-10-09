// Copyright 2018 Esote. All rights reserved. Use of this source code is
// governed by an MIT license that can be found in the LICENSE file.

package todo

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

//ReadJSON converts json into Item slice
func ReadJSON(name string) ([]Item, error) {
	b, err := ioutil.ReadFile(name)

	if err != nil {
		return nil, err
	}

	var items []Item
	err = json.Unmarshal(b, &items)

	return items, err
}

//WriteJSON writes Item slice to file
func WriteJSON(items []Item, name string) (err error) {
	var b []byte

	b, err = json.Marshal(items)

	if err != nil {
		return
	}

	return ioutil.WriteFile(name, b, 0600)
}

//ReadItem fills item with input from stderr
func ReadItem() (i Item, err error) {
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
	i.Details, err = ReadDetails(r)

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

//ReadDetails reads details from a specfied Reader
func ReadDetails(r *bufio.Reader) (string, error) {
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

func InitJSON(name string) error {
	if _, err := os.Stat(name); !os.IsNotExist(err) {
		if err == nil {
			return fmt.Errorf("file %s already exists", name)
		} else {
			return err
		}
	}

	b := []byte("[]\n")

	return ioutil.WriteFile(name, b, 0600)
}