package entity

type FuturesBalance struct {
	Asset   string `json:"asset" bson:"asset"`
	Balance string `json:"balance" bson:"balance"`
	Locked  string `json:"locked" bson:"locked"`
}

type Futures_InstrumentsInfo struct {
	Symbol         string `json:"symbol" bson:"symbol"`
	Base           string `json:"base" bson:"base"`
	Quote          string `json:"quote" bson:"quote"`
	MinQty         string `json:"minQty" bson:"minQty"`
	MinNotional    string `json:"minNotional" bson:"minNotional"`
	PricePrecision string `json:"pricePrecision" bson:"pricePrecision"`
	SizePrecision  string `json:"sizePrecision" bson:"sizePrecision"`
	State          string `json:"state" bson:"state"`
	MaxLeverage    string `json:"maxLeverage" bson:"maxLeverage"`
	Multiplier     string `json:"multiplier" bson:"multiplier"`
	ContractSize   string `json:"contractSize" bson:"contractSize"`
	IsSizeContract bool   `json:"isSizeContract" bson:"isSizeContract"`

	//==============
	// MaxLeverage        string `json:"maxLeverage" bson:"maxLeverage"`
	// StepTickPrice      string `json:"stepTickPrice" bson:"stepTickPrice"`
	// MinContractSize    string `json:"minContractSize" bson:"minContractSize"`
	// StepContractSize   string `json:"stepContractSize" bson:"stepContractSize"`
	// ContractSize       string `json:"contractSize" bson:"contractSize"`
	// ContractMultiplier string `json:"contractMultiplier" bson:"contractMultiplier"`
}

type Futures_Leverage struct {
	Symbol       string `json:"symbol" bson:"symbol"`
	Leverage     string `json:"leverage" bson:"leverage"`
	MarginMode   string `json:"marginMode" bson:"marginMode"`
	PositionSide string `json:"positionSide" bson:"positionSide"`
}

type Futures_Positions struct {
	Symbol           string `json:"symbol" bson:"symbol"`
	PositionSide     string `json:"positionSide" bson:"positionSide"`
	PositionAmt      string `json:"positionAmt" bson:"positionAmt"`
	PositionID       string `json:"positionID"`
	EntryPrice       string `json:"entryPrice" bson:"entryPrice"`
	MarkPrice        string `json:"markPrice" bson:"markPrice"`
	UnRealizedProfit string `json:"unRealizedProfit" bson:"unRealizedProfit"`
	RealizedProfit   string `json:"realizedProfit" bson:"realizedProfit"`
	Notional         string `json:"notional" bson:"notional"`
	InitialMargin    string `json:"initialMargin" bson:"initialMargin"`
	MarginRatio      string `json:"marginRatio" bson:"marginRatio"`
	AutoDeleveraging string `json:"autoDeleveraging" bson:"autoDeleveraging"`
	UpdateTime       int64  `json:"updateTime" bson:"updateTime"`
}

type Futures_OrdersList struct {
	Symbol        string `json:"symbol" bson:"symbol"`
	OrderID       string `json:"orderId" bson:"orderId"`
	ClientOrderID string `json:"clientOrderID" bson:"clientOrderID"`
	Side          string `json:"side" bson:"side"`
	PositionSide  string `json:"positionSide" bson:"positionSide"`
	PositionAmt   string `json:"positionAmt" bson:"positionAmt"`
	Price         string `json:"price" bson:"price"`
	TpPrice       string `json:"tpPrice" bson:"tpPrice"`
	SlPrice       string `json:"slPrice" bson:"slPrice"`
	Leverage      string `json:"leverage" bson:"leverage"`
	Type          string `json:"type" bson:"type"`
	TradeMode     string `json:"tradeMode" bson:"tradeMode"`
	InstType      string `json:"instType" bson:"instType"`
	Status        string `json:"status" bson:"status"`
	IsTpLimit     bool   `json:"isTpLimit" bson:"isTpLimit"`
	CreateTime    int64  `json:"createTime" bson:"createTime"`
	UpdateTime    int64  `json:"updateTime" bson:"updateTime"`
}
