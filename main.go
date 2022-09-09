package main

import (
	"github.com/kelseykm/kelcryptor/cryptography"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	printBanner()

	_, toEncrypt, toDecrypt, files := parseFlags()

	verifyFiles(files)

	password, err := scanPassword()
	checkErr(err)

	switch {
	case toEncrypt:
		for _, file := range files {
			cryptography.EncryptFile(password, file)
		}
	case toDecrypt:
		for _, file := range files {
			err := cryptography.DecryptFile(password, file)
			checkErr(err)
		}
	}
}
