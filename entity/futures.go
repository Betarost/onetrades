package entity

type Futures_InstrumentsInfo struct {
	Symbol           string `json:"symbol" bson:"symbol"`
	Base             string `json:"base" bson:"base"`
	Quote            string `json:"quote" bson:"quote"`
	InstType         string `json:"instType" bson:"instType"`
	State            string `json:"state" bson:"state"`
	MaxLeverage      string `json:"lever" bson:"lever"`
	StepTickPrice    string `json:"stepTickPrice" bson:"stepTickPrice"`
	MinContractSize  string `json:"minContractSize" bson:"minContractSize"`
	StepContractSize string `json:"stepContractSize" bson:"stepContractSize"`

	ContractSize       string `json:"contractSize" bson:"contractSize"`
	ContractMultiplier string `json:"contractMultiplier" bson:"contractMultiplier"`
}
