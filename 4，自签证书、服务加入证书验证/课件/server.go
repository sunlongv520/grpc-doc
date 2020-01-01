package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"gorpc.jtthink.com/services"
	"log"
	"net/http"
)

func main()  {

	creds, err := credentials.NewServerTLSFromFile("keys/server.crt", "keys/server_no_passwd.key")
	if err != nil {
		log.Fatal(err)
	}
	rpcServer:=grpc.NewServer(grpc.Creds(creds))
	services.RegisterProdServiceServer(rpcServer,new(services.ProdService))

	//lis,_:=net.Listen("tcp",":8081")
	//rpcServer.Serve(lis)

	mux:=http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hello"))
	})
	httpServer:=&http.Server{
		Addr:":8081",
		Handler:mux,
	}

	httpServer.ListenAndServeTLS("keys/server.crt","keys/server_no_passwd.key")




}
