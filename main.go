package main

import (
	"fmt"
	"os"
	"time"

	"github.com/kelseykm/kelcryptor/colour"
	"github.com/kelseykm/kelcryptor/cryptography"
)

func main() {
	var retVal int
	defer func() {
		os.Exit(retVal)
	}()

	printBanner()

	toIgnore, toRecordTime, toEncrypt, toDecrypt, files := parseFlags()

	if cleanFiles, err := verifyFiles(files); err != nil {
		fmt.Println(err.Error())
		retVal = 2
		if !toIgnore {
			return
		}
		files = cleanFiles
	}

	password, err := func() (string, error) {
		var acceptableError mismatchedPassword
		for {
			password, err := scanPassword()
			if err == nil {
				return password, nil
			} else if err != acceptableError {
				return "", err
			}
			fmt.Println(err.Error())
		}
	}()
	if err != nil {
		fmt.Println(err.Error())
		retVal = 2
		return
	}

	switch {
	case toEncrypt:
		for _, file := range files {
			if !toRecordTime {
				if err := cryptography.EncryptFile(password, file); err != nil {
					fmt.Println(err.Error())
					retVal = 2
					if toIgnore {
						continue
					}
					return
				}
			} else {
				start := time.Now()

				if err := cryptography.EncryptFile(password, file); err != nil {
					fmt.Println(err.Error())
					retVal = 2
					if toIgnore {
						continue
					}
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
					if toIgnore {
						continue
					}
					return
				}
			} else {
				start := time.Now()

				if err := cryptography.DecryptFile(password, file); err != nil {
					fmt.Println(err.Error())
					retVal = 2
					if toIgnore {
						continue
					}
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
