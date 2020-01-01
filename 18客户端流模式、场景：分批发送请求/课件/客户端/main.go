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

    ctx:=context.Background()
    userClient:=services.NewUserServiceClient(conn)
	var i int32
	if err!=nil{
		log.Fatal(err)
	}
	stream,err:=userClient.GetUserScoreByClientStream(ctx)
	if err!=nil{
		log.Fatal(err)
	}

		for j:=1;j<=3;j++{
			req:=services.UserScoreRequest{}
			req.Users=make([]*services.UserInfo,0)
			for i=1;i<=5;i++{  //加了5条用户信息  假设是一个耗时的过程
				req.Users=append(req.Users,&services.UserInfo{UserId:i})
			}
			err:=stream.Send(&req)
			if err!=nil{
				log.Println(err)
			}
		}
	res,_:=stream.CloseAndRecv()
	fmt.Println(res.Users)











}
