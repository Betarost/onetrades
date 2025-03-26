package onetrades

import (
	"context"
	"log"
	"os"
	"os/signal"
	"testing"

	"github.com/spf13/viper"
)

func TestOnetrades(t *testing.T) {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln("Error ReadInConfig", err)
	}

	//==========================BINANCE==========================
	// binanceKey := viper.GetString("BINANCE_API")
	// binanceSecret := viper.GetString("BINANCE_SECRET")
	// client := NewFutureBinanceClient(binanceKey, binanceSecret)
	//======================= GET Balance
	// res, err := client.NewGetAccountBalance().Do(context.Background())
	// t.Logf("Results: %+v  %v", res, err)
	//======================= GET Position
	// res, err := client.NewGetPositions().Do(context.Background())
	// t.Logf("Results: %+v  %v", res, err)
	// ======================= GET TradeHistoryOrder
	// res, err := client.NewGetTradeHistoryOrder().Symbol("BTCUSDT").Limit(1000).Do(context.Background())
	// t.Logf("Results: %+v  %v", res, err)
	// ======================= GET GetDownloadIdHistoryOrder GetDownloadLinkHistoryOrder
	// res, err := client.NewGetDownloadIdHistoryOrder().Begin(1741109298000).End(1742282098000).Do(context.Background())
	// t.Logf("Results: %+v  %v", res, err)
	// res, err := client.NewGetDownloadLinkHistoryOrder().DownloadId("954279024084131840").Do(context.Background())
	// t.Logf("Results: %+v  %v", res, err)
	//======================END BINANCE==========================

	//==========================BYBIT==========================
	// bybitKey := viper.GetString("BYBIT_API")
	// bybitSecret := viper.GetString("BYBIT_SECRET")
	// client := NewFutureBybitClient(bybitKey, bybitSecret)
	//======================= GET Balance
	// res, err := client.NewGetAccountBalance().Do(context.Background())
	// t.Logf("Results: %+v  %v", res, err)
	//======================= GET Position
	// res, err := client.NewGetPositions().Do(context.Background())
	// t.Logf("Results: %+v  %v", res, err)
	// ======================= GET TradeHistoryOrder
	// res, err := client.NewGetTradeHistoryOrder().Do(context.Background())
	// t.Logf("Results: %+v  %v", res, err)
	//======================END BYBIT==========================

	//==========================MEXC==========================
	// mexcKey := viper.GetString("MEX_API")
	// mexcSecret := viper.GetString("MEX_SECRET")
	// mexcMemo := viper.GetString("MEX_MEMO")
	// client := NewFutureMexcClient(mexcKey, mexcSecret, mexcMemo)
	//======================= GET Balance
	// res, err := client.NewGetAccountBalance().Do(context.Background())
	// t.Logf("Results: %+v  %v", res, err)
	//======================= GET Position
	// res, err := client.NewGetPositions().Do(context.Background())
	// t.Logf("Results: %+v  %v", res, err)
	//======================= GET ContractInfo
	// res, err := client.NewGetContractInfo().Symbol("DOGE_USDT").Do(context.Background())
	// t.Logf("Results: %+v  %v", res, err)
	//======================= GET FairPrice
	// res, err := client.NewGetFairPrice().Symbol("DOGE_USDT").Do(context.Background())
	// t.Logf("Results: %+v  %v", res, err)
	//======================END MEXC==========================

	//==========================BINGX==========================
	// bingxKey := viper.GetString("BINGX_API")
	// bingxSecret := viper.GetString("BINGX_SECRET")
	// client := NewFutureBingxClient(bingxKey, bingxSecret)
	//======================= GET Balance
	// res, err := client.NewGetAccountBalance().Do(context.Background())
	// t.Logf("Results: %+v  %v", res, err)
	//======================= GET Position
	// res, err := client.NewGetPositions().Do(context.Background())
	// t.Logf("Results: %+v  %v", res, err)
	//======================= GET OrdersHistory
	// startime := time.Now().Add(time.Hour * (-5)).UnixMilli()
	// res, err := client.NewGetOrdersHistory().StartTime(startime).Do(context.Background())
	// t.Logf("Results: %+v  %v", res, err)
	//======================END BINGX==========================

	//==========================OKX==========================
	okxKey := viper.GetString("OKX_API")
	okxSecret := viper.GetString("OKX_SECRET")
	okxMemo := viper.GetString("OKX_MEMO")
	client := NewFutureOKXClient(okxKey, okxSecret, okxMemo)
	//======================= GET FundsTransfer
	res, err := client.NewFundsTransfer().Way("2").Amount("0.05").From("6").To("6").SubID("Betarost").Do(context.Background())
	t.Logf("Results: %+v  %v", res, err)
	//======================= GET AccountValuation
	// res, err := client.NewGetAccountValuation().Do(context.Background())
	// t.Logf("Results: %+v  %v", res, err)
	//======================= GET AccountInfo
	// res, err := client.NewGetAccountInfo().Do(context.Background())
	// t.Logf("Results: %+v  %v", res, err)
	//======================= GET SubAccountsLists
	// res, err := client.NewGetSubAccountsLists().Do(context.Background())
	// t.Logf("Results: %+v  %v", res, err)
	//======================= GET SubAccountBalance
	// res, err := client.NewGetSubAccountBalance().SubID("Betarost").Do(context.Background())
	// t.Logf("Results: %+v  %v", res, err)
	//======================= GET SubAccountFundingBalance
	// res, err := client.NewGetSubAccountFundingBalance().SubID("Betarost").Do(context.Background())
	// t.Logf("Results: %+v  %v", res, err)
	//======================= GET Balance
	// res, err := client.NewGetAccountBalance().Do(context.Background())
	// t.Logf("Results: %+v  %v", res, err)
	//======================= GET Position
	// res, err := client.NewGetPositions().Do(context.Background())
	// t.Logf("Results: %+v  %v", res, err)
	//======================= GET HistoryPositions
	// res, err := client.NewGetHistoryPositions().Do(context.Background())
	// t.Logf("Results: %+v  %v", res, err)
	// ======================= GET TradeHistoryOrder
	// res, err := client.NewGetTradeHistoryOrder().Do(context.Background())
	// t.Logf("Results: %+v  %v", res, err)
	//======================= SET AccountMode
	// res, err := client.NewSetAccountMode().Mode("2").Do(context.Background())
	// t.Logf("Results: %+v  %v", res, err)
	//======================= SET PositionMode
	// res, err := client.NewSetPositionMode().Mode(entity.PositionModeTypeHedge).Do(context.Background())
	// t.Logf("Results: %+v  %v", res, err)
	//======================= SET Leverage
	// res, err := client.NewSetLeverage().Symbol("DOGE-USDT-SWAP").Leverage(50).Do(context.Background())
	// t.Logf("Results: %+v  %v", res, err)
	//======================= GET ContractsInfo
	// res, err := client.NewGetContractsInfo().Do(context.Background())
	// t.Logf("Results: %+v  %v", res, err)
	//======================= GET ContractsInfo One Symbol
	// res, err := client.NewGetContractsInfo().Symbol("DOGE-USDT-SWAP").Do(context.Background())
	// t.Logf("Results: %+v  %v", res, err)
	// ======================= GET MarkPrices
	// res, err := client.NewGetMarkPrices().Do(context.Background())
	// t.Logf("Results: %+v  %v", res, err)
	// ======================= GET MarkPrice
	// res, err := client.NewGetMarkPrice().Symbol("DOGE-USDT-SWAP").Do(context.Background())
	// t.Logf("Results: %+v  %v", res, err)
	//======================= GET Kline
	// res, err := client.NewGetKline().Symbol("DOGE-USDT-SWAP").TimeFrame(entity.TimeFrameType5m).Limit(13).Do(context.Background())
	// t.Logf("Results: %+v  %v", res, err)
	//=======================Trade PlaceOrder
	// res, err := client.NewTradePlaceOrder().Symbol("DOGE-USDT-SWAP").PositionSide(entity.PositionSideTypeLong).Side(entity.SideTypeBuy).Size("0.1").Price("0.19876").OrderType(entity.OrderTypeLimit).Do(context.Background())
	// t.Logf("Results: %+v  %v", res, err)
	//=======================Get OrderList
	// res, err := client.NewGetOrderList().Symbol("DOGE-USDT-SWAP").OrderType(entity.OrderTypeLimit).Do(context.Background())
	// res, err := client.NewGetOrderList().OrderType(entity.OrderTypeLimit).Do(context.Background())
	// t.Logf("Results: %+v  %v", res, err)
	//=======================Get CancelOrders
	// res, err := client.NewTradeCancelOrders().Symbol("DOGE-USDT-SWAP").OrderIDs([]string{"2284238701511041024", "2284179927031078912"}).Do(context.Background())
	// t.Logf("Results: %+v  %v", res, err)
	// //=======================WebSocket Public
	// ws := client.NewWebSocketPublicClient()
	// //=======================MarkPrice
	// wsPublicMarkPriceHandler := func(event *entity.WsPublicMarkPriceEvent) {
	// 	log.Printf("=wsPublicMarkPriceHandler= %+v", event)
	// }
	// errHandler := func(err error) {
	// 	log.Printf("wsPublicMarkPriceHandler Error: %s", err.Error())
	// }
	// ws.NewPublicMarkPrice([]string{"DOGE-USDT-SWAP", "TON-USDT-SWAP"}, wsPublicMarkPriceHandler, errHandler)
	//=======================WebSocket Private
	// ws := client.NewWebSocketPrivateClient()
	//=======================Orders
	// wsPrivateOrdersHandler := func(event *entity.WsPrivateOrdersEvent) {
	// 	log.Printf("=wsPrivateOrdersHandler= %+v", event)
	// }
	// errHandler := func(err error) {
	// 	log.Printf("wsPrivateOrdersHandler Error: %s", err.Error())
	// }
	// time.Sleep(1 * time.Second)
	// ws.NewPrivateOrders(wsPrivateOrdersHandler, errHandler)
	//======================= PlaceOrders
	// time.Sleep(1 * time.Second)
	// data := []map[string]string{
	// 	{
	// 		"instId":  "DOGE-USDT-SWAP",
	// 		"tdMode":  "cross",
	// 		"clOrdId": "BLONG",
	// 		"side":    "buy",
	// 		"posSide": "long",
	// 		"ordType": "limit",
	// 		"sz":      "0.01",
	// 		"px":      "0.20123",
	// 	},
	// 	{
	// 		"instId":  "DOGE-USDT-SWAP",
	// 		"tdMode":  "cross",
	// 		"clOrdId": "SSHORT",
	// 		"side":    "sell",
	// 		"posSide": "short",
	// 		"ordType": "limit",
	// 		"sz":      "0.01",
	// 		"px":      "0.21123",
	// 	},
	// }
	// ws.NewPrivatePlaceOrders(data)
	//======================= CancelOrders
	// time.Sleep(1 * time.Second)
	// data := []map[string]string{
	// 	{
	// 		"instId":  "DOGE-USDT-SWAP",
	// 		"clOrdId": "BLONG",
	// 	},
	// 	{
	// 		"instId":  "DOGE-USDT-SWAP",
	// 		"clOrdId": "SSHORT",
	// 	},
	// }
	// ws.NewPrivateCancelOrders(data)
	//=======================Positions
	// wsPrivatePositionsHandler := func(event *entity.WsPrivatePositionsEvent) {
	// 	log.Printf("=wsPrivatePositionsHandler= %+v", event)
	// }
	// errHandler := func(err error) {
	// 	log.Printf("wsPrivatePositionsHandler Error: %s", err.Error())
	// }
	// time.Sleep(1 * time.Second)
	// ws.NewPrivatePositions(wsPrivatePositionsHandler, errHandler)
	//======================END OKX==========================

	//=====================GATE GET BALANCE======================-=====
	// gateKey := viper.GetString("GATE_API")
	// gateSecret := viper.GetString("GATE_SECRET")
	// client := NewFutureGateClient(gateKey, gateSecret)
	// res, err := client.NewGetAccountBalance().Do(context.Background())
	// t.Logf("Results: %+v  %v", res, err)

	//=====================BITGET GET BALANCE======================-=====
	// bitgetKey := viper.GetString("BITGET_API")
	// bitgetSecret := viper.GetString("BITGET_SECRET")
	// bitgetMemo := viper.GetString("BITGET_MEMO")
	// client := NewFutureBitgetClient(bitgetKey, bitgetSecret, bitgetMemo)
	// client.Debug = true
	// res, err := client.NewGetAccountBalance().Do(context.Background())
	// t.Logf("Results: %+v  %v", res, err)

	//=====================HUOBI GET BALANCE======================-=====
	// bingxKey := viper.GetString("BINGX_API")
	// bingxSecret := viper.GetString("BINGX_SECRET")
	// client := NewFutureBingxClient(bingxKey, bingxSecret)
	// res, err := client.NewGetAccountBalance().Do(context.Background())
	// t.Logf("Results: %+v  %v", res, err)

	//===========Not Exit
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	for {
		<-c
		return
	}

}
