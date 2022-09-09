package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/kelseykm/kelcryptor/colour"
)

type mismatchedPassword struct{}

func (m mismatchedPassword) Error() string {
	return "Passwords do not match"
}

func scanPassword() (string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("%vEnter password: %s", colour.GreenBold, colour.Normal)
	fmt.Printf("%v", colour.Invisible)
	password, err := reader.ReadString('\n')
	fmt.Printf("%v", colour.Normal)
	checkErr(err)

	fmt.Printf("%vRepeat password: %s", colour.GreenBold, colour.Normal)
	fmt.Printf("%v", colour.Invisible)
	passwordConfirm, err := reader.ReadString('\n')
	fmt.Printf("%v", colour.Normal)
	checkErr(err)

	if password != passwordConfirm {
		return "", mismatchedPassword{}
	}

	return password, nil
}
