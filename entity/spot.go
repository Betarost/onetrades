package entity

// New SPOT

type AssetsBalance struct {
	Asset   string `json:"asset" bson:"asset"`
	Balance string `json:"balance" bson:"balance"`
	Locked  string `json:"locked" bson:"locked"`
}

type Spot_InstrumentsInfo struct {
	Symbol         string `json:"symbol" bson:"symbol"`
	Base           string `json:"base" bson:"base"`
	Quote          string `json:"quote" bson:"quote"`
	MinQty         string `json:"minQty" bson:"minQty"`
	MinNotional    string `json:"minNotional" bson:"minNotional"`
	PricePrecision string `json:"pricePrecision" bson:"pricePrecision"`
	SizePrecision  string `json:"sizePrecision" bson:"sizePrecision"`
	State          string `json:"state" bson:"state"`
}

type PlaceOrder struct {
	OrderID       string `json:"orderId" bson:"orderId"`
	ClientOrderID string `json:"clientOrderID" bson:"clientOrderID"`
	PositionID    string `json:"positionID,omitempty" bson:"positionID,omitempty"`
	Ts            int64  `json:"ts" bson:"ts"`
}

type Spot_OrdersList struct {
	Symbol        string `json:"symbol" bson:"symbol"`
	OrderID       string `json:"orderId" bson:"orderId"`
	ClientOrderID string `json:"clientOrderID" bson:"clientOrderID"`
	Side          string `json:"side" bson:"side"`
	Size          string `json:"size" bson:"size"`
	ExecutedSize  string `json:"executedSize" bson:"executedSize"`
	Price         string `json:"price" bson:"price"`
	Type          string `json:"type" bson:"type"`
	Status        string `json:"status" bson:"status"`
	CreateTime    int64  `json:"createTime" bson:"createTime"`
	UpdateTime    int64  `json:"updateTime" bson:"updateTime"`
}

type Spot_OrdersHistory struct {
	Symbol        string `json:"symbol" bson:"symbol"`
	OrderID       string `json:"orderId" bson:"orderId"`
	ClientOrderID string `json:"clientOrderID" bson:"clientOrderID"`
	Side          string `json:"side" bson:"side"`
	Size          string `json:"size" bson:"size"`
	ExecutedSize  string `json:"executedSize" bson:"executedSize"`
	Price         string `json:"price" bson:"price"`
	ExecutedPrice string `json:"executedPrice" bson:"executedPrice"`
	Fee           string `json:"fee" bson:"fee"`
	Type          string `json:"type" bson:"type"`
	Status        string `json:"status" bson:"status"`
	Cursor        string `json:"cursor,omitempty" bson:"cursor,omitempty"`
	CreateTime    int64  `json:"createTime" bson:"createTime"`
	UpdateTime    int64  `json:"updateTime" bson:"updateTime"`
}

// OLD

type TradingAccountBalance struct {
	TotalEquity      string                         `json:"totalEquity" bson:"totalEquity"`
	AvailableEquity  string                         `json:"availableEquity" bson:"availableEquity"`
	NotionalUsd      string                         `json:"notionalUsd" bson:"notionalUsd"`
	UnrealizedProfit string                         `json:"unrealizedProfit" bson:"unrealizedProfit"`
	UpdateTime       int64                          `json:"updateTime" bson:"updateTime"`
	Assets           []TradingAccountBalanceDetails `json:"assets" bson:"assets"`
}

type TradingAccountBalanceDetails struct {
	Asset            string `json:"asset" bson:"asset"`
	Balance          string `json:"balance" bson:"balance"`
	EquityBalance    string `json:"equityBalance" bson:"equityBalance"`
	AvailableBalance string `json:"availableBalance" bson:"availableBalance"`
	AvailableEquity  string `json:"availableEquity" bson:"availableEquity"`
	UnrealizedProfit string `json:"unrealizedProfit" bson:"unrealizedProfit"`
}

type FundingAccountBalance struct {
	Asset            string `json:"asset" bson:"asset"`
	Balance          string `json:"balance" bson:"balance"`
	AvailableBalance string `json:"availableBalance" bson:"availableBalance"`
	FrozenBalance    string `json:"frozenBalance" bson:"frozenBalance"`
}

type InstrumentsInfo struct {
	Symbol         string `json:"symbol" bson:"symbol"`                 // +
	Base           string `json:"base" bson:"base"`                     // +
	Quote          string `json:"quote" bson:"quote"`                   // +
	MinQty         string `json:"minQty" bson:"minQty"`                 // + размер монеты минимальный
	MinNotional    string `json:"minNotional" bson:"minNotional"`       // (не обязательный) размер в доларах минимальный
	PricePrecision string `json:"pricePrecision" bson:"pricePrecision"` // Отправляем если есть
	SizePrecision  string `json:"sizePrecision" bson:"sizePrecision"`   // Отправляем если есть
	State          string `json:"state" bson:"state"`                   // enum  LIVE и другие
	// InstType           string `json:"instType" bson:"instType"`
	StepTickPrice      string `json:"stepTickPrice" bson:"stepTickPrice"`
	StepContractSize   string `json:"stepContractSize" bson:"stepContractSize"`
	MinContractSize    string `json:"minContractSize" bson:"minContractSize"`
	MaxLeverage        string `json:"maxLeverage" bson:"maxLeverage"`
	ContractSize       string `json:"contractSize" bson:"contractSize"`
	ContractMultiplier string `json:"contractMultiplier" bson:"contractMultiplier"`
}

type OrdersPendingList struct {
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
