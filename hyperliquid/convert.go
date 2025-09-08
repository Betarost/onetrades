package hyperliquid

import (
	"strconv"
	"strings"
	"time"

	"github.com/Betarost/onetrades/entity"
	"github.com/sonirico/go-hyperliquid"
)

// Конвертация сайдов
func convertSide(side string) entity.SideType {
	switch strings.ToLower(side) {
	case "a", "ask", "sell":
		return entity.SideTypeSell
	case "b", "bid", "buy":
		return entity.SideTypeBuy
	default:
		return entity.SideTypeBuy
	}
}

// Конвертация статусов ордеров
func convertOrderStatus(status string) entity.OrderStatusType {
	switch strings.ToLower(status) {
	case "open":
		return entity.OrderStatusTypeNew
	case "filled":
		return entity.OrderStatusTypeFilled
	case "canceled":
		return entity.OrderStatusTypeCanceled
	case "partially_filled":
		return entity.OrderStatusTypePartiallyFilled
	case "rejected":
		return entity.OrderStatusTypeRejected
	default:
		return entity.OrderStatusTypeNew
	}
}

// Конвертация типов ордеров
func convertOrderType(orderType string) entity.OrderType {
	switch strings.ToLower(orderType) {
	case "limit":
		return entity.OrderTypeLimit
	case "market":
		return entity.OrderTypeMarket
	case "stop_limit":
		return entity.OrderTypeStopLimit
	case "stop_market":
		return entity.OrderTypeStopMarket
	default:
		return entity.OrderTypeLimit
	}
}

// Конвертация TimeInForce
func convertTimeInForce(tif string) entity.TimeInForceType {
	switch strings.ToUpper(tif) {
	case "GTC":
		return entity.TimeInForceTypeGTC
	case "IOC":
		return entity.TimeInForceTypeIOC
	case "FOK":
		return entity.TimeInForceTypeFOK
	default:
		return entity.TimeInForceTypeGTC
	}
}

// Конвертация позиций
func convertPosition(pos *hyperliquid.Position) *entity.Position {
	if pos == nil {
		return nil
	}

	size, _ := strconv.ParseFloat(pos.Szi, 64)
	entryPrice, _ := strconv.ParseFloat(pos.EntryPx, 64)
	markPrice, _ := strconv.ParseFloat(pos.MarkPx, 64)
	unrealizedPnl, _ := strconv.ParseFloat(pos.UnrealizedPnl, 64)

	var side entity.PositionSideType
	if size > 0 {
		side = entity.PositionSideTypeLong
	} else if size < 0 {
		side = entity.PositionSideTypeShort
		size = -size // Делаем размер положительным
	} else {
		side = entity.PositionSideTypeBoth
	}

	return &entity.Position{
		Symbol:        pos.Coin,
		Size:          strconv.FormatFloat(size, 'f', -1, 64),
		Side:          side,
		EntryPrice:    strconv.FormatFloat(entryPrice, 'f', -1, 64),
		MarkPrice:     strconv.FormatFloat(markPrice, 'f', -1, 64),
		UnrealizedPnL: strconv.FormatFloat(unrealizedPnl, 'f', -1, 64),
		Percentage:    pos.ReturnOnEquity,
		UpdateTime:    time.Now().UnixMilli(),
	}
}

// Конвертация фьючерсных балансов (маржи)
func convertFuturesBalances(userState *hyperliquid.UserState) []*entity.Balance {
	if userState == nil {
		return nil
	}

	balances := make([]*entity.Balance, 0, 1)

	// Основной маржинальный баланс USDC
	if userState.MarginSummary.AccountValue != "" {
		// Вычисляем заблокированные средства как разность между общим балансом и доступным для вывода
		accountValue, _ := strconv.ParseFloat(userState.MarginSummary.AccountValue, 64)
		withdrawable, _ := strconv.ParseFloat(userState.Withdrawable, 64)
		locked := accountValue - withdrawable

		balance := &entity.Balance{
			Asset:  "USDC",
			Free:   userState.Withdrawable,
			Locked: strconv.FormatFloat(locked, 'f', -1, 64),
			Total:  userState.MarginSummary.AccountValue,
		}
		balances = append(balances, balance)
	}

	return balances
}

// Конвертация спот-балансов всех активов
func convertBalances(spotState *hyperliquid.SpotUserState, userState *hyperliquid.UserState) []*entity.Balance {
	if spotState == nil {
		return nil
	}

	balances := make([]*entity.Balance, 0, len(spotState.Balances))

	// Проходим по всем спот-балансам
	for _, spotBalance := range spotState.Balances {
		// Вычисляем свободные средства (total - hold)
		total, _ := strconv.ParseFloat(spotBalance.Total, 64)
		hold, _ := strconv.ParseFloat(spotBalance.Hold, 64)
		free := total - hold

		balance := &entity.Balance{
			Asset:  spotBalance.Coin,
			Free:   strconv.FormatFloat(free, 'f', -1, 64),
			Locked: spotBalance.Hold,
			Total:  spotBalance.Total,
		}

		balances = append(balances, balance)
	}

	return balances
}

// Конвертация информации об инструментах
func convertInstrumentInfo(meta *hyperliquid.Meta) []*entity.Symbol {
	if meta == nil || meta.Universe == nil {
		return nil
	}

	symbols := make([]*entity.Symbol, 0, len(meta.Universe))

	for _, asset := range meta.Universe {
		if asset.Name == "" {
			continue
		}

		minQty, _ := strconv.ParseFloat(asset.SzDecimals, 64)
		if minQty == 0 {
			minQty = 0.001 // Значение по умолчанию
		}

		symbol := &entity.Symbol{
			Symbol:             asset.Name,
			Status:             entity.SymbolStatusTypeTrading,
			BaseAsset:          asset.Name,
			QuoteAsset:         "USDC",
			BaseAssetPrecision: 8,
			QuotePrecision:     8,
			OrderTypes: []entity.OrderType{
				entity.OrderTypeLimit,
				entity.OrderTypeMarket,
			},
			IcebergAllowed:         false,
			OcoAllowed:             false,
			IsSpotTradingAllowed:   false,
			IsMarginTradingAllowed: false,
			Filters: []entity.SymbolFilter{
				{
					FilterType: entity.SymbolFilterTypeLotSize,
					MinQty:     strconv.FormatFloat(minQty, 'f', -1, 64),
					MaxQty:     "1000000",
					StepSize:   strconv.FormatFloat(minQty, 'f', -1, 64),
				},
				{
					FilterType:  entity.SymbolFilterTypeMinNotional,
					MinNotional: "1",
				},
			},
			Permissions: []entity.AccountType{
				entity.AccountTypeSpot,
				entity.AccountTypeMargin,
			},
		}

		symbols = append(symbols, symbol)
	}

	return symbols
}

// Конвертация ордера
func convertOrder(order *hyperliquid.Order) *entity.Order {
	if order == nil {
		return nil
	}

	price, _ := strconv.ParseFloat(order.LimitPx, 64)
	qty, _ := strconv.ParseFloat(order.Sz, 64)
	executedQty, _ := strconv.ParseFloat(order.FilledSz, 64)

	return &entity.Order{
		Symbol:              order.Coin,
		OrderId:             strconv.FormatUint(order.Oid, 10),
		ClientOrderId:       order.Cloid,
		Price:               strconv.FormatFloat(price, 'f', -1, 64),
		OrigQty:             strconv.FormatFloat(qty, 'f', -1, 64),
		ExecutedQty:         strconv.FormatFloat(executedQty, 'f', -1, 64),
		CummulativeQuoteQty: "0",                        // Hyperliquid не предоставляет это значение
		Status:              convertOrderStatus("open"), // По умолчанию
		TimeInForce:         convertTimeInForce("GTC"),
		Type:                convertOrderType("limit"),
		Side:                convertSide(order.Side),
		Time:                order.Timestamp,
		UpdateTime:          order.Timestamp,
	}
}
