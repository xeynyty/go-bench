# HTTP(s) bench

### **Go-bench** is a simple cross-platform HTTP(s) server benchmark application written in Go using the [fasthttp](https://github.com/valyala/fasthttp?ysclid=lcwgg8cpz3782494501) library.

## Application use

    ./file_name -url localhost:8080 -rps 100

- URL set endpoint
- RPS set count of requests in one second (max is 65535)