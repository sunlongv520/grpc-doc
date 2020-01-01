package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"gweb.jtthink.com/helper"
	"gweb.jtthink.com/services"
	"io"
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

	 ctx:=context.Background()

    userClient:=services.NewUserServiceClient(conn)
	var i int32
		req:=services.UserScoreRequest{}
		req.Users=make([]*services.UserInfo,0)

		for i=1;i<6;i++{  //加了5条用户信息
			req.Users=append(req.Users,&services.UserInfo{UserId:i})
		}

		stream,err:=userClient.GetUserScoreByServerStream(ctx,&req)
		if err!=nil{
			log.Fatal(err)
		}
		for{
			res,err:=stream.Recv()
			if err==io.EOF{
				break
			}
			if err!=nil{
				log.Fatal(err)
			}
			fmt.Println(res.Users)

		}






}
