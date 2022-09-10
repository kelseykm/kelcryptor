package cryptography

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"regexp"
	"sync"

	"github.com/kelseykm/kelcryptor/colour"
	"golang.org/x/crypto/scrypt"
)

const (
	keySize = 32
	saltSize
	fileHmacSize
	nonceSize = 12
	tagSize   = 16
	passwordHashSize
	passwordSaltSize
	chunkSize = 16 * 1024
)

func getKey(password, salt []byte, size int) ([]byte, []byte) {
	if salt == nil {
		salt = make([]byte, size)
		rand.Read(salt)
	}

	key, err := scrypt.Key(password, salt, 1<<15, 8, 1, size)
	checkErr(err)
	return key, salt
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

// EncryptFile encrypts the src using a key derived from the password
// and writes the encrypted file to dest
func EncryptFile(password, src string) {
	key, salt := getKey([]byte(password), nil, keySize)
	passwordHash, passwordSalt := getKey([]byte(password), nil, passwordHashSize)

	hmacHash := hmac.New(sha256.New, key)

	file, err := os.Open(src)
	checkErr(err)
	defer func() {
		err := file.Close()
		checkErr(err)
	}()

	dest := src + ".enc"
	destFile, err := os.Create(dest)
	checkErr(err)
	defer func() {
		err := destFile.Close()
		checkErr(err)
	}()

	var wg sync.WaitGroup
	defer wg.Wait()
	chunksEncryptedChannel := make(chan int, 100)

	fileInfo, err := file.Stat()
	checkErr(err)
	fileSize := fileInfo.Size()

	wg.Add(1)
	go printProgress(
		&wg,
		src,
		'e',
		fileSize,
		chunksEncryptedChannel,
	)

	// make space for file hmac
	_, err = destFile.Write(make([]byte, fileHmacSize))
	checkErr(err)

	_, err = destFile.Write(passwordHash)
	checkErr(err)

	_, err = destFile.Write(passwordSalt)
	checkErr(err)

	_, err = destFile.Write(salt)
	checkErr(err)

	block, err := aes.NewCipher(key)
	checkErr(err)

	buffer := make([]byte, chunkSize)
	for {
		nRead, err := file.Read(buffer)
		if err == io.EOF {
			break
		} else {
			checkErr(err)
		}

		hmacHash.Write(buffer[:nRead])

		aead, err := cipher.NewGCM(block)
		checkErr(err)

		nonce := make([]byte, nonceSize)
		rand.Read(nonce)

		_, err = destFile.Write(nonce)
		checkErr(err)

		encryptedBuffer := aead.Seal(nil, nonce, buffer[:nRead], nil)

		_, err = destFile.Write(encryptedBuffer)
		checkErr(err)

		chunksEncryptedChannel <- nRead
	}

	close(chunksEncryptedChannel)

	computedHmac := hmacHash.Sum(nil)
	destFile.Seek(0, 0)
	_, err = destFile.Write(computedHmac)
	checkErr(err)
}

type wrongPassword struct{}

func (w wrongPassword) Error() string {
	return fmt.Sprintf("%s %s",
		colour.Error(),
		colour.Message("Incorrect password"),
	)
}

type fileModified struct{}

func (f fileModified) Error() string {
	return fmt.Sprintf("%s %s",
		colour.Error(),
		colour.Message("File interity compromised"),
	)
}

// DecryptFile decrypts the src with a key derived from
// the password and writes the decrypted file to the dest
func DecryptFile(password, src string) error {
	file, err := os.Open(src)
	checkErr(err)
	defer func() {
		err := file.Close()
		checkErr(err)
	}()

	storedHmac := make([]byte, fileHmacSize)
	_, err = file.Read(storedHmac)
	checkErr(err)

	passwordHash := make([]byte, passwordHashSize)
	_, err = file.Read(passwordHash)
	checkErr(err)

	passwordSalt := make([]byte, passwordSaltSize)
	_, err = file.Read(passwordSalt)
	checkErr(err)

	compPasswordHash, _ := getKey([]byte(password), passwordSalt, passwordSaltSize)
	if eq := hmac.Equal(compPasswordHash, passwordHash); !eq {
		return wrongPassword{}
	}

	salt := make([]byte, saltSize)
	_, err = file.Read(salt)
	checkErr(err)

	var dest string
	if re, _ := regexp.Compile("enc$"); re.Match([]byte(src)) {
		dest = string(re.ReplaceAll([]byte(src), []byte("dec")))
	} else {
		dest = src + ".dec"
	}

	var wg sync.WaitGroup
	defer wg.Wait()
	chunksDecryptedChannel := make(chan int, 100)

	fileInfo, err := file.Stat()
	checkErr(err)
	fileSize := fileInfo.Size()

	wg.Add(1)
	go printProgress(
		&wg,
		src,
		'd',
		fileSize,
		chunksDecryptedChannel,
	)

	destFile, err := os.Create(dest)
	checkErr(err)
	defer func() {
		err := destFile.Close()
		checkErr(err)
	}()

	key, _ := getKey([]byte(password), salt, keySize)

	hmacHash := hmac.New(sha256.New, key)

	block, err := aes.NewCipher(key)
	checkErr(err)

	buffer := make([]byte, chunkSize+nonceSize+tagSize)
	for {
		nRead, err := file.Read(buffer)
		if err == io.EOF {
			break
		} else {
			checkErr(err)
		}

		aead, err := cipher.NewGCM(block)
		checkErr(err)

		nonce := buffer[:nonceSize]

		decryptedBuffer, err := aead.Open(nil, nonce, buffer[nonceSize:nRead], nil)
		checkErr(err)

		hmacHash.Write(decryptedBuffer)

		_, err = destFile.Write(decryptedBuffer)
		checkErr(err)

		chunksDecryptedChannel <- nRead
	}

	close(chunksDecryptedChannel)

	computedHmac := hmacHash.Sum(nil)
	if fileAuthentic := hmac.Equal(storedHmac, computedHmac); !fileAuthentic {
		return fileModified{}
	}

	return nil
}
