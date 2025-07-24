package entity

// New FUTURES

type FuturesBalance struct {
	Asset            string `json:"asset" bson:"asset"`
	Balance          string `json:"balance" bson:"balance"`
	Equity           string `json:"equity" bson:"equity"`
	Available        string `json:"available" bson:"available"`
	UnrealizedProfit string `json:"unrealizedProfit" bson:"unrealizedProfit"`
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

type Futures_PositionsMode struct {
	HedgeMode bool `json:"hedgeMode" bson:"hedgeMode"`
}

type Futures_Leverage struct {
	Symbol        string `json:"symbol" bson:"symbol"`
	Leverage      string `json:"leverage" bson:"leverage"`
	LongLeverage  string `json:"longLeverage" bson:"longLeverage"`
	ShortLeverage string `json:"shortLeverage" bson:"shortLeverage"`
	// MarginMode   string `json:"marginMode" bson:"marginMode"`
	// PositionSide string `json:"positionSide" bson:"positionSide"`
}

type Futures_MarginMode struct {
	MarginMode string `json:"marginMode" bson:"marginMode"`
}

type Futures_ListenKey struct {
	ListenKey string `json:"listenKey" bson:"listenKey"`
}

// OLD

type Futures_PositionsHistory struct {
	Symbol              string `json:"symbol" bson:"symbol"`
	PositionID          string `json:"positionID" bson:"positionID"`
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

type Futures_OrdersHistory struct {
	Symbol         string `json:"symbol" bson:"symbol"`
	OrderID        string `json:"orderId" bson:"orderId"`
	ClientOrderID  string `json:"clientOrderID" bson:"clientOrderID"`
	PositionID     string `json:"positionID" bson:"positionID"`
	Side           string `json:"side" bson:"side"`
	PositionSide   string `json:"positionSide" bson:"positionSide"`
	PositionSize   string `json:"positionSize" bson:"positionSize"`
	ExecutedSize   string `json:"executedSize" bson:"executedSize"`
	Price          string `json:"price" bson:"price"`
	ExecutedPrice  string `json:"executedPrice" bson:"executedPrice"`
	RealisedProfit string `json:"realisedProfit" bson:"realisedProfit"`
	Fee            string `json:"fee" bson:"fee"`
	Leverage       string `json:"leverage"  bson:"leverage"`
	Type           string `json:"type" bson:"type"`
	Status         string `json:"status" bson:"status"`
	HedgeMode      bool   `json:"hedgeMode" bson:"hedgeMode"`
	MarginMode     string `json:"marginMode" bson:"marginMode"`
	// Cursor         string `json:"cursor,omitempty" bson:"cursor,omitempty"`
	CreateTime int64 `json:"createTime" bson:"createTime"`
	UpdateTime int64 `json:"updateTime" bson:"updateTime"`
}

type Futures_Positions struct {
	Symbol           string `json:"symbol" bson:"symbol"`
	PositionSide     string `json:"positionSide" bson:"positionSide"`
	PositionSize     string `json:"positionSize"`
	Leverage         string `json:"leverage"`
	PositionID       string `json:"positionID"`
	EntryPrice       string `json:"entryPrice" bson:"entryPrice"`
	MarkPrice        string `json:"markPrice" bson:"markPrice"`
	UnRealizedProfit string `json:"unRealizedProfit" bson:"unRealizedProfit"`
	RealizedProfit   string `json:"realizedProfit" bson:"realizedProfit"`
	Notional         string `json:"notional" bson:"notional"`
	HedgeMode        bool   `json:"hedgeMode" bson:"hedgeMode"`
	MarginMode       string `json:"marginMode" bson:"marginMode"`
	CreateTime       int64  `json:"createTime" bson:"createTime"`
	UpdateTime       int64  `json:"updateTime" bson:"updateTime"`

	//
	// InitialMargin    string `json:"initialMargin" bson:"initialMargin"`
	// MarginRatio      string `json:"marginRatio" bson:"marginRatio"`
	// AutoDeleveraging string `json:"autoDeleveraging" bson:"autoDeleveraging"`
	// PositionAmt      string `json:"positionAmt" bson:"positionAmt"`
}

type Futures_OrdersList struct {
	Symbol        string `json:"symbol" bson:"symbol"`
	OrderID       string `json:"orderId" bson:"orderId"`
	ClientOrderID string `json:"clientOrderID" bson:"clientOrderID"`
	PositionID    string `json:"positionID" bson:"positionID"`
	Side          string `json:"side" bson:"side"`
	PositionSide  string `json:"positionSide" bson:"positionSide"`
	PositionSize  string `json:"positionSize" bson:"positionSize"`
	ExecutedSize  string `json:"executedSize" bson:"executedSize"`
	Price         string `json:"price" bson:"price"`
	Leverage      string `json:"leverage" bson:"leverage"`
	Type          string `json:"type" bson:"type"`
	Status        string `json:"status" bson:"status"`
	CreateTime    int64  `json:"createTime" bson:"createTime"`
	UpdateTime    int64  `json:"updateTime" bson:"updateTime"`

	MarginMode string `json:"marginMode" bson:"marginMode"`

	//==========
	// TpPrice     string `json:"tpPrice" bson:"tpPrice"`
	// SlPrice     string `json:"slPrice" bson:"slPrice"`
	// TradeMode   string `json:"tradeMode" bson:"tradeMode"`
	// InstType    string `json:"instType" bson:"instType"`
	// IsTpLimit   bool   `json:"isTpLimit" bson:"isTpLimit"`
	// PositionAmt string `json:"positionAmt" bson:"positionAmt"`
}
