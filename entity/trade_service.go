package entity

type OrdersHistory struct {
	Symbol       string  `json:"symbol" bson:"symbol"`
	OrderID      string  `json:"orderId" bson:"orderId"`
	Side         string  `json:"side" bson:"side"`
	PositionSide string  `json:"positionSide" bson:"positionSide"`
	Price        float64 `json:"price" bson:"price"`
	OrigQty      float64 `json:"origQty" bson:"origQty"`
	AvgPrice     float64 `json:"avgPrice" bson:"avgPrice"`
	Type         string  `json:"type" bson:"type"`
	Status       string  `json:"status" bson:"status"`
	Time         int64   `json:"time" bson:"time"`
	UpdateTime   int64   `json:"updateTime" bson:"updateTime"`
	Hex          string  `json:"hex" bson:"hex"`
}

type OrderList struct {
	Symbol  string `json:"symbol" bson:"symbol"`
	OrderID string `json:"orderId" bson:"orderId"`
}
