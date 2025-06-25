package entity

type AccountInformation struct {
	UID      string `json:"uid" bson:"uid"`
	Label    string `json:"label" bson:"label"`
	IP       string `json:"ip" bson:"ip"`
	CanRead  bool   `json:"canRead" bson:"canRead"`
	CanTrade bool   `json:"canTrade" bson:"canTrade"`
	// CanTransfer bool   `json:"canTransfer" bson:"canTransfer"`
	PermSpot    bool `json:"permSpot" bson:"permSpot"`
	PermFutures bool `json:"permFutures" bson:"permFutures"`
	// HedgeMode   bool   `json:"hedgeMode" bson:"hedgeMode"`
}

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

type PlaceOrder struct {
	OrderID       string `json:"orderId" bson:"orderId"`
	ClientOrderID string `json:"clientOrderID" bson:"clientOrderID"`
	Ts            int64  `json:"ts" bson:"ts"`
}

type AssetsBalance struct {
	Asset   string `json:"asset" bson:"asset"`
	Balance string `json:"balance" bson:"balance"`
	Locked  string `json:"locked" bson:"locked"`
}
