[![Build Status](https://travis-ci.org/emahiro/aelog.svg?branch=master)](https://travis-ci.org/emahiro/ae-plain-logger)
[![Go Report Card](https://goreportcard.com/badge/github.com/emahiro/ae-plain-logger)](https://goreportcard.com/report/github.com/emahiro/ae-plain-logger)
[![GoDoc](https://godoc.org/github.com/emahiro/ae-plain-logger?status.svg)](https://godoc.org/github.com/emahiro/ae-plain-logger)


# App Engine Plain Logger
This is Plain Logger for App Engine 2nd Generation.  
App Engine Plain Logger only supports structured logging (ref: https://cloud.google.com/logging/docs/structured-logging), so this does not support logging experience which proprietary App Engine API provided.

# How to use

```go
mux := http.NewServeMux()
mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
    ctx := request.Context()
    aelog.Infof(ctx, "some log message")
})

h := middleware.AELogger("ServeHTTP")(mux)

if err := http.ListenAndServe(fmt.Sprintf(":%s", port), h); err != nil {
    panic(err)
}
```

# License
MIT
