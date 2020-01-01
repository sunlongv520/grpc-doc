package helper

import (
	"crypto/tls"
	"crypto/x509"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
)


//获取客户端证书配置
func  GetClientCreds() credentials.TransportCredentials   {
	cert,_:=tls.LoadX509KeyPair("cert/client.pem","cert/client.key")
	certPool := x509.NewCertPool()
	ca, _ := ioutil.ReadFile("cert/ca.pem")
	certPool.AppendCertsFromPEM(ca)

	creds:=credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},//客户端证书
		ServerName: "localhost",
		RootCAs:      certPool,
	})
	return creds
}