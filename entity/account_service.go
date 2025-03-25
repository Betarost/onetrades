package entity

type AccountInfo struct {
	UID       string `json:"uid" bson:"uid"`
	MainUID   string `json:"mainUID" bson:"mainUID"`
	IsMain    bool   `json:"isMain" bson:"isMain"`
	Label     string `json:"label" bson:"label"`
	Level     string `json:"level" bson:"level"`
	CanRead   bool   `json:"canRead" bson:"canRead"`
	CanTrade  bool   `json:"canTrade" bson:"canTrade"`
	HedgeMode bool   `json:"hedgeMode" bson:"hedgeMode"`
}

type AccountBalance struct {
	Asset              string  `json:"asset" bson:"asset"`
	Balance            float64 `json:"balance" bson:"balance"`
	EquityBalance      float64 `json:"equityBalance" bson:"equityBalance"`
	AvailableBalance   float64 `json:"availableBalance" bson:"availableBalance"`
	CrossWalletBalance float64 `json:"crossWalletBalance" bson:"crossWalletBalance"`
	UnrealizedProfit   float64 `json:"unrealizedProfit" bson:"unrealizedProfit"`
}

type Position struct {
	Symbol           string  `json:"symbol" bson:"symbol"`
	PositionSide     string  `json:"positionSide" bson:"positionSide"`
	PositionAmt      float64 `json:"positionAmt" bson:"positionAmt"`
	PositionID       string  `json:"positionID"`
	EntryPrice       float64 `json:"entryPrice" bson:"entryPrice"`
	MarkPrice        float64 `json:"markPrice" bson:"markPrice"`
	UnRealizedProfit float64 `json:"unRealizedProfit" bson:"unRealizedProfit"`
	RealizedProfit   float64 `json:"realizedProfit" bson:"realizedProfit"`
	Notional         float64 `json:"notional" bson:"notional"`
	InitialMargin    float64 `json:"initialMargin" bson:"initialMargin"`
	UpdateTime       int64   `json:"updateTime" bson:"updateTime"`

	MarginRatio      string `json:"marginRatio" bson:"marginRatio"`
	AutoDeleveraging string `json:"autoDeleveraging" bson:"autoDeleveraging"`
}

type HistoryPosition struct {
	PositionID            string             `json:"positionID"`
	Symbol                string             `json:"symbol" bson:"symbol"`
	Status                PositionStatusType `json:"status" bson:"status"`
	PositionSide          string             `json:"positionSide" bson:"positionSide"`
	AvgOpenPrice          float64            `json:"avgOpenPrice" bson:"avgOpenPrice"`
	AvgClosePrice         float64            `json:"avgClosePrice" bson:"avgClosePrice"`
	PositionOpenAmt       float64            `json:"positionOpenAmt" bson:"positionOpenAmt"`
	PositionCloseAmt      float64            `json:"positionCloseAmt" bson:"positionCloseAmt"`
	PositionOpenNotional  float64            `json:"positionOpenNotional" bson:"positionOpenNotional"`
	PositionCloseNotional float64            `json:"positionCloseNotional" bson:"positionCloseNotional"`
	RealizedProfit        float64            `json:"realizedProfit" bson:"realizedProfit"`
	Pnl                   float64            `json:"pnl" bson:"pnl"`
	PnlRatio              float64            `json:"pnlRatio" bson:"pnlRatio"`
	Fee                   float64            `json:"fee" bson:"fee"`
	FundingFee            float64            `json:"fundingFee" bson:"fundingFee"`
	LiqPenalty            float64            `json:"liqPenalty" bson:"liqPenalty"`
	CreateTime            int64              `json:"createTime" bson:"createTime"`
	UpdateTime            int64              `json:"updateTime" bson:"updateTime"`
}
