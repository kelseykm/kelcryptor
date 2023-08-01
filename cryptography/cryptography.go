package cryptography

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"io"
	"os"
	"regexp"
	"sync"

	"github.com/kelseykm/kelcryptor/errors"
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

func getKey(password, salt []byte, size int) ([]byte, []byte, error) {
	if salt == nil {
		salt = make([]byte, size)
		rand.Read(salt)
	}

	key, err := scrypt.Key(password, salt, 1<<15, 8, 1, size)
	if err != nil {
		return nil, nil, errors.GenericError{Message: err.Error()}
	}
	return key, salt, nil
}

// EncryptFile encrypts the src using a key derived from the password
// and writes the encrypted file to dest
func EncryptFile(password []byte, src string) error {
	key, salt, err := getKey(password, nil, keySize)
	if err != nil {
		return errors.GenericError{Message: err.Error()}
	}

	passwordHash, passwordSalt, err := getKey(password, nil, passwordHashSize)
	if err != nil {
		return errors.GenericError{Message: err.Error()}
	}

	hmacHash := hmac.New(sha256.New, key)

	file, err := os.Open(src)
	if err != nil {
		return errors.GenericError{Message: err.Error()}
	}
	defer file.Close()

	dest := src + ".enc"
	destFile, err := os.Create(dest)
	if err != nil {
		return errors.GenericError{Message: err.Error()}
	}
	defer destFile.Close()

	var wg sync.WaitGroup
	defer wg.Wait()
	chunksEncryptedChannel := make(chan int, 100)

	fileInfo, err := file.Stat()
	if err != nil {
		return errors.GenericError{Message: err.Error()}
	}
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
	if err != nil {
		return errors.GenericError{Message: err.Error()}
	}

	_, err = destFile.Write(passwordHash)
	if err != nil {
		return errors.GenericError{Message: err.Error()}
	}

	_, err = destFile.Write(passwordSalt)
	if err != nil {
		return errors.GenericError{Message: err.Error()}
	}

	_, err = destFile.Write(salt)
	if err != nil {
		return errors.GenericError{Message: err.Error()}
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return errors.GenericError{Message: err.Error()}
	}

	buffer := make([]byte, chunkSize)
	for {
		nRead, err := file.Read(buffer)
		if err == io.EOF {
			break
		} else if err != nil {
			return errors.GenericError{Message: err.Error()}
		}

		hmacHash.Write(buffer[:nRead])

		aead, err := cipher.NewGCM(block)
		if err != nil {
			return errors.GenericError{Message: err.Error()}
		}

		nonce := make([]byte, nonceSize)
		rand.Read(nonce)

		_, err = destFile.Write(nonce)
		if err != nil {
			return errors.GenericError{Message: err.Error()}
		}

		encryptedBuffer := aead.Seal(nil, nonce, buffer[:nRead], nil)

		_, err = destFile.Write(encryptedBuffer)
		if err != nil {
			return errors.GenericError{Message: err.Error()}
		}

		chunksEncryptedChannel <- nRead
	}

	close(chunksEncryptedChannel)

	computedHmac := hmacHash.Sum(nil)
	destFile.Seek(0, 0)
	_, err = destFile.Write(computedHmac)
	if err != nil {
		return errors.GenericError{Message: err.Error()}
	}

	return nil
}

// DecryptFile decrypts the src with a key derived from
// the password and writes the decrypted file to the dest
func DecryptFile(password []byte, src string) error {
	file, err := os.Open(src)
	if err != nil {
		return errors.GenericError{Message: err.Error()}
	}
	defer file.Close()

	storedHmac := make([]byte, fileHmacSize)
	_, err = file.Read(storedHmac)
	if err != nil {
		return errors.GenericError{Message: err.Error()}
	}

	passwordHash := make([]byte, passwordHashSize)
	_, err = file.Read(passwordHash)
	if err != nil {
		return errors.GenericError{Message: err.Error()}
	}

	passwordSalt := make([]byte, passwordSaltSize)
	_, err = file.Read(passwordSalt)
	if err != nil {
		return errors.GenericError{Message: err.Error()}
	}

	compPasswordHash, _, err := getKey(password, passwordSalt, passwordSaltSize)
	if err != nil {
		return errors.GenericError{Message: err.Error()}
	}

	if !hmac.Equal(compPasswordHash, passwordHash) {
		return errors.WrongPasswordError{}
	}

	salt := make([]byte, saltSize)
	_, err = file.Read(salt)
	if err != nil {
		return errors.GenericError{Message: err.Error()}
	}

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
	if err != nil {
		return errors.GenericError{Message: err.Error()}
	}
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
	if err != nil {
		return errors.GenericError{Message: err.Error()}
	}
	defer destFile.Close()

	key, _, err := getKey(password, salt, keySize)
	if err != nil {
		return errors.GenericError{Message: err.Error()}
	}

	hmacHash := hmac.New(sha256.New, key)

	block, err := aes.NewCipher(key)
	if err != nil {
		return errors.GenericError{Message: err.Error()}
	}

	buffer := make([]byte, chunkSize+nonceSize+tagSize)
	for {
		nRead, err := file.Read(buffer)
		if err == io.EOF {
			break
		} else if err != nil {
			return errors.GenericError{Message: err.Error()}
		}

		aead, err := cipher.NewGCM(block)
		if err != nil {
			return errors.GenericError{Message: err.Error()}
		}

		nonce := buffer[:nonceSize]

		decryptedBuffer, err := aead.Open(nil, nonce, buffer[nonceSize:nRead], nil)
		if err != nil {
			return errors.FileModifiedError{}
		}

		hmacHash.Write(decryptedBuffer)

		_, err = destFile.Write(decryptedBuffer)
		if err != nil {
			return errors.GenericError{Message: err.Error()}
		}

		chunksDecryptedChannel <- nRead
	}

	close(chunksDecryptedChannel)

	computedHmac := hmacHash.Sum(nil)
	if !hmac.Equal(storedHmac, computedHmac) {
		return errors.FileModifiedError{}
	}

	return nil
}
