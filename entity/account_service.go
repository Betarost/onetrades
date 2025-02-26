package entity

type AccountBalance struct {
	Asset              string  `json:"asset" bson:"asset"`
	Balance            float64 `json:"balance" bson:"balance"`
	AvailableBalance   float64 `json:"availableBalance" bson:"availableBalance"`
	CrossWalletBalance float64 `json:"crossWalletBalance" bson:"crossWalletBalance"`
	UnrealizedProfit   float64 `json:"unrealizedProfit" bson:"unrealizedProfit"`
}

type Position struct {
	Symbol           string  `json:"symbol" bson:"symbol"`
	PositionSide     string  `json:"positionSide" bson:"positionSide"`
	PositionAmt      float64 `json:"positionAmt" bson:"positionAmt"`
	EntryPrice       float64 `json:"entryPrice" bson:"entryPrice"`
	MarkPrice        float64 `json:"markPrice" bson:"markPrice"`
	UnRealizedProfit float64 `json:"unRealizedProfit" bson:"unRealizedProfit"`
	RealizedProfit   float64 `json:"realizedProfit" bson:"realizedProfit"`
	Notional         float64 `json:"notional" bson:"notional"`
	InitialMargin    float64 `json:"initialMargin" bson:"initialMargin"`
	UpdateTime       int64   `json:"updateTime" bson:"updateTime"`
}
