package onetrades

import "github.com/Betarost/onetrades/futurebinance"

func NewFutureBinanceClient(apiKey, secretKey string) *futurebinance.Client {
	return futurebinance.NewClient(apiKey, secretKey)
}
