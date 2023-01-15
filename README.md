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
        "req_count":259,
        "err_count":0,
        "average_response_time_ms":57.96139,
        "max_response_time_ms":265,
        "min_response_time_ms":38,
        "time_of_bench_sec":2.7747707,
        "percent_of_errors":0
    }
