package huobi

import (
	"fmt"
	"strings"

	"github.com/Betarost/onetrades/entity"
	"github.com/Betarost/onetrades/utils"
)

// ===============ACCOUNT=================
type account_converts struct{}

func (c *account_converts) convertAccountInfo(in accountInfo) (out entity.AccountInformation) {

	out.UID = fmt.Sprintf("%d", in.ID)
	// out.Label = in.Label
	// out.IP = in.Ip
	// out.PermSpot = true

	// if strings.Contains(in.Perm, "read") {
	// 	out.CanRead = true
	// }

	// if strings.Contains(in.Perm, "trade") {
	// 	out.CanTrade = true
	// }

	// if in.PosMode == "long_short_mode" {
	// 	// out.HedgeMode = true
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

		if item.Status == "online" {
			item.Status = "LIVE"
		}

		rec := entity.Spot_InstrumentsInfo{
			Symbol: strings.ToUpper(item.Symbol),
			Base:   strings.ToUpper(item.BaseCoin),
			Quote:  strings.ToUpper(item.QuoteCoin),
			// MinQty:         utils.FloatToStringAll(item.MinQty),
			// MinNotional:    item.MinTradeUSDT,
			PricePrecision: utils.FloatToStringAll(item.PricePrecision),
			SizePrecision:  utils.FloatToStringAll(item.QuantityPrecision),
			State:          item.Status,
		}
		out = append(out, rec)
	}
	return
}

func (c *spot_converts) convertBalance(in []spot_Balance) (out []entity.AssetsBalance) {
	if len(in) == 0 {
		return out
	}

	mapsAssets := map[string]entity.AssetsBalance{}

	for _, item := range in {
		t, is := mapsAssets[item.Currency]
		if !is {
			t = entity.AssetsBalance{Asset: item.Currency}

		}

		if item.Type == "trade" {
			t.Balance = item.Balance
		} else if item.Type == "frozen" {
			t.Locked = item.Balance
		}
		mapsAssets[item.Currency] = t

	}

	for _, i := range mapsAssets {
		if i.Balance != "0" || i.Locked != "0" {
			out = append(out, i)
		}
	}
	return out
}

func (c *spot_converts) convertOrdersHistory(in []spot_ordersHistory_Response) (out []entity.Spot_OrdersHistory) {
	if len(in) == 0 {
		return out
	}

	for _, item := range in {

		side := ""
		typeOrd := ""

		sp := strings.Split(item.Type, "-")
		if len(sp) == 2 {
			side = sp[0]
			typeOrd = sp[1]
		}

		stat := item.State

		if item.State == "submitted" {
			stat = "FILLED"
		}
		out = append(out, entity.Spot_OrdersHistory{
			Symbol:        strings.ToUpper(item.Symbol),
			OrderID:       fmt.Sprintf("%d", item.ID),
			ClientOrderID: item.Client_order_id,
			Side:          strings.ToUpper(side),
			Size:          item.Amount,
			Price:         item.Price,
			ExecutedSize:  item.Filled_amount,
			// ExecutedPrice: item.Fill_price,
			// Fee:           item.Fee,
			Type:       strings.ToUpper(typeOrd),
			Status:     stat,
			CreateTime: item.Created_at,
			UpdateTime: item.Updated_at,
		})
	}
	return out
}

func (c *spot_converts) convertPlaceOrder(in placeOrder_Response) (out []entity.PlaceOrder) {

	// out = append(out, entity.PlaceOrder{
	// 	OrderID:       in.ID,
	// 	ClientOrderID: in.Text,
	// 	Ts:            time.Now().UTC().UnixMilli(),
	// })
	return out
}

func (c *spot_converts) convertOrderList(in []spot_orderList) (out []entity.Spot_OrdersList) {
	if len(in) == 0 {
		return out
	}

	for _, item := range in {

		side := ""
		typeOrd := ""

		sp := strings.Split(item.Type, "-")
		if len(sp) == 2 {
			side = sp[0]
			typeOrd = sp[1]
		}
		out = append(out, entity.Spot_OrdersList{
			Symbol:        strings.ToUpper(item.Symbol),
			OrderID:       fmt.Sprintf("%d", item.ID),
			ClientOrderID: item.Client_order_id,
			Side:          strings.ToUpper(side),
			Size:          item.Amount,
			Price:         item.Price,
			ExecutedSize:  item.Filled_amount,
			Type:          strings.ToUpper(typeOrd),
			Status:        item.State,
			CreateTime:    item.Created_at,
			// UpdateTime:    utils.StringToInt64(item.UTime),
		})
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
		state := "OTHER"
		if item.Contract_status == 1 {
			state = "LIVE"
		}
		rec := entity.Futures_InstrumentsInfo{
			Symbol:         item.Contract_code,
			Base:           item.Symbol,
			Quote:          item.Trade_partition,
			MinQty:         "1",
			PricePrecision: utils.GetPrecisionFromStr(utils.FloatToStringAll(item.Price_tick)),
			SizePrecision:  utils.GetPrecisionFromStr(utils.FloatToStringAll(item.Contract_size)),
			State:          state,
			ContractSize:   utils.FloatToStringAll(item.Contract_size),
			Multiplier:     "1",
			IsSizeContract: true,
		}
		out = append(out, rec)
	}
	return
}

func (c *futures_converts) convertBalance(in futures_Balance) (out []entity.FuturesBalance) {

	out = append(out, entity.FuturesBalance{
		Asset:   in.Currency,
		Balance: in.Total,
		// Equity:  in.Equity,
		Available:        in.Available,
		UnrealizedProfit: in.Unrealised_pnl,
	})
	return out
}
