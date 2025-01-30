package onetrades

import "github.com/Betarost/onetrades/futuremexc"

type TradeName string

// Global enums
const (
	TradeNameMexc TradeName = "MEXC"
)

func NewFutureMexcClient(apiKey, secretKey string) *futuremexc.Client {
	return futuremexc.NewClient(apiKey, secretKey)
}
