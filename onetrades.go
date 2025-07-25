package onetrades

import (
	"github.com/Betarost/onetrades/futurebinance"
	"github.com/Betarost/onetrades/futurebingx"
	"github.com/Betarost/onetrades/futurebitget"
	"github.com/Betarost/onetrades/futurebybit"
	"github.com/Betarost/onetrades/futuregate"
	"github.com/Betarost/onetrades/futuremexc"
	"github.com/Betarost/onetrades/futureokx"
	"github.com/Betarost/onetrades/optionokx"
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

func NewFutureGateClient(apiKey, secretKey string) *futuregate.Client {
	return futuregate.NewClient(apiKey, secretKey)
}

func NewFutureBitgetClient(apiKey, secretKey, memo string) *futurebitget.Client {
	return futurebitget.NewClient(apiKey, secretKey, memo)
}

func NewFutureOKXClient(apiKey, secretKey, memo string) *futureokx.Client {
	return futureokx.NewClient(apiKey, secretKey, memo)
}

func NewOptionOKXClient(apiKey, secretKey, memo string) *optionokx.Client {
	return optionokx.NewClient(apiKey, secretKey, memo)
}
