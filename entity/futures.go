package entity

type FuturesBalance struct {
	Asset            string `json:"asset" bson:"asset"`
	Balance          string `json:"balance" bson:"balance"`
	Equity           string `json:"equity" bson:"equity"`
	Available        string `json:"available" bson:"available"`
	UnrealizedProfit string `json:"unrealizedProfit" bson:"unrealizedProfit"`
}

type Futures_PositionsMode struct {
	HedgeMode bool `json:"hedgeMode" bson:"hedgeMode"`
}

type Futures_MarginMode struct {
	MarginMode string `json:"marginMode" bson:"marginMode"`
}

type Futures_Leverage struct {
	Symbol   string `json:"symbol" bson:"symbol"`
	Leverage string `json:"leverage" bson:"leverage"`
	// MarginMode   string `json:"marginMode" bson:"marginMode"`
	// PositionSide string `json:"positionSide" bson:"positionSide"`
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
}

type Futures_PositionsHistory struct {
	Symbol              string `json:"symbol" bson:"symbol"`
	PositionId          string `json:"positionId" bson:"positionId"`
	PositionSide        string `json:"positionSide" bson:"positionSide"`
	PositionAmt         string `json:"positionAmt" bson:"positionAmt"`
	ExecutedPositionAmt string `json:"executedPositionAmt" bson:"executedPositionAmt"`
	AvgPrice            string `json:"avgPrice" bson:"avgPrice"`
	ExecutedAvgPrice    string `json:"executedAvgPrice" bson:"executedAvgPrice"`
	RealisedProfit      string `json:"realisedProfit" bson:"realisedProfit"`
	Fee                 string `json:"fee" bson:"fee"`
	Funding             string `json:"funding" bson:"funding"`
	MarginMode          string `json:"marginMode" bson:"marginMode"`
	CreateTime          int64  `json:"createTime" bson:"createTime"`
	UpdateTime          int64  `json:"updateTime" bson:"updateTime"`
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
