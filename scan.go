package main

import (
	"bytes"
	"fmt"
	"syscall"

	"github.com/kelseykm/kelcryptor/colour"
	"github.com/kelseykm/kelcryptor/errors"
	"golang.org/x/term"
)

func scanPassword() ([]byte, error) {

	fmt.Printf("%s %s",
		colour.Input(),
		colour.Message("Enter password: "),
	)
	password, err := term.ReadPassword(int(syscall.Stdin))
	fmt.Println()
	if err != nil {
		return nil, errors.GenericError{err.Error()}
	}

	fmt.Printf("%s %s",
		colour.Input(),
		colour.Message("Repeat password: "),
	)
	passwordConfirm, err := term.ReadPassword(int(syscall.Stdin))
	fmt.Println()
	if err != nil {
		return nil, errors.GenericError{err.Error()}
	}

	if !bytes.Equal(password, passwordConfirm) {
		return nil, errors.MismatchedPasswordError{}
	}

	return password, nil
}
