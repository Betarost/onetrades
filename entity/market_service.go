package entity

type ContractInfo struct {
	Symbol       string  `json:"symbol" bson:"symbol"`
	ContractSize float64 `json:"contractSize" bson:"contractSize"`
	MaxLeverage  int     `json:"lever" bson:"lever"`
}

type Kline struct {
	OpenPrice    float64 `json:"openPrice" bson:"openPrice"`
	HighestPrice float64 `json:"highestPrice" bson:"highestPrice"`
	LowestPrice  float64 `json:"lowestPrice" bson:"lowestPrice"`
	ClosePrice   float64 `json:"closePrice" bson:"closePrice"`
	Time         int64   `json:"time" bson:"time"`
	Complete     bool    `json:"complete" bson:"complete"`
}
