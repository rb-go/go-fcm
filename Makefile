GOPACKAGE=github.com/riftbit/gofcm

default: prepare

prepare:
	go mod download
	go mod verify
	go mod tidy

generate: prepare
	ffjson -force-regenerate message.go
	ffjson -force-regenerate response.go

test: prepare
	go test -json ./...
	go test -short all
	go test -race $(GPACKAGE)
