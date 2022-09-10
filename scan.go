package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/kelseykm/kelcryptor/colour"
)

type mismatchedPassword struct{}

func (m mismatchedPassword) Error() string {
	return fmt.Sprintf("%s %s",
		colour.Error(),
		colour.Message("Passwords do not match"),
	)
}

type genericError struct{ message string }

func (g genericError) Error() string {
	return fmt.Sprintf("%s %s",
		colour.Error(),
		colour.Message(g.message),
	)
}

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
		return "", genericError{err.Error()}
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
		return "", genericError{err.Error()}
	}

	if password != passwordConfirm {
		return "", mismatchedPassword{}
	}

	return password, nil
}
