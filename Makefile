.PHONY: build

build-windows:
	GOOS=windows GOARCH=amd64 go build -o ./build/bench-win.exe -ldflags "-s -w" main.go

build-linux:
	GOOS=linux GOARCH=amd64 go build -o ./build/bench-linux -ldflags "-s -w" main.go

build-mac-arm:
	GOOS=darwin GOARCH=arm64 go build -o ./build/bench-mac-arm -ldflags "-s -w" main.go

build-mac-adm64:
	GOOS=darwin GOARCH=amd64 go build -o ./build/bench-mac-amd64 -ldflags "-s -w" main.go

run:
	go run main.go

.DEFAULT_GOAL = run