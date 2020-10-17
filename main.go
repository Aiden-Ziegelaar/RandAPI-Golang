package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"log"
	"net"
	"randAPI/helpers"
	"randAPI/proto/grpcService"
	"time"
)

var (
	port 		= 10001
	certFilePath 	= "./x509/server-cert.pem"
	keyFilePath		= "./x509/server-key.pem"
	caCertFilePath	= "./x509/ca-cert.pem"
)

func seedRotator () {
	for {
		helpers.RotateSeed()
		time.Sleep(1000)
	}
}

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	// Load server's certificate and private key
	serverCert, err := tls.LoadX509KeyPair(certFilePath, keyFilePath)
	if err != nil {
		return nil, err
	}

	pemClientCA, err := ioutil.ReadFile(caCertFilePath)
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemClientCA) {
		return nil, fmt.Errorf("failed to add client CA's certificate")
	}

	// Create the credentials and return it
	config := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	}

	return credentials.NewTLS(config), nil
}


func main() {
	go seedRotator()

	var opts []grpc.ServerOption

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))

	if err != nil {
		log.Fatal("Could not instantiate listener")
	}

	tlsCredentials, err := loadTLSCredentials()

	if err != nil {
		log.Fatal("cannot load TLS credentials: ", err)
	}

	opts = []grpc.ServerOption{grpc.Creds(tlsCredentials)}

	grpcServer := grpc.NewServer(opts...)

	var service = server{serverName: "RandAPI"}

	grpcService.RegisterRandAPIServer(grpcServer, &service)

	grpcServer.Serve(lis)
}