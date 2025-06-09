package entity

type AccountInformation struct {
	UID         string `json:"uid" bson:"uid"`
	Label       string `json:"label" bson:"label"`
	IP          string `json:"ip" bson:"ip"`
	CanRead     bool   `json:"canRead" bson:"canRead"`
	CanTrade    bool   `json:"canTrade" bson:"canTrade"`
	CanTransfer bool   `json:"canTransfer" bson:"canTransfer"`
	PermSpot    bool   `json:"permSpot" bson:"permSpot"`
	PermFutures bool   `json:"permFutures" bson:"permFutures"`
	HedgeMode   bool   `json:"hedgeMode" bson:"hedgeMode"`
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
	Symbol             string `json:"symbol" bson:"symbol"`
	Base               string `json:"base" bson:"base"`
	Quote              string `json:"quote" bson:"quote"`
	InstType           string `json:"instType" bson:"instType"`
	State              string `json:"state" bson:"state"`
	MaxLeverage        string `json:"maxLeverage" bson:"maxLeverage"`
	StepTickPrice      string `json:"stepTickPrice" bson:"stepTickPrice"`
	MinContractSize    string `json:"minContractSize" bson:"minContractSize"`
	StepContractSize   string `json:"stepContractSize" bson:"stepContractSize"`
	ContractSize       string `json:"contractSize" bson:"contractSize"`
	ContractMultiplier string `json:"contractMultiplier" bson:"contractMultiplier"`
}

type OrdersPendingList struct {
	Symbol       string  `json:"symbol" bson:"symbol"`
	OrderID      string  `json:"orderId" bson:"orderId"`
	Side         string  `json:"side" bson:"side"`
	PositionSide string  `json:"positionSide" bson:"positionSide"`
	PositionAmt  float64 `json:"positionAmt" bson:"positionAmt"`
	Price        float64 `json:"price" bson:"price"`
	Notional     float64 `json:"notional" bson:"notional"`
	Type         string  `json:"type" bson:"type"`
	TradeMode    string  `json:"tradeMode" bson:"tradeMode"`
	Status       string  `json:"status" bson:"status"`
	Time         int64   `json:"time" bson:"time"`
	UpdateTime   int64   `json:"updateTime" bson:"updateTime"`
}

type PlaceOrder struct {
	OrderID       string `json:"orderId" bson:"orderId"`
	ClientOrderID string `json:"clientOrderID"`
	Ts            int64  `json:"ts" bson:"ts"`
}
