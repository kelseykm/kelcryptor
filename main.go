package main

import (
	"fmt"

	"github.com/kelseykm/kelcryptor/colour"
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
			fmt.Printf("%s%s[INFO]%s %s%s%s %sencrypted%s\n",
				colour.BlueBackground,
				colour.BlueBold,
				colour.Normal,
				colour.WhiteUnderlined,
				file,
				colour.Normal,
				colour.WhiteBold,
				colour.Normal,
			)
		}
	case toDecrypt:
		for _, file := range files {
			err := cryptography.DecryptFile(password, file)
			if err != nil {
				panic(err)
			}
			fmt.Printf("%s%s[INFO]%s %s%s%s decrypted\n",
				colour.BlueBackground,
				colour.BlueBold,
				colour.Normal,
				colour.WhiteUnderlined,
				file,
				colour.Normal,
				colour.WhiteBold,
				colour.Normal,
			)
		}
	}
}
