# HTTP(s) bench

### **Go-bench** is a simple cross-platform HTTP(s) server benchmark application written in Go using the [fasthttp](https://github.com/valyala/fasthttp?ysclid=lcwgg8cpz3782494501) library.

## Application use

    ./file_name -url localhost:8080 -rps 100

- URL set endpoint
- RPS set count of requests in one second (max is 65535)

Result:

    Host: localhost:8080
    RPS: 100

    ^C       // Ctrl + C for stop

    {
        "req_count":282,
        "err_count":282,
        "average_response_time_ms":0,
        "max_response_time_ms":0,
        "min_response_time_ms":0,
        "time_of_bench_sec":3.085026,
        "percent_of_errors":100
    }
