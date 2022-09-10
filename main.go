package main

import (
	"fmt"
	"os"
	"time"

	"github.com/kelseykm/kelcryptor/colour"
	"github.com/kelseykm/kelcryptor/cryptography"
)

func main() {
	retVal := 0
	defer func() {
		os.Exit(retVal)
	}()

	printBanner()

	toRecordTime, toEncrypt, toDecrypt, files := parseFlags()

	if err := verifyFiles(files); err != nil {
		fmt.Println(err.Error())
		retVal = 2
		return
	}

	password := func() string {
		for {
			password, err := scanPassword()
			if err == nil {
				return password
			}
			fmt.Println(err.Error())
		}
	}()

	switch {
	case toEncrypt:
		for _, file := range files {
			if !toRecordTime {
				if err := cryptography.EncryptFile(password, file); err != nil {
					fmt.Println(err.Error())
					retVal = 2
					return
				}
			} else {
				start := time.Now()

				if err := cryptography.EncryptFile(password, file); err != nil {
					fmt.Println(err.Error())
					retVal = 2
					return
				}

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
				if err := cryptography.DecryptFile(password, file); err != nil {
					fmt.Println(err.Error())
					retVal = 2
					return
				}
			} else {
				start := time.Now()

				if err := cryptography.DecryptFile(password, file); err != nil {
					fmt.Println(err.Error())
					retVal = 2
					return
				}

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
