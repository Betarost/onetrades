package entity

type PositionModeType string
type PositionSideType string
type SideType string
type OrderType string

const (
	PositionModeTypeOneWay PositionModeType = "ONE_WAY"
	PositionModeTypeHedge  PositionModeType = "HEDGE"

	PositionSideTypeLong  PositionSideType = "LONG"
	PositionSideTypeShort PositionSideType = "SHORT"

	SideTypeBuy  SideType = "BUY"
	SideTypeSell SideType = "SELL"

	OrderTypeLimit  OrderType = "LIMIT"
	OrderTypeMarket OrderType = "MARKET"
	// OrderTypeLimitMaker      OrderType = "LIMIT_MAKER"
	// OrderTypeStopLoss        OrderType = "STOP_LOSS"
	// OrderTypeStopLossLimit   OrderType = "STOP_LOSS_LIMIT"
	// OrderTypeTakeProfit      OrderType = "TAKE_PROFIT"
	// OrderTypeTakeProfitLimit OrderType = "TAKE_PROFIT_LIMIT"
)
