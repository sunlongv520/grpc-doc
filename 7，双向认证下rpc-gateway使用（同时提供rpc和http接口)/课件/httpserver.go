package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"gorpc.jtthink.com/helper"
	"gorpc.jtthink.com/services"
	"log"
	"net/http"
)

func main()  {
	gwmux:=runtime.NewServeMux()
	opt:=[]grpc.DialOption{grpc.WithTransportCredentials(helper.GetClientCreds())}
	err:=services.RegisterProdServiceHandlerFromEndpoint(context.Background(),
		gwmux,"localhost:8081",opt)
	if err != nil {
		log.Fatal(err)
	}
	httpServer:=&http.Server{
		Addr:":8080",
		Handler:gwmux,
	}
	httpServer.ListenAndServe()

}