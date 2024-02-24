package main

import (
	"fmt"
	"os"

	"github.com/kelseykm/kelcryptor/colour"
	"github.com/kelseykm/kelcryptor/errors"
)

func verifyFiles(files []string) ([]string, error) {
	var cleanFiles, nonExistentFiles, nonRegularFiles []string
	var badFilesErr *errors.BadFilesError

	for _, file := range files {
		fileInfo, err := os.Stat(file)
		if err != nil {
			nonExistentFiles = append(nonExistentFiles, file)
		} else if !fileInfo.Mode().IsRegular() {
			nonRegularFiles = append(nonRegularFiles, file)
		} else {
			cleanFiles = append(cleanFiles, file)
		}
	}

	if len(nonExistentFiles) != 0 {
		badFilesErr = &errors.BadFilesError{}

		badFilesErr.AddToMesg(
			fmt.Sprintf(
				"%s %s",
				colour.Error,
				colour.Message("Some files do not exist:\n"),
			))

		for index, file := range nonExistentFiles {
			badFilesErr.AddToMesg(
				fmt.Sprintf(
					"  %v: %s\n",
					index+1,
					colour.FileName(file),
				))
		}

	}

	if len(nonRegularFiles) != 0 {
		if badFilesErr == nil {
			badFilesErr = &errors.BadFilesError{}
		}

		badFilesErr.AddToMesg(
			fmt.Sprintf(
				"%s %s",
				colour.Error,
				colour.Message("Some files are not regular files:\n"),
			))

		for index, file := range nonRegularFiles {
			badFilesErr.AddToMesg(
				fmt.Sprintf(
					"  %v: %s\n",
					index+1,
					colour.FileName(file),
				))
		}
	}

	if badFilesErr == nil {
		return cleanFiles, nil
	}

	return cleanFiles, badFilesErr
}
