package main

import (
	"fmt"

	"github.com/kelseykm/kelcryptor/colour"
	"github.com/kelseykm/kelcryptor/cryptography"
)

func main() {
	PrintBanner()

	toRecordTimeTaken, toEncrypt, toDecrypt, files := ParseFlags()
	fmt.Printf("time: %v, enc: %v, dec: %v, files: %v\n",
		toRecordTimeTaken, toEncrypt, toDecrypt, files)

	VerifyFiles(files)

	switch {
	case toEncrypt:
		for _, file := range files {
			cryptography.EncryptFile("password", file)
			fmt.Printf("%s[INFO]%s%s %s%s encrypted\n",
				colour.BlueBackground, colour.Normal, colour.WhiteBold, file, colour.Normal,
			)
		}
	case toDecrypt:
		for _, file := range files {
			err := cryptography.DecryptFile("password", file)
			if err != nil {
				panic(err)
			}
			fmt.Printf("%s[INFO]%s%s %s%s decrypted\n",
				colour.BlueBackground, colour.Normal, colour.WhiteBold, file, colour.Normal,
			)
		}
	}
}
