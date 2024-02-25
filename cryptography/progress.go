package cryptography

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/kelseykm/kelcryptor/colour"
	"golang.org/x/term"
)

func printProgress(wg *sync.WaitGroup, fileName string, operation byte, totalSize int64, ch <-chan int) {
	var terminalWidth int

	if width, _, err := term.GetSize(int(os.Stdout.Fd())); err != nil {
		terminalWidth = 100
	} else {
		terminalWidth = width
	}

	var operationVerb string

	switch operation {
	case 'e':
		operationVerb = "encrypted"
	case 'd':
		operationVerb = "decrypted"
	}

	finalMesg := fmt.Sprintf("\r%s %s %s",
		colour.Info,
		colour.FileName(fileName),
		colour.Message(operationVerb),
	)

	var finalMesgExtraSpaces int
	finalMesgExtraSpaces = terminalWidth - len(finalMesg)
	if finalMesgExtraSpaces < 0 {
		finalMesgExtraSpaces = terminalWidth
	}

	defer fmt.Printf("%s%s\n", finalMesg, strings.Repeat(" ", finalMesgExtraSpaces))

	defer wg.Done()

	chunks := 0

	for chunk := range ch {
		chunks += chunk

		percentage := float64(chunks) / float64(totalSize) * 100
		mesg := fmt.Sprintf("%06.2f%% %s:",
			percentage,
			operationVerb,
		)

		fullMesg := fmt.Sprintf("\r%s %s %s",
			colour.Info,
			colour.Message(mesg),
			colour.FileName(fileName),
		)

		var extraSpaces int
		extraSpaces = terminalWidth - len(fullMesg)
		if extraSpaces < 0 {
			extraSpaces = terminalWidth
		}

		fmt.Printf("%s%s", fullMesg, strings.Repeat(" ", extraSpaces))
	}
}
