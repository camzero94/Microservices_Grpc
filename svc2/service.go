package main 

import (
	"context"
	"fmt"
)

type FetchAvailableMoney interface {
	FetchMoney(context.Context,string)(float64,error)
}

// Implementation of the FetchAbailableMoney interface
// This is where the business logic is implemented
type MoneyFetcher struct {}

func (s *MoneyFetcher) FetchMoney(ctx context.Context, username string)(money float64, err error)  {
	return  MockMoneyFetcher(ctx,username)
}

var mockBankAccount = map[string] float64{
	"camilo": 606,
	"daniel": 52000,
}
// All the business logic 
func MockMoneyFetcher (ctx context.Context , username string) (float64, error){
	balance, ok:= mockBankAccount[username] 
	if !ok{
		return 0, fmt.Errorf("Record for %s not found",username)
	}
	return balance , nil
}




