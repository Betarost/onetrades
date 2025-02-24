package onetrades

import (
	"github.com/Betarost/onetrades/futurebinance"
	"github.com/Betarost/onetrades/futurebingx"
	"github.com/Betarost/onetrades/futurebybit"
	"github.com/Betarost/onetrades/futuremexc"
)

func NewFutureBinanceClient(apiKey, secretKey string) *futurebinance.Client {
	return futurebinance.NewClient(apiKey, secretKey)
}

func NewFutureBybitClient(apiKey, secretKey string) *futurebybit.Client {
	return futurebybit.NewClient(apiKey, secretKey)
}

func NewFutureMexcClient(apiKey, secretKey, memo string) *futuremexc.Client {
	return futuremexc.NewClient(apiKey, secretKey, memo)
}

func NewFutureBingxClient(apiKey, secretKey string) *futurebingx.Client {
	return futurebingx.NewClient(apiKey, secretKey)
}
