package entity

type ContractInfo struct {
	Symbol       string  `json:"symbol" bson:"symbol"`
	ContractSize float64 `json:"contractSize" bson:"contractSize"`
}
