package hyperliquid

import (
	"strconv"

	"github.com/Betarost/onetrades/entity"
	hyperliquid "github.com/sonirico/go-hyperliquid"
)

// Константы для Hyperliquid API
const (
	// Стороны ордеров в Hyperliquid
	HyperliquidSideBuy  = "B"
	HyperliquidSideSell = "A"

	// Статусы ордеров
	HyperliquidStatusOpen     = "open"
	HyperliquidStatusFilled   = "filled"
	HyperliquidStatusCanceled = "canceled"

	// Типы ордеров
	HyperliquidOrderTypeLimit  = "limit"
	HyperliquidOrderTypeMarket = "market"

	// Валюта маржи по умолчанию
	DefaultMarginAsset = "USDC"
)

// convertSide конвертирует сторону заказа
func convertSide(side string) entity.SideType {
	switch side {
	case HyperliquidSideBuy:
		return entity.SideTypeBuy
	case HyperliquidSideSell:
		return entity.SideTypeSell
	default:
		return entity.SideTypeBuy
	}
}

// convertOrderStatus конвертирует статус заказа
func convertOrderStatus(status string) string {
	switch status {
	case HyperliquidStatusOpen:
		return "NEW"
	case HyperliquidStatusFilled:
		return "FILLED"
	case HyperliquidStatusCanceled:
		return "CANCELED"
	default:
		return "UNKNOWN"
	}
}

// convertOrderType конвертирует тип заказа
func convertOrderType(orderType string) entity.OrderType {
	switch orderType {
	case HyperliquidOrderTypeLimit:
		return entity.OrderTypeLimit
	case HyperliquidOrderTypeMarket:
		return entity.OrderTypeMarket
	default:
		return entity.OrderTypeLimit
	}
}

// convertTimeInForce конвертирует время действия заказа
func convertTimeInForce(tif string) string {
	switch tif {
	case "Gtc":
		return "GTC"
	case "Ioc":
		return "IOC"
	case "Fok":
		return "FOK"
	default:
		return "GTC"
	}
}

// convertPosition конвертирует позицию из Hyperliquid в общий формат
func convertPosition(pos *hyperliquid.Position) *entity.Futures_Positions {
	if pos == nil {
		return nil
	}

	// Конвертируем сторону позиции
	var positionSide string
	if pos.Szi != "" {
		if szi, err := strconv.ParseFloat(pos.Szi, 64); err == nil {
			switch {
			case szi > 0:
				positionSide = string(entity.PositionSideTypeLong)
			case szi < 0:
				positionSide = string(entity.PositionSideTypeShort)
			default:
				positionSide = string(entity.PositionSideTypeBoth)
			}
		} else {
			positionSide = string(entity.PositionSideTypeBoth)
		}
	}

	// Обрабатываем указатель на EntryPx
	entryPrice := ""
	if pos.EntryPx != nil {
		entryPrice = *pos.EntryPx
	}

	return &entity.Futures_Positions{
		Symbol:           pos.Coin,
		PositionSide:     positionSide,
		PositionSize:     pos.Szi,
		EntryPrice:       entryPrice,
		MarkPrice:        pos.PositionValue,
		UnRealizedProfit: pos.UnrealizedPnl,
		Notional:         pos.PositionValue,
		UpdateTime:       0, // Hyperliquid не предоставляет время обновления в Position
	}
}

// convertFuturesBalances конвертирует баланс фьючерсов из Hyperliquid в общий формат
func convertFuturesBalances(userState *hyperliquid.UserState) []entity.FuturesBalance {
	if userState == nil {
		return nil
	}

	balances := make([]entity.FuturesBalance, 0, 1)

	// Основной маржинальный баланс
	if userState.MarginSummary.AccountValue != "" {
		accountValue, err1 := strconv.ParseFloat(userState.MarginSummary.AccountValue, 64)
		withdrawable, err2 := strconv.ParseFloat(userState.Withdrawable, 64)

		// Если не удалось распарсить, используем значения по умолчанию
		if err1 != nil {
			accountValue = 0
		}
		if err2 != nil {
			withdrawable = 0
		}

		locked := accountValue - withdrawable
		if locked < 0 {
			locked = 0
		}

		balance := entity.FuturesBalance{
			Asset:            DefaultMarginAsset, // Стандартная валюта маржи в Hyperliquid
			Balance:          userState.MarginSummary.AccountValue,
			Equity:           userState.MarginSummary.AccountValue,
			Available:        userState.Withdrawable,
			UnrealizedProfit: strconv.FormatFloat(locked, 'f', -1, 64),
		}
		balances = append(balances, balance)
	}
	return balances
}

// convertSpotBalances конвертирует спот-балансы из Hyperliquid в общий формат
func convertSpotBalances(spotUserState *hyperliquid.SpotUserState) []entity.AssetsBalance {
	if spotUserState == nil {
		return nil
	}

	balances := make([]entity.AssetsBalance, 0, len(spotUserState.Balances))

	for _, balance := range spotUserState.Balances {
		total, err1 := strconv.ParseFloat(balance.Total, 64)
		hold, err2 := strconv.ParseFloat(balance.Hold, 64)

		// Пропускаем балансы с некорректными данными
		if err1 != nil || err2 != nil {
			continue
		}

		free := total - hold
		if free < 0 {
			free = 0
		}

		assetBalance := entity.AssetsBalance{
			Asset:   balance.Coin, // Используем реальное название валюты из API
			Balance: strconv.FormatFloat(free, 'f', -1, 64),
			Locked:  balance.Hold,
		}
		balances = append(balances, assetBalance)
	}

	return balances
}

// convertInstrumentInfo конвертирует информацию об инструменте
func convertInstrumentInfo(assetInfo *hyperliquid.AssetInfo) *entity.Futures_InstrumentsInfo {
	if assetInfo == nil {
		return nil
	}

	return &entity.Futures_InstrumentsInfo{
		Symbol:         assetInfo.Name,
		Base:           assetInfo.Name,
		Quote:          DefaultMarginAsset,
		MinQty:         "0.001", // Значение по умолчанию
		PricePrecision: strconv.Itoa(assetInfo.SzDecimals),
		SizePrecision:  strconv.Itoa(assetInfo.SzDecimals),
		State:          "TRADING",
		MaxLeverage:    "100", // Значение по умолчанию, AssetInfo не содержит MaxLeverage
	}
}

// convertOrder конвертирует ордер из Hyperliquid в общий формат
func convertOrder(order *hyperliquid.OpenOrder) *entity.Futures_OrdersList {
	if order == nil {
		return nil
	}

	// OpenOrder не имеет поля Cloid, используем пустую строку
	clientOrderID := ""

	return &entity.Futures_OrdersList{
		Symbol:        order.Coin,
		OrderID:       strconv.FormatInt(order.Oid, 10),
		ClientOrderID: clientOrderID,
		Side:          string(convertSide(string(order.Side))),
		PositionSize:  strconv.FormatFloat(order.Size, 'f', -1, 64), // Size вместо Sz
		Price:         strconv.FormatFloat(order.LimitPx, 'f', -1, 64),
		Type:          "limit",                    // OpenOrder всегда лимитные
		Status:        convertOrderStatus("open"), // Все открытые ордера имеют статус "open"
		CreateTime:    order.Timestamp,
		UpdateTime:    order.Timestamp,
	}
}
