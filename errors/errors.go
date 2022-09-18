package errors

import (
	"fmt"

	"github.com/kelseykm/kelcryptor/colour"
)

// MismatchedPassword error
type MismatchedPassword struct{}

func (m MismatchedPassword) Error() string {
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

// BadFiles error
type BadFiles struct{ Message string }

// AddToMesg takes a string and adds it to the error message
func (b *BadFiles) AddToMesg(more string) {
	b.Message += more
}

func (b *BadFiles) Error() string {
	return b.Message
}

// WrongPassword error
type WrongPassword struct{}

func (w WrongPassword) Error() string {
	return fmt.Sprintf("%s %s",
		colour.Error(),
		colour.Message("Incorrect password"),
	)
}

// FileModified error
type FileModified struct{}

func (f FileModified) Error() string {
	return fmt.Sprintf("%s %s",
		colour.Error(),
		colour.Message("File interity compromised"),
	)
}
