package hyperliquid

import (
	"crypto/ecdsa"

	"github.com/sonirico/go-hyperliquid"
)

// FuturesClient для фьючерсной торговли на Hyperliquid
type FuturesClient struct {
	client     *hyperliquid.Client
	exchange   *hyperliquid.Exchange
	info       *hyperliquid.Info
	privateKey *ecdsa.PrivateKey
	isTestnet  bool
}

// NewFuturesClient создает новый клиент для фьючерсной торговли
func NewFuturesClient(privateKeyHex string, isTestnet bool) (*FuturesClient, error) {
	var baseURL string
	if isTestnet {
		baseURL = hyperliquid.TestnetAPIURL
	} else {
		baseURL = hyperliquid.MainnetAPIURL
	}

	client := hyperliquid.NewClient(baseURL)

	// Создаем Exchange клиент с приватным ключом
	exchange, err := hyperliquid.NewExchange(client, privateKeyHex)
	if err != nil {
		return nil, err
	}

	// Создаем Info клиент
	info := hyperliquid.NewInfo(client)

	return &FuturesClient{
		client:    client,
		exchange:  exchange,
		info:      info,
		isTestnet: isTestnet,
	}, nil
}

// Методы для фьючерсной торговли

// NewGetInstrumentsInfo получает информацию об инструментах
func (c *FuturesClient) NewGetInstrumentsInfo() *futures_getInstrumentsInfo {
	return &futures_getInstrumentsInfo{
		client: c,
	}
}

// NewGetBalance получает баланс аккаунта
func (c *FuturesClient) NewGetBalance() *futures_getBalance {
	return &futures_getBalance{
		client: c,
	}
}

// NewGetPositions получает позиции
func (c *FuturesClient) NewGetPositions() *futures_getPositions {
	return &futures_getPositions{
		client: c,
	}
}

// NewPlaceOrder размещает ордер
func (c *FuturesClient) NewPlaceOrder() *futures_placeOrder {
	return &futures_placeOrder{
		client: c,
	}
}

// NewCancelOrder отменяет ордер
func (c *FuturesClient) NewCancelOrder() *futures_cancelOrder {
	return &futures_cancelOrder{
		client: c,
	}
}

// NewGetOrderList получает список ордеров
func (c *FuturesClient) NewGetOrderList() *futures_getOrderList {
	return &futures_getOrderList{
		client: c,
	}
}

// NewOrdersHistory получает историю ордеров
func (c *FuturesClient) NewOrdersHistory() *futures_ordersHistory {
	return &futures_ordersHistory{
		client: c,
	}
}

// NewAmendOrder изменяет ордер
func (c *FuturesClient) NewAmendOrder() *futures_amendOrder {
	return &futures_amendOrder{
		client: c,
	}
}

// NewGetLeverage получает кредитное плечо
func (c *FuturesClient) NewGetLeverage() *futures_getLeverage {
	return &futures_getLeverage{
		client: c,
	}
}

// NewSetLeverage устанавливает кредитное плечо
func (c *FuturesClient) NewSetLeverage() *futures_setLeverage {
	return &futures_setLeverage{
		client: c,
	}
}

// NewGetMarginMode получает режим маржи
func (c *FuturesClient) NewGetMarginMode() *futures_getMarginMode {
	return &futures_getMarginMode{
		client: c,
	}
}

// NewSetMarginMode устанавливает режим маржи
func (c *FuturesClient) NewSetMarginMode() *futures_setMarginMode {
	return &futures_setMarginMode{
		client: c,
	}
}

// NewGetPositionMode получает режим позиций
func (c *FuturesClient) NewGetPositionMode() *futures_getPositionMode {
	return &futures_getPositionMode{
		client: c,
	}
}

// NewSetPositionMode устанавливает режим позиций
func (c *FuturesClient) NewSetPositionMode() *futures_setPositionMode {
	return &futures_setPositionMode{
		client: c,
	}
}

// NewPositionsHistory получает историю позиций
func (c *FuturesClient) NewPositionsHistory() *futures_positionsHistory {
	return &futures_positionsHistory{
		client: c,
	}
}
