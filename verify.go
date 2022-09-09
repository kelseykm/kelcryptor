package main

import (
	"fmt"
	"os"

	"github.com/kelseykm/kelcryptor/colour"
)

type badFiles struct {
	message string
}

func (b *badFiles) addToMesg(more string) {
	b.message += more
}

func (b *badFiles) Error() string {
	return b.message
}

func verifyFiles(files []string) error {
	var nonExistentFiles, nonRegularFiles []string
	var badFilesErr *badFiles

	for _, file := range files {
		fileInfo, err := os.Stat(file)
		if err != nil {
			nonExistentFiles = append(nonExistentFiles, file)
		} else if !fileInfo.Mode().IsRegular() {
			nonRegularFiles = append(nonRegularFiles, file)
		}
	}

	if len(nonExistentFiles) != 0 {
		badFilesErr = &badFiles{}

		badFilesErr.addToMesg(
			fmt.Sprintf(
				"%s %s",
				colour.Error(),
				colour.Message("Some files do not exist:\n"),
			))

		for index, file := range nonExistentFiles {
			badFilesErr.addToMesg(
				fmt.Sprintf(
					"  %v: %s\n",
					index+1,
					colour.FileName(file),
				))
		}

	}

	if len(nonRegularFiles) != 0 {
		if badFilesErr == nil {
			badFilesErr = &badFiles{}
		}

		badFilesErr.addToMesg(
			fmt.Sprintf(
				"%s %s",
				colour.Error(),
				colour.Message("Some files are not regular files:\n"),
			))

		for index, file := range nonRegularFiles {
			badFilesErr.addToMesg(
				fmt.Sprintf(
					"  %v: %s\n",
					index+1,
					colour.FileName(file),
				))
		}
	}

	if badFilesErr == nil {
		return nil
	}
	return badFilesErr
}
