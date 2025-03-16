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
