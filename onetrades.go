package onetrades

import (
	"github.com/Betarost/onetrades/futurebinance"
	"github.com/Betarost/onetrades/futurebybit"
)

func NewFutureBinanceClient(apiKey, secretKey string) *futurebinance.Client {
	return futurebinance.NewClient(apiKey, secretKey)
}

func NewFutureBybitClient(apiKey, secretKey string) *futurebybit.Client {
	return futurebybit.NewClient(apiKey, secretKey)
}
