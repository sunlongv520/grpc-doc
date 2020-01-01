package main

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/grpc"
	"gweb.jtthink.com/helper"
	"gweb.jtthink.com/services"
	"log"
	"time"
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

	//prodClient:=NewProdServiceClient(conn)
	 ctx:=context.Background()
	 t:=timestamp.Timestamp{Seconds:time.Now().Unix()}
	 orderClient:=services.NewOrderSerivceClient(conn)
	res,_:= orderClient.NewOrder(ctx,&services.OrderMain{
		 OrderId:1001,
		 OrderNo:"20190809",
		 OrderMoney:90,
		 OrderTime:&t,
	 })
	fmt.Println(res)


}
