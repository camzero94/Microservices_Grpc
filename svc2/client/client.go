package client

import (
	"github.com/camzero94/svc2/proto"
	"google.golang.org/grpc"
)

func NewClientGrpc (listenAddr string) (proto.MoneyFetcherClient, error){
	conn, err := grpc.Dial(listenAddr,grpc.WithInsecure())
	if err 	!= nil{
		return nil, err
	}
	client := proto.NewMoneyFetcherClient(conn)
	return  client, nil
}
