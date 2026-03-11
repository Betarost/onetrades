package hyperliquid

import (
	"encoding/json"
	"fmt"
	"math"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/Betarost/onetrades/entity"
	"github.com/Betarost/onetrades/utils"
)

type account_converts struct{}

func (c *account_converts) convertAccountInfo(_ hyperliquidAccountInfo) (out entity.AccountInformation) {
	out.CanRead = true
	out.CanTrade = true
	out.CanTransfer = false
	out.PermSpot = true
	out.PermFutures = true
	out.UID = ""
	return out
}

type spot_converts struct{}

func (c *spot_converts) convertInstrumentsInfo(in spotMetaResponse) (out []entity.Spot_InstrumentsInfo) {
	if len(in.Universe) == 0 || len(in.Tokens) == 0 {
		return out
	}

	tokenByIndex := make(map[int]spotMetaToken, len(in.Tokens))
	for _, t := range in.Tokens {
		tokenByIndex[t.Index] = t
	}

	for _, u := range in.Universe {
		if len(u.Tokens) < 2 {
			continue
		}

		baseTok, okBase := tokenByIndex[u.Tokens[0]]
		quoteTok, okQuote := tokenByIndex[u.Tokens[1]]
		if !okBase || !okQuote {
			continue
		}

		symbol := u.Name
		if !strings.Contains(symbol, "/") {
			symbol = baseTok.Name + "/" + quoteTok.Name
		}

		// minQty := pow10Str(-baseTok.SzDecimals)
		sizePrec := strconv.Itoa(baseTok.SzDecimals)
		state := "LIVE"

		out = append(out, entity.Spot_InstrumentsInfo{
			Symbol: symbol,
			Base:   baseTok.Name,
			Quote:  quoteTok.Name,
			// MinQty:         minQty,
			MinNotional:    "",
			PricePrecision: "",
			SizePrecision:  sizePrec,
			State:          state,
			TokenId:        strconv.Itoa(10000 + u.Index),
		})
	}

	return out
}

func (c *spot_converts) convertSpotBalances(in spot_spotClearinghouseState) (out []entity.AssetsBalance) {
	for _, b := range in.Balances {
		if utils.StringToFloat(b.Total) == 0 && utils.StringToFloat(b.Hold) == 0 {
			continue
		}
		out = append(out, entity.AssetsBalance{
			Asset:   b.Coin,
			Balance: b.Total,
			Locked:  b.Hold,
		})
	}
	return out
}

func (c *spot_converts) convertPlaceOrderSpot(in spot_placeOrderResponse) (out []entity.PlaceOrder) {
	if len(in.Response.Data.Statuses) == 0 {
		return out
	}

	st := in.Response.Data.Statuses[0]
	if strings.TrimSpace(st.Error) != "" {
		return out
	}

	out = append(out, entity.PlaceOrder{
		OrderID:       stringifyHLID(st.Oid),
		ClientOrderID: strings.TrimSpace(st.Cloid),
		Ts:            time.Now().UTC().UnixMilli(),
	})
	return out
}

func (c *spot_converts) convertSpotOpenOrders(in []hlFrontendOpenOrder) (out []entity.Spot_OrdersList) {
	out = make([]entity.Spot_OrdersList, 0, len(in))

	for _, o := range in {
		assetID, ok := parseSpotAssetIDFromCoin(o.Coin)
		if !ok {
			continue
		}

		side := "BUY"
		if strings.EqualFold(strings.TrimSpace(o.Side), "A") {
			side = "SELL"
		}

		typ := strings.ToUpper(strings.TrimSpace(o.OrderType))
		if typ == "" {
			typ = "LIMIT"
		}

		size := strings.TrimSpace(o.OrigSz)
		if size == "" {
			size = strings.TrimSpace(o.Sz)
		}

		executed := "0"
		if strings.TrimSpace(o.OrigSz) != "" && strings.TrimSpace(o.Sz) != "" {
			if ex, err := decSub(o.OrigSz, o.Sz); err == nil {
				executed = ex
			}
		}

		out = append(out, entity.Spot_OrdersList{
			Symbol:        assetID,
			OrderID:       strconv.FormatInt(o.Oid, 10),
			ClientOrderID: o.clientOrderID(),
			Side:          side,
			Size:          size,
			ExecutedSize:  executed,
			Price:         strings.TrimSpace(o.LimitPx),
			Type:          typ,
			Status:        "NEW",
			CreateTime:    o.Timestamp,
			UpdateTime:    o.Timestamp,
		})
	}

	return out
}

func (c *spot_converts) convertCancelOrderSpot(_ spot_placeOrderResponse, oid int64) (out []entity.PlaceOrder) {
	out = append(out, entity.PlaceOrder{
		OrderID: strconv.FormatInt(oid, 10),
		Ts:      time.Now().UTC().UnixMilli(),
	})
	return out
}

func (c *spot_converts) convertSpotOrdersHistory(in []hlHistoricalOrder) (out []entity.Spot_OrdersHistory) {
	if len(in) == 0 {
		return out
	}

	out = make([]entity.Spot_OrdersHistory, 0, len(in))

	for _, item := range in {
		assetID, ok := parseSpotAssetIDFromHistoricalCoin(item.Order.Coin)
		if !ok {
			continue
		}

		status := strings.ToUpper(strings.TrimSpace(item.Status))
		if status != "FILLED" {
			continue
		}

		side := "BUY"
		if strings.EqualFold(strings.TrimSpace(item.Order.Side), "A") {
			side = "SELL"
		}

		orderType := strings.ToUpper(strings.TrimSpace(item.Order.OrderType))
		if orderType == "" {
			orderType = "LIMIT"
		}

		size := strings.TrimSpace(item.Order.OrigSz)
		if size == "" {
			size = strings.TrimSpace(item.Order.Sz)
		}

		createTime := item.Order.Timestamp
		updateTime := item.StatusTimestamp
		if updateTime == 0 {
			updateTime = createTime
		}

		out = append(out, entity.Spot_OrdersHistory{
			Symbol:        assetID,
			OrderID:       strconv.FormatInt(item.Order.Oid, 10),
			ClientOrderID: item.clientOrderID(),
			Side:          side,
			Size:          size,
			Price:         strings.TrimSpace(item.Order.LimitPx),
			ExecutedSize:  size,
			ExecutedPrice: strings.TrimSpace(item.Order.LimitPx),
			Type:          orderType,
			Status:        status,
			CreateTime:    createTime,
			UpdateTime:    updateTime,
		})
	}

	return out
}

type spotMetaResponse struct {
	Tokens   []spotMetaToken    `json:"tokens"`
	Universe []spotMetaUniverse `json:"universe"`
}

type spotMetaToken struct {
	Name        string `json:"name"`
	SzDecimals  int    `json:"szDecimals"`
	WeiDecimals int    `json:"weiDecimals"`
	Index       int    `json:"index"`
	TokenId     string `json:"tokenId"`
	IsCanonical bool   `json:"isCanonical"`
}

type spotMetaUniverse struct {
	Name        string `json:"name"`
	Tokens      []int  `json:"tokens"`
	Index       int    `json:"index"`
	IsCanonical bool   `json:"isCanonical"`
}

type futures_converts struct{}

func (c *futures_converts) convertInstrumentsInfo(in perpMetaResponse) (out []entity.Futures_InstrumentsInfo) {
	if len(in.Universe) == 0 {
		return out
	}

	for _, u := range in.Universe {
		state := "LIVE"
		if u.IsDelisted {
			state = "OFFLINE"
		}

		out = append(out, entity.Futures_InstrumentsInfo{
			Symbol:         u.Name,
			Base:           u.Name,
			Quote:          "USDC",
			MinQty:         pow10Str(-u.SzDecimals),
			MinNotional:    "",
			PricePrecision: "",
			SizePrecision:  strconv.Itoa(u.SzDecimals),
			State:          state,
			MaxLeverage:    utils.Int64ToString(int64(u.MaxLeverage)),
			Multiplier:     "1",
			ContractSize:   "1",
			IsSizeContract: false,
		})
	}

	return out
}

type perpMetaResponse struct {
	Universe []perpMetaUniverse `json:"universe"`
}

type perpMetaUniverse struct {
	Name         string `json:"name"`
	SzDecimals   int    `json:"szDecimals"`
	MaxLeverage  int    `json:"maxLeverage"`
	OnlyIsolated bool   `json:"onlyIsolated,omitempty"`
	IsDelisted   bool   `json:"isDelisted,omitempty"`
}

func pow10Str(exp int) string {
	return utils.FloatToStringAll(math.Pow10(exp))
}

func stringifyHLID(v interface{}) string {
	if v == nil {
		return ""
	}

	switch t := v.(type) {
	case string:
		return t
	case json.Number:
		return t.String()
	case float64:
		return strconv.FormatInt(int64(t), 10)
	case int:
		return strconv.Itoa(t)
	case int64:
		return utils.Int64ToString(t)
	case uint64:
		return strconv.FormatUint(t, 10)
	default:
		b, _ := json.Marshal(t)
		return string(b)
	}
}

func decSub(a, b string) (string, error) {
	a = strings.TrimSpace(a)
	b = strings.TrimSpace(b)
	if a == "" || b == "" {
		return "", fmt.Errorf("empty a/b")
	}

	ra := new(big.Rat)
	if _, ok := ra.SetString(a); !ok {
		return "", fmt.Errorf("bad decimal %q", a)
	}
	rb := new(big.Rat)
	if _, ok := rb.SetString(b); !ok {
		return "", fmt.Errorf("bad decimal %q", b)
	}

	ra.Sub(ra, rb)

	s := ra.FloatString(18)
	s = strings.TrimRight(s, "0")
	s = strings.TrimRight(s, ".")
	if s == "" {
		return "0", nil
	}
	return s, nil
}
