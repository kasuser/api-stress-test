# kasuser/stresstest
Simple backend API service and synthetic stress test.

## Requirements
* Go 1.14 or higher
* ab

## Usage

Run backend API service (add -race flag to diagnose data races)
```bash
$ go run main.go
```

## Manually testing

Run profiling tool
```bash
$ pprof -http=:6060 http://127.0.0.1:8080/debug/pprof/profile
```

Run one of the load utilities
```bash
$ go-wrk -c 80 -d 30 http://127.0.0.1:8080/request
```

or

```bash
$ ab -c 80 -n 1000000 http://127.0.0.1:8080/request
```

## Run tests
```bash
$ go test ./...
```