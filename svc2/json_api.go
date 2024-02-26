package main

import (
	"context"
	// "time"
	"encoding/json"
	"math/rand"
	"net/http"
	"github.com/camzero94/svc2/types"
	// "github.com/camzero94/svc2/client"
	// "github.com/camzero94/svc2/proto"
	// "fmt"
)

// Same signature as the handler
type APIFunc func(context.Context, http.ResponseWriter, *http.Request) error

// JSON server that it is separated of the business logic aka service  and logging logic aka logging service, etc 
type JSONServer struct {
	svc FetchAvailableMoney
	addr string
}

func NewServer(svc FetchAvailableMoney,addr string) *JSONServer {
	return &JSONServer{
		addr: addr,
		svc: svc,
	}
}
func (s *JSONServer) Start() error {
	http.HandleFunc("/",makeHTTPHandlerFunc(s.handlerFetchBalance))
	return http.ListenAndServe(s.addr,nil)
}

// Create a HTTPAPIFunc
func makeHTTPHandlerFunc(apiFn APIFunc) http.HandlerFunc{
	ctx := context.Background()
	ctx = context.WithValue(ctx,"req_id",rand.Intn(100000))
	return func (w http.ResponseWriter, req *http.Request){
		if err:= apiFn(ctx, w, req); err != nil{
			writeJSON(w,http.StatusBadRequest,map[string]any{"error":err.Error()})
		}
	}
}

func (s *JSONServer) handlerFetchBalance (ctx context.Context,w http.ResponseWriter, r *http.Request) error{
	name := r.URL.Query().Get("name")
	balance, err := s.svc.FetchMoney(ctx,name)
	if err != nil {
		return err
	}

	
	// // Connect to Grpc svc1
	// grpcClient, err:= client.NewClientGrpc(":9000")	
	// if err != nil{
	// 	return err
	// }
	//
	// chErr := make(chan error) 
	// go func(chErr chan error){
	// 	time.Sleep(3 * time.Second)
	// 	resp , err := grpcClient.FetchMoney(ctx, &proto.BalanceRequest{Name:"camilo"})	
	// 	if err != nil{
	// 		chErr <- err
	// 	}
	// 	close(chErr)
	// 	fmt.Printf("%+v\n",resp)
	// }(chErr)
	//
	// go func() {
	// 	for {
	// 		// Attempt to read from the error channel
	// 		err, ok := <-chErr
	// 		// Check if the error channel is closed and empty
	// 		if !ok {
	// 			break
	// 		}
	// 		// Process the error (you can handle it as needed)
	// 		if err != nil {
	// 			fmt.Println("Received error:", err)
	// 		}
	// 	}
	// }()

	resp := &types.BalanceResponse{
		Name:name,
		Balance: balance,
	}
	return writeJSON(w,http.StatusOK,resp)
}

// Function that generalize the response of a JSON 
func writeJSON (w http.ResponseWriter,status int, v any) error{
	 w.WriteHeader(status)
	 return json.NewEncoder(w).Encode(v)
}
