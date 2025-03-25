package entity

type SubAccountBalance struct {
	EquityBalance    float64 `json:"equityBalance" bson:"equityBalance"`
	UnrealizedProfit float64 `json:"unrealizedProfit" bson:"unrealizedProfit"`
}

type SubAccountFundingBalance struct {
	Asset            string  `json:"asset" bson:"asset"`
	Balance          float64 `json:"balance" bson:"balance"`
	FrozenBalance    float64 `json:"frozenBalance" bson:"frozenBalance"`
	AvailableBalance float64 `json:"availableBalance" bson:"availableBalance"`
}
