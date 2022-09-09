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

// ScanPassword scans password from stdin
func ScanPassword() (string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("%vEnter password: %s", colour.GreenBold, colour.Normal)
	password, err := reader.ReadString('\n')
	checkErr(err)

	fmt.Printf("%vRepeat password: %s", colour.GreenBold, colour.Normal)
	passwordConfirm, err := reader.ReadString('\n')
	checkErr(err)

	if password != passwordConfirm {
		return "", mismatchedPassword{}
	}

	return password, nil
}
