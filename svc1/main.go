// Service 2 is in charge of comunicate with service 1  using gRCP
package main

import (
	// "context"
	"flag"
	"fmt"
	"log"
	// "time"
	// "github.com/camzero94/svc1/client"
	// "github.com/camzero94/svc1/proto"
)

func main() {
	// Start server with port and service
	// Create my service with logging 
	var (
		jsonPort = flag.String("json", ":7000", "Specify the port to start the JSON server")
		grpcPort = flag.String("grpc",":6000", "Specify the port to start the Grpc server")
		svc = loggingService{&MoneyFetcher{}}
		// ctx = context.Background()
	)
	flag.Parse()

	// // ---------------------- Testing only Grpc Client  ------------------------ 
	// grpcClient, err:= client.NewClientGrpc(":6000")	
	// if err != nil{
	// 	log.Fatal(err)
	// }
	// go func(){
	// 	for {
	// 		time.Sleep(3 * time.Second)
	// 		resp , err := grpcClient.FetchMoney(ctx, &proto.BalanceRequest{Name:"camilo"})	
	// 		if err != nil{
	// 			log.Fatal(err)
	// 		}
	// 		fmt.Printf("%+v\n",resp)
	// 	}
	// }()

	// Create the Grpc server and pass my service to it 
	// (Because the implememntation of the Grpc service ) use gorutine
	go MakeGrcpServerandRun(svc,*grpcPort)

	// Create the JSON server and pass my service to it
	jsonServer := NewServer(svc,*jsonPort)
	fmt.Printf("Server started at port %s\n",*jsonPort )
	log.Fatal(jsonServer.Start())
}
