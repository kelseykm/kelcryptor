package errors

import (
	"fmt"

	"github.com/kelseykm/kelcryptor/colour"
)

// MismatchedPasswordError error
type MismatchedPasswordError struct{}

func (m MismatchedPasswordError) Error() string {
	return fmt.Sprintf("%s %s",
		colour.Error(),
		colour.Message("Passwords do not match"),
	)
}

// GenericError error
type GenericError struct{ Message string }

func (g GenericError) Error() string {
	return fmt.Sprintf("%s %s",
		colour.Error(),
		colour.Message(g.Message),
	)
}

// BadFilesError error
type BadFilesError struct{ Message string }

// AddToMesg takes a string and adds it to the error message
func (b *BadFilesError) AddToMesg(more string) {
	b.Message += more
}

func (b *BadFilesError) Error() string {
	return b.Message
}

// WrongPasswordError error
type WrongPasswordError struct{}

func (w WrongPasswordError) Error() string {
	return fmt.Sprintf("%s %s",
		colour.Error(),
		colour.Message("Incorrect password"),
	)
}

// FileModifiedError error
type FileModifiedError struct{}

func (f FileModifiedError) Error() string {
	return fmt.Sprintf("%s %s",
		colour.Error(),
		colour.Message("File interity compromised"),
	)
}
