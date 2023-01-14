# HTTP(s) bench

## Application

    ./ddos -url localhost:8080 -rps 100

- URL set endpoint
- RPS set count of requests in one second (max is 255)

# How to compile application?

Clone repo

    git clone https://github.com/xeynyty/go-bench

Compile (GO >= 1.19 recommended)

    go build -o bench main.go