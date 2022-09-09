package main

import (
	"fmt"
	"time"

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

	toRecordTime, toEncrypt, toDecrypt, files := parseFlags()

	verifyFiles(files)

	password, err := scanPassword()
	checkErr(err)

	switch {
	case toEncrypt:
		for _, file := range files {
			if !toRecordTime {
				cryptography.EncryptFile(password, file)
			} else {
				start := time.Now()

				cryptography.EncryptFile(password, file)

				timeTaken := time.Since(start).Seconds()

				mesg := fmt.Sprintf("Done in %.2f seconds",
					timeTaken,
				)

				fmt.Printf("%s %s\n",
					colour.Info(),
					colour.Message(mesg),
				)
			}
		}
	case toDecrypt:
		for _, file := range files {
			if !toRecordTime {
				err := cryptography.DecryptFile(password, file)
				checkErr(err)
			} else {
				start := time.Now()

				err := cryptography.DecryptFile(password, file)
				checkErr(err)

				timeTaken := time.Since(start).Seconds()

				mesg := fmt.Sprintf("Done in %.2f seconds",
					timeTaken,
				)

				fmt.Printf("%s %s\n",
					colour.Info(),
					colour.Message(mesg),
				)
			}
		}
	}
}
