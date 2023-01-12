# HTTP(s) bench / DDOS

## Like a application

    ./nameOfAppFile -url localhost:8080 -rps 100
- URL set endpoint
- RPS set count of requests in one second (0 to DDOS)

## Like a library

    func main() {

        // Need URL and RPS (count of requests in one second)
        test := bench.New(
            "localhost:8080",
            100)
    
        // Start test
        test.Start()

        // waiting some time
        time.Sleep(time.Second)

        // stop sending request
        test.Stop()
    }