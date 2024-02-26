package main

// The main idea of the logging service is to log to the client useful information
// about the service the client is using

import (
	"context"
	"log/slog"
	"os"
	"time"
)

type loggingService struct {
	next FetchAvailableMoney
}

func (l loggingService) FetchMoney(ctx context.Context, username string) (balance float64, err error) {

	// Before exiting the logger function calculates the time it took
	defer func(begin_t time.Time) {
		logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
		req_id := ctx.Value("req_id")
		logger.Info(
			"Incoming,req",
			slog.Int("Request_Id", req_id.(int)), 
			"fetchMoneyApi_took", time.Since(begin_t),
			"balance: ", balance,
			"err: ", err,
		)
	}(time.Now())

	// Business logic
	return l.next.FetchMoney(ctx, username)
}
