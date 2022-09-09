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

	fmt.Printf("%s%s[INPUT]%s %sEnter password: %s",
		colour.GreenBackground,
		colour.GreenBold,
		colour.Normal,
		colour.GreenBold,
		colour.Normal,
	)
	fmt.Printf("%s", colour.Invisible)
	password, err := reader.ReadString('\n')
	fmt.Printf("%s", colour.Normal)
	checkErr(err)

	fmt.Printf("%s%s[INPUT]%s %sRepeat password: %s",
		colour.GreenBackground,
		colour.GreenBold,
		colour.Normal,
		colour.GreenBold,
		colour.Normal,
	)
	fmt.Printf("%s", colour.Invisible)
	passwordConfirm, err := reader.ReadString('\n')
	fmt.Printf("%s", colour.Normal)
	checkErr(err)

	if password != passwordConfirm {
		return "", mismatchedPassword{}
	}

	return password, nil
}
