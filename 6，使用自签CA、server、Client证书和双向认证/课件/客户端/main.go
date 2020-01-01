package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"gweb.jtthink.com/services"
	"io/ioutil"
	"log"
)

func main(){
	//creds, err := credentials.NewClientTLSFromFile("keys/server.crt", "localhost")
	//if err != nil {
	//	log.Fatal(err)
	//}

	cert,_:=tls.LoadX509KeyPair("cert/client.pem","cert/client.key")
	certPool := x509.NewCertPool()
	ca, _ := ioutil.ReadFile("cert/ca.pem")
	certPool.AppendCertsFromPEM(ca)

	creds:=credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},//客户端证书
		ServerName: "localhost",
		RootCAs:      certPool,
	})


	conn,err:=grpc.Dial(":8081",grpc.WithTransportCredentials(creds))
	if err!=nil{
		log.Fatal(err)
	}
	defer conn.Close()

	prodClient:=services.NewProdServiceClient(conn)
	prodRes,err:=prodClient.GetProdStock(context.Background(),
		&services.ProdRequest{ProdId:12})
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println(prodRes.ProdStock)
}
