PROJECT = $(shell basename $(CURDIR))
build:
	go build -o bin/${PROJECT} ./src
build-win:
	GOOS=windows GOARCH=amd64 go build -o bin/${PROJECT}.bat ./src