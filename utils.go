// Copyright 2018 Esote. All rights reserved. Use of this source code is
// governed by an MIT license that can be found in the LICENSE file.

package todo

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"unicode"
)

//Confirm wraps a message and asks for confirm
func Confirm(msg string) error {
	r := bufio.NewReader(os.Stdin)

	fmt.Print(msg + " [y/N]: ")

	choice, err := r.ReadString('\n')

	if err != nil {
		return err
	}

	if unicode.ToLower(rune(choice[0])) != 'y' {
		return errors.New("operation aborted")
	}

	return nil
}

// FindItem finds an item with ID
func FindItem(items []Item, ID int) (index int, err error) {
	var ok bool

	for n, i := range items {
		if i.ID == ID {
			index = n
			ok = true
			break
		}
	}

	if !ok {
		err = fmt.Errorf("no such item with ID '%d'", ID)
	}

	return
}

// NextID gets the next available ID
func NextID(items []Item) int {
	used := make(map[int]bool)

	for _, i := range items {
		used[i.ID] = true
	}

	ID := 1

	for used[ID] {
		ID++
	}

	return ID
}
