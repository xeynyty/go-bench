.PHONY: build

build-windows:
	GOOS=windows GOARCH=amd64 go build -o ./build/ddos-win main.go

build-linux:
	GOOS=linux GOARCH=amd64 go build -o ./build/ddos-linux main.go

build-mac-arm:
	GOOS=darwin GOARCH=arm64 go build -o ./build/ddos-mac-arm main.go

build-mac-adm64:
	GOOS=darwin GOARCH=amd64 go build -o ./build/ddos-mac-amd64 main.go

run:
	go run main.go

.DEFAULT_GOAL = run