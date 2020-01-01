package main

import (
	"google.golang.org/grpc"
	"gorpc.jtthink.com/helper"
	"gorpc.jtthink.com/services"
	"net"
)

func main()  {

	rpcServer:=grpc.NewServer(grpc.Creds(helper.GetServerCreds()))
	services.RegisterProdServiceServer(rpcServer,new(services.ProdService))

	lis,_:=net.Listen("tcp",":8081")
	rpcServer.Serve(lis)

	//mux:=http.NewServeMux()
	//mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
	//	 rpcServer.ServeHTTP(writer,request)
	//})
	//httpServer:=&http.Server{
	//	Addr:":8081",
	//	Handler:mux,
	//}
	//httpServer.ListenAndServeTLS("keys/server.crt","keys/server.key")


}
