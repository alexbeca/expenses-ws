package main

import (
    "context"
    "flag"
    "fmt"
    "log"
    "net/http"
    "os"
    "os/signal"
    "syscall"

    "expenses-ws/services"
)

func main() {
    var (
        httpAddr = flag.String("http", ":8088", "http listen address")
    )
    flag.Parse()
    ctx := context.Background()
    // our app service
    srv := services.NewService()
    errChan := make(chan error)

    go func() {
        c := make(chan os.Signal, 1)
        signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
        errChan <- fmt.Errorf("%s", <-c)
    }()

    // mapping endpoints
    endpoints := services.Endpoints{
        GetEndpoint:      services.MakeGetEndpoint(srv),
        StatusEndpoint:   services.MakeStatusEndpoint(srv),
        ValidateEndpoint: services.MakeValidateEndpoint(srv),
    }

    // HTTP transport
    go func() {
        log.Println("app ws is listening on port:", *httpAddr)
        handler := services.NewHTTPServer(ctx, endpoints)
        errChan <- http.ListenAndServe(*httpAddr, handler)
    }()

    log.Fatalln(<-errChan)
}