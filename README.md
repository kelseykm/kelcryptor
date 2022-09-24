![kelcryptor](./assets/banner.png)

# kelcryptor

A commandline file encryption and decryption application, written in [Go](https://go.dev).

## Features

- AES-256-GCM encryption
- No limit on file size to be encrypted
- Keeps track of the integrity of encrypted file
- Get alert if wrong password is used without having to wait for the entire file to be decrypted
- Pass multiple files at a go to be encrypted/decrypted
- Ignore and skip files with errors when encrypting/decrypting many files
- Keeps track of time taken to encrypt/decrypt file(s)
- Shows progress of encryption/decryption

## Installation

`go install -ldflags="-s -w" github.com/kelseykm/kelcryptor`

You can also clone this repository and run `make install` at the repository root

### Local build

Clone this repository and run `make` at the repository root
