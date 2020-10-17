package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/keepalive"
	"io/ioutil"
	"log"
	"randAPI/helpers"
	"randAPI/proto/grpcService"
	"testing"
	"time"
)

var kacp = keepalive.ClientParameters{
	Time:                10 * time.Second, // send pings every 10 seconds if there is no activity
	Timeout:             time.Second,      // wait 1 second for ping back
	PermitWithoutStream: true,             // send pings even without active streams
}

func loadClientTLSCredentials() (credentials.TransportCredentials, error) {
	// Load certificate of the CA who signed server's certificate
	pemServerCA, err := ioutil.ReadFile("./x509/ca-cert.pem")
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemServerCA) {
		return nil, fmt.Errorf("failed to add server CA's certificate")
	}

	// Load client's certificate and private key
	clientCert, err := tls.LoadX509KeyPair("./x509/client-cert.pem", "./x509/client-key.pem")
	if err != nil {
		return nil, err
	}

	// Create the credentials and return it
	config := &tls.Config{
		Certificates: []tls.Certificate{clientCert},
		RootCAs:      certPool,
	}

	return credentials.NewTLS(config), nil
}

func connectMain() grpcService.RandAPIClient  {
	tlsCredentials, err := loadClientTLSCredentials()
	if err != nil {
		log.Fatal("cannot load TLS credentials: ", err)
	}

	cc1, err := grpc.Dial("0.0.0.0:10001", grpc.WithTransportCredentials(tlsCredentials),
		grpc.WithKeepaliveParams(kacp))
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}
	return grpcService.NewRandAPIClient(cc1)
}

func MainConnectionT (t *testing.T) {

	tlsCredentials, err := loadClientTLSCredentials()
	if err != nil {
		log.Fatal("cannot load TLS credentials: ", err)
	}

	cc1, err := grpc.Dial("localhost:10001", grpc.WithTransportCredentials(tlsCredentials))
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}
	if cc1.GetState() != connectivity.Idle {
		t.Errorf("Expected connection to be in state %s, got %s", connectivity.Idle, cc1.GetState())
	}
}

func MainReqBoolT (t *testing.T) {
	client := connectMain()

	req := empty.Empty{}

	var _, err = client.Boolean(context.Background(), &req)

	if err != nil {
		t.Errorf("Error calling Boolean: %s", err)
	}
}

func MainReqBoolLoadT (t *testing.T) {
	client := connectMain()

	req := empty.Empty{}

	var rounds = 10000
	var a, b int
	helpers.RotateSeed()
	for i := 0; i < rounds; i++ {
		var res, _ = client.Boolean(context.Background(), &req)
		if res.Result {
			a++
		} else {
			b++
		}
	}
	result := float64(a) / float64(a+b)
	if !(0.49 <  result && result < 0.51) {
		t.Errorf("Expected bias to be in range 0.49-0.51, got %f", result)
	}
}

func TestMainFunction(t *testing.T) {
	go main()
	t.Run("Connect", MainConnectionT)
	t.Run("RequestBoolean", MainReqBoolT)
	t.Run("RequestBooleanLoad", MainReqBoolLoadT)
}