package cryptography

import (
	"fmt"
	"sync"

	"github.com/kelseykm/kelcryptor/colour"
)

func printProgress(wg *sync.WaitGroup, fileName string, operation byte, totalSize int64, ch <-chan int) {
	var operationVerb string

	switch operation {
	case 'e':
		operationVerb = "encrypted"
	case 'd':
		operationVerb = "decrypted"
	}

	defer fmt.Printf("%s%s %s %s\n",
		colour.Overwrite,
		colour.Info(),
		colour.FileName(fileName),
		colour.Message(operationVerb),
	)

	defer wg.Done()

	chunks := 0

	for chunk := range ch {
		chunks += chunk

		percentage := float64(chunks) / float64(totalSize) * 100
		mesg := fmt.Sprintf("%06.2f%% %s:",
			percentage,
			operationVerb,
		)

		fmt.Printf("%s%s %s %s",
			colour.Overwrite,
			colour.Info(),
			colour.Message(mesg),
			colour.FileName(fileName),
		)
	}
}
