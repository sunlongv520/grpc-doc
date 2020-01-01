package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"gweb.jtthink.com/helper"
	"gweb.jtthink.com/services"
	"log"
)

func main(){
	//creds, err := credentials.NewClientTLSFromFile("keys/server.crt", "localhost")
	//if err != nil {
	//	log.Fatal(err)
	//}


	conn,err:=grpc.Dial(":8081",grpc.WithTransportCredentials(helper.GetClientCreds()))
	if err!=nil{
		log.Fatal(err)
	}
	defer conn.Close()

	prodClient:=services.NewProdServiceClient(conn)
	ctx:=context.Background()
	//prodRes,err:=prodClient.GetProdStock(context.Background(),
	//	&services.ProdRequest{ProdId:12})
	//if err!=nil{
	//	log.Fatal(err)
	//}
	//fmt.Println(prodRes.ProdStock)
	response,err:=prodClient.GetProdStocks(ctx,&services.QuerySize{Size:10})
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println(response.Prodres[2].ProdStock)


}
