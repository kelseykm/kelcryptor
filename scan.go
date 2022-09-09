package main

import (
	"fmt"

	"github.com/kelseykm/kelcryptor/colour"
)

type mismatchedPassword struct{}

func (m mismatchedPassword) Error() string {
	return "Passwords do not match"
}

// TODO: Find better way to read from stdin
func ScanPassword() (string, error) {
	var password string
	var passwordConfirm string

	fmt.Printf("%vEnter password: %s", colour.GreenBold, colour.Normal)
	fmt.Scanln(&password)

	fmt.Printf("%vRepeat password: %s", colour.GreenBold, colour.Normal)
	fmt.Scanln(&passwordConfirm)

	if password == passwordConfirm {
		return password, nil
	} else {
		return "", mismatchedPassword{}
	}
}
