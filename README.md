[![Build Status](https://travis-ci.org/emahiro/ae-plain-logger.svg?branch=master)](https://travis-ci.org/emahiro/ae-plain-logger)


# App Engine Plain Logger
This is Plain Logger for App Engine 2nd Generation.  
App Engine Plain Logger only supports structured logging (ref: https://cloud.google.com/logging/docs/structured-logging), but does not support logging experience which proprietary App Engine API provided.

# How to use

```go

mux := http.NewServeMux()
mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
    // some
})

server := http.Server {
    Addr: fmt.Sprintf(":%s", port)
    Handler: middleware.MwAEPlainLogger("label")(mux)
}

if err := server.ListenAndServe(); err != nil {
    log.Fatalf("shutdown server. err: %v", err)
}

```

# LICENSE
MIT
