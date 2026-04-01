package bitget

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/Betarost/onetrades/entity"
	"github.com/Betarost/onetrades/utils"
)

type futures_ordersHistory struct {
	callAPI func(ctx context.Context, r *utils.Request, opts ...utils.RequestOption) (data []byte, header *http.Header, err error)
	convert futures_converts

	symbol    *string
	startTime *int64
	endTime   *int64
	limit     *int64
	page      *int64
}

func (s *futures_ordersHistory) Symbol(symbol string) *futures_ordersHistory {
	s.symbol = &symbol
	return s
}

func (s *futures_ordersHistory) StartTime(startTime int64) *futures_ordersHistory {
	s.startTime = &startTime
	return s
}

func (s *futures_ordersHistory) EndTime(endTime int64) *futures_ordersHistory {
	s.endTime = &endTime
	return s
}

func (s *futures_ordersHistory) Limit(limit int64) *futures_ordersHistory {
	s.limit = &limit
	return s
}

func (s *futures_ordersHistory) Page(page int64) *futures_ordersHistory {
	s.page = &page
	return s
}

func (s *futures_ordersHistory) Do(ctx context.Context, opts ...utils.RequestOption) (res []entity.Futures_OrdersHistory, err error) {
	// ------------------------------------------------
	// 1) Обычная история ордеров
	// ------------------------------------------------
	r1 := &utils.Request{
		Method:   http.MethodGet,
		Endpoint: "/api/v2/mix/order/orders-history",
		SecType:  utils.SecTypeSigned,
	}

	m1 := utils.Params{
		"productType": "USDT-FUTURES",
		"marginCoin":  "USDT",
	}

	if s.symbol != nil {
		m1["symbol"] = *s.symbol
	}
	if s.limit != nil && *s.limit > 0 {
		m1["limit"] = *s.limit
	}
	if s.startTime != nil {
		m1["startTime"] = *s.startTime
	}
	if s.endTime != nil {
		m1["endTime"] = *s.endTime
	}

	r1.SetParams(m1)

	data1, _, err := s.callAPI(ctx, r1, opts...)
	if err != nil {
		return res, err
	}

	var answ1 struct {
		Result struct {
			Orders []futures_ordersHistory_Response `json:"entrustedList"`
		} `json:"data"`
	}

	err = json.Unmarshal(data1, &answ1)
	if err != nil {
		return res, err
	}

	out := convertOrdersHistoryBitget(answ1.Result.Orders)

	// ------------------------------------------------
	// 2) История TP/SL plan orders
	// ------------------------------------------------
	r2 := &utils.Request{
		Method:   http.MethodGet,
		Endpoint: "/api/v2/mix/order/orders-plan-history",
		SecType:  utils.SecTypeSigned,
	}

	m2 := utils.Params{
		"productType": "USDT-FUTURES",
		"marginCoin":  "USDT",
		"planType":    "profit_loss",
	}

	if s.symbol != nil {
		m2["symbol"] = *s.symbol
	}
	if s.limit != nil && *s.limit > 0 {
		m2["limit"] = *s.limit
	}
	if s.startTime != nil {
		m2["startTime"] = *s.startTime
	}
	if s.endTime != nil {
		m2["endTime"] = *s.endTime
	}

	r2.SetParams(m2)

	data2, _, err := s.callAPI(ctx, r2, opts...)
	if err != nil {
		return res, err
	}

	var answ2 struct {
		Result struct {
			Orders []futures_planOrdersHistory_Response `json:"entrustedList"`
		} `json:"data"`
	}

	err = json.Unmarshal(data2, &answ2)
	if err != nil {
		return res, err
	}

	planOut := convertPlanOrdersHistoryBitget(answ2.Result.Orders)

	// ------------------------------------------------
	// 3) Merge без дублей:
	//    если trigger order уже породил обычный ордер и он есть
	//    в orders-history, то просто проставляем ему TpOrder/SlOrder.
	// ------------------------------------------------
	indexByOrderID := make(map[string]int, len(out))
	for i := range out {
		if out[i].OrderID != "" {
			indexByOrderID[out[i].OrderID] = i
		}
	}

	for _, p := range planOut {
		if idx, ok := indexByOrderID[p.OrderID]; ok {
			if p.TpOrder {
				out[idx].TpOrder = true
			}
			if p.SlOrder {
				out[idx].SlOrder = true
			}

			// На случай если в обычной истории чего-то нет, а в plan history есть
			if out[idx].ClientOrderID == "" {
				out[idx].ClientOrderID = p.ClientOrderID
			}
			if out[idx].PositionSide == "" {
				out[idx].PositionSide = p.PositionSide
			}
			if out[idx].Side == "" {
				out[idx].Side = p.Side
			}
			if out[idx].MarginMode == "" {
				out[idx].MarginMode = p.MarginMode
			}
			if out[idx].CreateTime == 0 {
				out[idx].CreateTime = p.CreateTime
			}
			if out[idx].UpdateTime == 0 {
				out[idx].UpdateTime = p.UpdateTime
			}
			continue
		}

		indexByOrderID[p.OrderID] = len(out)
		out = append(out, p)
	}

	return out, nil
}

type futures_ordersHistory_Response struct {
	Symbol       string `json:"symbol"`
	Size         string `json:"size"`
	OrderId      string `json:"orderId"`
	ClientOid    string `json:"clientOid"`
	BaseVolume   string `json:"baseVolume"`
	Fee          string `json:"fee"`
	Price        string `json:"price"`
	PriceAvg     string `json:"priceAvg"`
	Status       string `json:"status"`
	Side         string `json:"side"`
	TotalProfits string `json:"totalProfits"`
	PosSide      string `json:"posSide"`
	Leverage     string `json:"leverage"`
	MarginMode   string `json:"marginMode"`
	PosMode      string `json:"posMode"`
	OrderType    string `json:"orderType"`
	CTime        string `json:"cTime"`
	UTime        string `json:"uTime"`
}

type futures_planOrdersHistory_Response struct {
	PlanType       string `json:"planType"`
	Symbol         string `json:"symbol"`
	Size           string `json:"size"`
	OrderId        string `json:"orderId"`        // plan order id
	ExecuteOrderId string `json:"executeOrderId"` // actual order id after trigger
	ClientOid      string `json:"clientOid"`
	PlanStatus     string `json:"planStatus"` // executed / fail_execute / cancelled
	Price          string `json:"price"`
	ExecutePrice   string `json:"executePrice"`
	PriceAvg       string `json:"priceAvg"`
	BaseVolume     string `json:"baseVolume"`
	TriggerPrice   string `json:"triggerPrice"`
	Side           string `json:"side"`
	PosSide        string `json:"posSide"`
	MarginMode     string `json:"marginMode"`
	TradeSide      string `json:"tradeSide"`
	PosMode        string `json:"posMode"`
	OrderType      string `json:"orderType"` // limit / market
	CTime          string `json:"cTime"`
	UTime          string `json:"uTime"`
}

func convertOrdersHistoryBitget(in []futures_ordersHistory_Response) (out []entity.Futures_OrdersHistory) {
	if len(in) == 0 {
		return out
	}

	for _, item := range in {
		status := strings.ToUpper(item.Status)
		if status != "FILLED" {
			continue
		}

		marginMode := string(entity.MarginModeTypeIsolated)
		hedgeMode := false

		if item.PosMode != "one_way_mode" {
			hedgeMode = true
		}

		if item.MarginMode == "crossed" {
			marginMode = string(entity.MarginModeTypeCross)
		}

		positionSide := "LONG"
		if item.PosSide == "net" {
			if strings.ToUpper(item.Side) == "SELL" {
				positionSide = "SHORT"
			}
		} else {
			positionSide = strings.ToUpper(item.PosSide)
		}

		out = append(out, entity.Futures_OrdersHistory{
			Symbol:         item.Symbol,
			OrderID:        item.OrderId,
			ClientOrderID:  item.ClientOid,
			Side:           strings.ToUpper(item.Side),
			PositionSide:   strings.ToUpper(positionSide),
			PositionSize:   item.Size,
			ExecutedSize:   item.BaseVolume,
			Price:          item.Price,
			ExecutedPrice:  item.PriceAvg,
			RealisedProfit: item.TotalProfits,
			Fee:            item.Fee,
			Type:           strings.ToUpper(item.OrderType),
			Leverage:       item.Leverage,
			Status:         strings.ToUpper(item.Status),
			HedgeMode:      hedgeMode,
			MarginMode:     marginMode,
			CreateTime:     utils.StringToInt64(item.CTime),
			UpdateTime:     utils.StringToInt64(item.UTime),
		})
	}

	return out
}

func convertPlanOrdersHistoryBitget(in []futures_planOrdersHistory_Response) (out []entity.Futures_OrdersHistory) {
	if len(in) == 0 {
		return out
	}

	for _, item := range in {
		if strings.ToLower(item.PlanStatus) != "executed" {
			continue
		}

		planType := strings.ToLower(item.PlanType)

		isTP := false
		isSL := false

		switch planType {
		case "profit_plan", "pos_profit":
			isTP = true
		case "loss_plan", "pos_loss":
			isSL = true
		default:
			// Если вдруг придет не TP/SL внутри profit_loss, пропускаем
			continue
		}

		marginMode := string(entity.MarginModeTypeIsolated)
		if strings.ToLower(item.MarginMode) == "crossed" {
			marginMode = string(entity.MarginModeTypeCross)
		}

		hedgeMode := false
		if item.PosMode != "" && strings.ToLower(item.PosMode) != "one_way_mode" {
			hedgeMode = true
		}

		side := strings.ToUpper(item.Side)
		positionSide := "LONG"

		if strings.ToUpper(item.PosSide) == "LONG" && strings.ToUpper(item.TradeSide) == "CLOSE" {
			side = "SELL"
		} else if strings.ToUpper(item.PosSide) == "SHORT" && strings.ToUpper(item.TradeSide) == "CLOSE" {
			side = "BUY"
		}

		if strings.ToLower(item.PosSide) == "net" {
			if strings.ToUpper(side) == "SELL" {
				positionSide = "SHORT"
			}
		} else if item.PosSide != "" {
			positionSide = strings.ToUpper(item.PosSide)
		}

		// Для triggered plan order используем executeOrderId как основной OrderID,
		// чтобы можно было матчить с обычной history и не плодить дубли.
		orderID := item.OrderId
		if item.ExecuteOrderId != "" {
			orderID = item.ExecuteOrderId
		}

		price := item.TriggerPrice
		if price == "" {
			price = item.Price
		}
		if price == "" {
			price = item.ExecutePrice
		}

		executedPrice := item.PriceAvg
		if executedPrice == "" {
			executedPrice = item.ExecutePrice
		}

		out = append(out, entity.Futures_OrdersHistory{
			Symbol:        item.Symbol,
			OrderID:       orderID,
			ClientOrderID: item.ClientOid,
			Side:          side,
			PositionSide:  positionSide,
			PositionSize:  item.Size,
			ExecutedSize:  item.BaseVolume,
			Price:         price,
			ExecutedPrice: executedPrice,
			Type:          strings.ToUpper(item.OrderType),
			Status:        strings.ToUpper(item.PlanStatus),
			HedgeMode:     hedgeMode,
			MarginMode:    marginMode,
			CreateTime:    utils.StringToInt64(item.CTime),
			UpdateTime:    utils.StringToInt64(item.UTime),
			TpOrder:       isTP,
			SlOrder:       isSL,
		})
	}

	return out
}
