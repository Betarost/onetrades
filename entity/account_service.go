package entity

type AccountBalance struct {
	Asset              string  `json:"asset" bson:"asset"`
	Balance            float64 `json:"balance" bson:"balance"`
	AvailableBalance   float64 `json:"availableBalance" bson:"availableBalance"`
	CrossWalletBalance float64 `json:"crossWalletBalance" bson:"crossWalletBalance"`
	UnrealizedProfit   float64 `json:"unrealizedProfit" bson:"unrealizedProfit"`
}
