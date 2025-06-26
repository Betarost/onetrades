package gateio

import (
	"fmt"
	"strings"
	"time"

	"github.com/Betarost/onetrades/entity"
	"github.com/Betarost/onetrades/utils"
)

// ===============ACCOUNT=================
type account_converts struct{}

func (c *account_converts) convertAccountInfo(in accountInfo) (out entity.AccountInformation) {

	out.UID = fmt.Sprintf("%d", in.User_id)
	out.IP = strings.Join(in.Ip_whitelist, ",")
	// out.PermSpot = true

	// if strings.Contains(in.Perm, "read") {
	// 	out.CanRead = true
	// }

	// if strings.Contains(in.Perm, "trade") {
	// 	out.CanTrade = true
	// }

	// if in.PosMode == "long_short_mode" {
	// 	out.HedgeMode = true
	// }

	// if in.AcctLv != "1" {
	// 	out.PermFutures = true
	// }
	return out
}

// ===============SPOT=================

type spot_converts struct{}

func (c *spot_converts) convertInstrumentsInfo(in []spot_instrumentsInfo) (out []entity.Spot_InstrumentsInfo) {
	if len(in) == 0 {
		return out
	}
	for _, item := range in {
		if item.Trade_status == "tradable" {
			item.Trade_status = "LIVE"
		}

		rec := entity.Spot_InstrumentsInfo{
			Symbol:         item.ID,
			Base:           item.Base,
			Quote:          item.Quote,
			State:          item.Trade_status,
			MinQty:         item.Min_base_amount,
			MinNotional:    item.Min_quote_amount,
			PricePrecision: fmt.Sprintf("%d", item.Precision),
			SizePrecision:  fmt.Sprintf("%d", item.Amount_precision),
		}
		out = append(out, rec)
	}
	return
}

func (c *spot_converts) convertBalance(in []spot_Balance) (out []entity.AssetsBalance) {
	if len(in) == 0 {
		return out
	}
	for _, item := range in {
		out = append(out, entity.AssetsBalance{
			Asset:   item.Currency,
			Balance: item.Available,
			Locked:  item.Locked,
		})
	}
	return out
}

func (c *spot_converts) convertPlaceOrder(in placeOrder_Response) (out []entity.PlaceOrder) {

	out = append(out, entity.PlaceOrder{
		OrderID:       in.ID,
		ClientOrderID: in.Text,
		Ts:            time.Now().UTC().UnixMilli(),
	})
	return out
}

func (c *spot_converts) convertOrderList(in []spot_orderList) (out []entity.Spot_OrdersList) {
	if len(in) == 0 {
		return out
	}
	for _, i := range in {
		for _, item := range i.Orders {

			out = append(out, entity.Spot_OrdersList{
				Symbol:        item.Currency_pair,
				OrderID:       item.ID,
				ClientOrderID: item.Text,
				Side:          strings.ToUpper(item.Side),
				Size:          item.Amount,
				Price:         item.Price,
				ExecutedSize:  item.Filled_amount,
				Type:          strings.ToUpper(item.Type),
				Status:        item.Status,
				CreateTime:    item.Create_time,
				UpdateTime:    item.Update_time,
			})
		}
	}
	return out
}

// ===============FUTURES=================

type futures_converts struct{}

func (c *futures_converts) convertInstrumentsInfo(in []futures_instrumentsInfo) (out []entity.Futures_InstrumentsInfo) {
	if len(in) == 0 {
		return out
	}
	for _, item := range in {
		if item.Status == "trading" {
			item.Status = "LIVE"
		}
		if item.In_delisting {
			item.Status = "OFF"
		}

		base := ""
		quote := ""

		sp := strings.Split(item.Name, "_")
		if len(sp) == 2 {
			base = sp[0]
			quote = sp[1]
		}

		rec := entity.Futures_InstrumentsInfo{
			Symbol:         item.Name,
			Base:           base,
			Quote:          quote,
			MinQty:         fmt.Sprintf("%d", item.Order_size_min),
			PricePrecision: utils.GetPrecisionFromStr(item.Mark_price_round),
			SizePrecision:  "0",
			MaxLeverage:    item.Leverage_max,
			State:          item.Status,
			IsSizeContract: true,
			Multiplier:     item.Quanto_multiplier,
			ContractSize:   item.Quanto_multiplier,
		}
		out = append(out, rec)
	}
	return
}

func (c *futures_converts) convertPlaceOrder(in futures_placeOrder_Response) (out []entity.PlaceOrder) {

	out = append(out, entity.PlaceOrder{
		OrderID:       fmt.Sprintf("%d", in.ID),
		ClientOrderID: in.Text,
		Ts:            int64(in.Create_time),
	})
	return out
}

func (c *futures_converts) convertPositions(answ []futures_Position) (res []entity.Futures_Positions) {
	for _, item := range answ {
		positionSide := "LONG"
		if item.Mode == "single" {
			if item.Size < 0 {
				positionSide = "SHORT"
			}
		} else {
			positionSide = strings.ToUpper(item.Mode)
		}

		res = append(res, entity.Futures_Positions{
			Symbol:       item.Contract,
			PositionSide: positionSide,
			// PositionID:       item.PosID,
			PositionAmt:      fmt.Sprintf("%d", item.Size),
			EntryPrice:       item.Entry_price,
			MarkPrice:        item.Mark_price,
			InitialMargin:    item.Initial_margin,
			UnRealizedProfit: item.Unrealised_pnl,
			RealizedProfit:   item.Realised_pnl,
			Notional:         item.Value,
			MarginRatio:      item.Maintenance_rate,
			UpdateTime:       item.Update_time,
		})
	}
	return res
}

func (c *futures_converts) convertOrderList(answ []futures_orderList) (res []entity.Futures_OrdersList) {
	for _, item := range answ {
		positionSide := "LONG"
		// if item.PosSide == "net" {
		// 	if strings.ToUpper(item.Side) == "SELL" {
		// 		positionSide = "SHORT"
		// 	}
		// } else {
		// 	positionSide = strings.ToUpper(item.PosSide)
		// }

		res = append(res, entity.Futures_OrdersList{
			Symbol:        item.Contract,
			OrderID:       fmt.Sprintf("%d", item.ID),
			ClientOrderID: item.Text,
			PositionSide:  positionSide,
			// Side:          item.Side,
			PositionAmt: fmt.Sprintf("%d", item.Size),
			Price:       item.Price,
			// TpPrice:       tp,
			// SlPrice:       sl,
			// Type:          strings.ToUpper(item.OrdType),
			// TradeMode:     item.TdMode,
			// InstType:      item.InstType,
			// Leverage:      item.Lever,
			// Status:        item.State,
			// IsTpLimit:     b,
			CreateTime: int64(item.Create_time),
			UpdateTime: int64(item.Update_time),
		})
	}
	return res
}
