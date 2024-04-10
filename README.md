# mop
Mop allows a microservice to shutdown gracefully, waiting for pending requests to finish and closing idling connections before exiting. It uses the Shutdown method provided in go v1.8 and higher.

## usage
```go
import "github.com/crmathieu/mop"
```

Only one call is needed in the main:
```go
mop.SetServer(&http.Server{Addr: < :port >, Handler: < handler >})
```

where 
- < :port > is a string corresponding to the local port number (ie ":8080")
- < handler > is a server handler

Most services have in their main function a call to a server "ListenAndServe" function. This should be removed and replaced by a mop.SetServer call. SetServer is a wrapper function that will call ListenAndServe on the behalf of the service and, at the same time, set up the necessary system signals and channels. 

Adding Mop is easy. Here is an example of a service main function:

```go
func main() {
    ...
    var mux http.Server
    mux.HandleFunc("/", router)

    handler := cors.New(cors.Options{
        AllowedHeaders: []string{
            "Authorization",
            "Origin",
            "X-Requested-With",
            "Accept",
            "X-CSRF-Token",
            "Content-Type"},
        AllowCredentials: true,
    }).Handler(mux)
    http.ListenAndServe(":8022", handler)
}
```
to integrate mop, only the last line changes:

```go
func main() {
    ...
    var mux http.Server
    mux.HandleFunc("/", router)

    handler := cors.New(cors.Options{
        AllowedHeaders: []string{
            "Authorization",
            "Origin",
            "X-Requested-With",
            "Accept",
            "X-CSRF-Token",
            "Content-Type"},
        AllowCredentials: true,
    }).Handler(mux)
    // http.ListenAndServe(":8022", handler)
    mop.SetServer(&http.Server{Addr: ":8022", Handler: handler})
}
```
Any service using Mop will shutdown safely upon receiving 2 types of signals: SIGTERM or SIGINT (Ctrl+C). 
