package main

import (
	"context"
	"github.com/camzero94/svc2/proto"
	"google.golang.org/grpc"
	"net"
	"math/rand"
	// "fmt"
)

type GrpcApiHandler func (ctx context.Context, req *proto.BalanceRequest)

type GrpcServer struct{
	svc FetchAvailableMoney 
	proto.UnimplementedMoneyFetcherServer
}

func NewGrpcServer (svc FetchAvailableMoney) *GrpcServer{
	return &GrpcServer{
		svc: svc,
	}
}

func MakeGrcpServerandRun (svc FetchAvailableMoney, listAddr string) error{
 	grpcSrvrBalanceFetcher := NewGrpcServer(svc)
	ln , err := net.Listen("tcp",listAddr)
	if err != nil {
		return err
	}
	opts := []grpc.ServerOption{}
	server := grpc.NewServer(opts...)
	proto.RegisterMoneyFetcherServer(server,grpcSrvrBalanceFetcher)
	return server.Serve(ln)
}
// We need to implement  MoneyFetcherServer interface  from service_proto_grpc generated fie 
func (grpcSrv * GrpcServer) FetchMoney (ctx context.Context, req *proto.BalanceRequest) (*proto.BalanceResponse, error) {
	reqId := rand.Intn(1000)
	ctx = context.WithValue(ctx,"req_id",reqId)
	req.Name = "daniel"
	balance , err := grpcSrv.svc.FetchMoney(ctx,req.Name)
	if err != nil{
		return nil, err
	}
	
	resp := &proto.BalanceResponse{
		Name:req.Name,
		Balance: float32(balance),
	}

	return resp,err
}

