package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var (
	url       string
	reqCount  uint64 = 0
	errCount  uint64 = 0
	startDate        = time.Now()
	isCreate         = true
)

func main() {
	for i := 0; i < len(os.Args); i++ {
		switch os.Args[i] {
		case "--url":
			existParam := strings.Contains(os.Args[i+1], "--")
			if i+1 < len(os.Args) && !existParam {
				url = os.Args[i+1]
				fmt.Println("  >  url:", url)
				i++
			} else {
				log.Fatal(" > url not set")
			}
		}
	}

	go func() {
		for {
			go request()
			if !isCreate {
				break
			}
		}
	}()

	func() {
		var cmd string
		for {
			fmt.Scan(&cmd)
			switch cmd {
			case "stop":
				isCreate = false
				end()
			}
		}
	}()
}

func request() {
	_, err := http.Get(url)
	if err != nil {
		errCount += 1
	}
	reqCount += 1
	return
}

func end() {
	workTime := time.Now().Sub(startDate)

	fmt.Print("\n\n")
	fmt.Println(" > works:", workTime.Seconds(), "sec.")
	fmt.Println("=<>==<>==<>==<>==<>==<>==<>==<>==<>==<>==<>==<>=")
	fmt.Println(" > requests:", reqCount)
	fmt.Println(" > request per second:", float32(reqCount)/float32(workTime.Seconds()))
	fmt.Println("=<>==<>==<>==<>==<>==<>==<>==<>==<>==<>==<>==<>=")
	fmt.Println(" > errors:", errCount)
	fmt.Println(" > error per second:", float32(errCount)/float32(workTime.Seconds()))
	fmt.Println("=<>==<>==<>==<>==<>==<>==<>==<>==<>==<>==<>==<>=")

	procentErrors := (float32(errCount) / float32(reqCount)) * 100
	fmt.Println(" > errors % =", procentErrors)
	os.Exit(0)
}
