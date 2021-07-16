run:
	go run main.go

build:
	go build -o build/discord-clone-api main.go

dev:
	ls **/*.go | entr -r make run
