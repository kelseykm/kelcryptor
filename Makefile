kelcryptor:
	go build -ldflags="-s -w"

.PHONY: install
install:
	go install -ldflags="-s -w"
