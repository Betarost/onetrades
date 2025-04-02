package entity

type ContractInfo struct {
	Symbol             string  `json:"symbol" bson:"symbol"`
	ContractSize       float64 `json:"contractSize" bson:"contractSize"`
	ContractMultiplier float64 `json:"contractMultiplier" bson:"contractMultiplier"`
	StepContractSize   float64 `json:"stepContractSize" bson:"stepContractSize"`
	MinContractSize    float64 `json:"minContractSize" bson:"minContractSize"`
	StepTickPrice      float64 `json:"stepTickPrice" bson:"stepTickPrice"`
	MaxLeverage        int     `json:"lever" bson:"lever"`
	State              string  `json:"state" bson:"state"`
	Type               string  `json:"type" bson:"type"`
}

type Kline struct {
	OpenPrice    float64 `json:"openPrice" bson:"openPrice"`
	HighestPrice float64 `json:"highestPrice" bson:"highestPrice"`
	LowestPrice  float64 `json:"lowestPrice" bson:"lowestPrice"`
	ClosePrice   float64 `json:"closePrice" bson:"closePrice"`
	Time         int64   `json:"time" bson:"time"`
	Complete     bool    `json:"complete" bson:"complete"`
}

type MarkPrice struct {
	Symbol string  `json:"symbol" bson:"symbol"`
	Price  float64 `json:"price" bson:"price"`
}

type Ticker struct {
	Symbol        string  `json:"symbol" bson:"symbol"`
	Open24hPrice  float64 `json:"open24hPrice" bson:"open24hPrice"`
	LastPrice     float64 `json:"lastPrice" bson:"lastPrice"`
	Volume24hCoin float64 `json:"volume24hCoin" bson:"volume24hCoin"`
	Volume24hUSDT float64 `json:"volume24hUSDT" bson:"volume24hUSDT"`
	Change24h     float64 `json:"change24h" bson:"change24h"`
	Time          int64   `json:"time" bson:"time"`
}
