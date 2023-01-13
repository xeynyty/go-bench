package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/xeynyty/go-ddos/pkg/bench"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	url := flag.String(
		"url",
		"localhost:8080",
		"set endpoint like a localhost:8080 or https//google.com")
	reqPerSec := flag.Uint(
		"rps",
		100,
		"set a count of request per second\n"+
			"not more than 65535")
	flag.Parse()

	fmt.Printf(" Host: %v\n RPS: %v\n", *url, *reqPerSec)

	b := bench.New(
		*url,
		uint16(*reqPerSec))

	b.Start()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		syscall.SIGALRM)

	select {
	case <-sigChan:
		fmt.Print("\n")

		res := b.Stop()

		resByte, err := json.Marshal(res)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(resByte))

		os.Exit(0)
	}
}
