package entity

type PositionModeType string
type PositionSideType string
type SideType string
type OrderType string
type TimeFrameType string

const (
	PositionModeTypeOneWay PositionModeType = "ONE_WAY"
	PositionModeTypeHedge  PositionModeType = "HEDGE"

	PositionSideTypeLong  PositionSideType = "LONG"
	PositionSideTypeShort PositionSideType = "SHORT"

	SideTypeBuy  SideType = "BUY"
	SideTypeSell SideType = "SELL"

	OrderTypeLimit  OrderType = "LIMIT"
	OrderTypeMarket OrderType = "MARKET"

	TimeFrameType1m  TimeFrameType = "1m"
	TimeFrameType3m  TimeFrameType = "3m"
	TimeFrameType5m  TimeFrameType = "5m"
	TimeFrameType15m TimeFrameType = "15m"
	TimeFrameType30m TimeFrameType = "30m"
	TimeFrameType1H  TimeFrameType = "1H"
	TimeFrameType2H  TimeFrameType = "2H"
	TimeFrameType4H  TimeFrameType = "4H"
	TimeFrameType8H  TimeFrameType = "8H"
	TimeFrameType1D  TimeFrameType = "1D"

	// OrderTypeLimitMaker      OrderType = "LIMIT_MAKER"
	// OrderTypeStopLoss        OrderType = "STOP_LOSS"
	// OrderTypeStopLossLimit   OrderType = "STOP_LOSS_LIMIT"
	// OrderTypeTakeProfit      OrderType = "TAKE_PROFIT"
	// OrderTypeTakeProfitLimit OrderType = "TAKE_PROFIT_LIMIT"
)
