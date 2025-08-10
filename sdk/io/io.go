// Package inout provides functions to handle basic input output operations
package inout

import (
	"errors"
	"os"
)

func SaveImage(name string, data *[]byte) error {

	f, err := os.Create(name)
	if nil != err {
		return err
	}

	defer f.Close()

	n, err := f.Write(*data)
	if nil != err {
		return err
	}

	if n != len(*data) {
		return errors.New("Cannot write all the file.")
	}

	return nil
}
