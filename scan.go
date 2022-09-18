package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/kelseykm/kelcryptor/colour"
	"github.com/kelseykm/kelcryptor/errors"
)

func scanPassword() (string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("%s %s",
		colour.Input(),
		colour.Message("Enter password: "),
	)
	fmt.Printf("%s", colour.Invisible)
	password, err := reader.ReadString('\n')
	fmt.Printf("%s", colour.Normal)
	if err != nil {
		fmt.Println()
		return "", errors.GenericError{err.Error()}
	}

	fmt.Printf("%s %s",
		colour.Input(),
		colour.Message("Repeat password: "),
	)
	fmt.Printf("%s", colour.Invisible)
	passwordConfirm, err := reader.ReadString('\n')
	fmt.Printf("%s", colour.Normal)
	if err != nil {
		fmt.Println()
		return "", errors.GenericError{err.Error()}
	}

	if password != passwordConfirm {
		return "", errors.MismatchedPassword{}
	}

	return password, nil
}
