package mop

import (
	"net/http"
	"os"
	"os/signal"
	"context"
	"fmt"
	"syscall"
)

type gracefulSD struct {
	Sigint *chan os.Signal 
	Srv    *http.Server
}

func shutdown(gsd *gracefulSD) {

	fmt.Println("Waiting to shutdown")
	<-(*gsd.Sigint)
	fmt.Println("server is shutting down")

	//ctx,cancel:=context.WithTimeout(context.Background(),time.Second*5)
	//defer cancel()

	fmt.Println("Obtaining Context...")
	ctx := context.Background()
	fmt.Println("Shutting down...")
	err := gsd.Srv.Shutdown(ctx)
	if err!=nil{
		fmt.Println(err)
	} else {
		fmt.Println("Shutdown complete.")
	}
}


func setSignals() *chan os.Signal {
	
	// create a signal channel and feed it only from
	// interrupt signal sent from terminal or kubernetes
	// SIGTERM signal sent during rolling update
	sigint := make(chan os.Signal, 1)

	// interrupt signal sent from terminal
	signal.Notify(sigint, os.Interrupt)

	// sigterm signal sent from kubernetes
	signal.Notify(sigint, syscall.SIGTERM)

	fmt.Println("SIGTERM and SIGINT signals initialized!")
	return &sigint
}

// ex: SetServer(&http.Server{Addr: ":8080", Handler: handler})
func SetServer(srv *http.Server) {
	go func() {
		if err:=srv.ListenAndServe();err != http.ErrServerClosed{
			fmt.Println("ListenAndServe(): "+err.Error())
		}
	}()
	shutdown(&gracefulSD{Sigint: setSignals(), Srv: srv})
}
