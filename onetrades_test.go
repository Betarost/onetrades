package onetrades

import (
	"context"
	"log"
	"testing"

	"github.com/Betarost/onetrades/binance"
	"github.com/Betarost/onetrades/bingx"
	"github.com/Betarost/onetrades/bitget"
	"github.com/Betarost/onetrades/bullish"
	"github.com/Betarost/onetrades/bybit"
	"github.com/Betarost/onetrades/gateio"
	"github.com/Betarost/onetrades/huobi"
	"github.com/Betarost/onetrades/mexc"
	"github.com/Betarost/onetrades/okx"
	"github.com/spf13/viper"
)

var n = ""

func printAnswers(r interface{}, e error) {
	log.Printf("=%s= %+v %v", n, r, e)
}
func TestOnetrades(t *testing.T) {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln("Error ReadInConfig", err)
	}
	ctx := context.Background()
	//==========================KEYS==========================
	binanceKey := viper.GetString("BINANCE_API")
	binanceSecret := viper.GetString("BINANCE_SECRET")
	bingxKey := viper.GetString("BINGX_API")
	bingxSecret := viper.GetString("BINGX_SECRET")
	bybitKey := viper.GetString("BYBIT_API")
	bybitSecret := viper.GetString("BYBIT_SECRET")
	gateKey := viper.GetString("GATE_API")
	gateSecret := viper.GetString("GATE_SECRET")
	mexcKey := viper.GetString("MEX_API")
	mexcSecret := viper.GetString("MEX_SECRET")
	bitgetKey := viper.GetString("BITGET_API")
	bitgetSecret := viper.GetString("BITGET_SECRET")
	bitgetMemo := viper.GetString("BITGET_MEMO")
	okxKey := viper.GetString("OKX_API")
	okxSecret := viper.GetString("OKX_SECRET")
	okxMemo := viper.GetString("OKX_MEMO")
	huobiKey := viper.GetString("HUOBI_API")
	huobiSecret := viper.GetString("HUOBI_SECRET")
	bullishKey := viper.GetString("BULLISH_API")
	bullishSecret := viper.GetString("BULLISH_SECRET")
	bullishJWT := viper.GetString("BULLISH_JWT")
	//==========================CLIENTS==========================
	binanceSpot := binance.NewSpotClient(binanceKey, binanceSecret)
	binanceFutures := binance.NewFuturesClient(binanceKey, binanceSecret)
	bingxSpot := bingx.NewSpotClient(bingxKey, bingxSecret)
	bingxFutures := bingx.NewFuturesClient(bingxKey, bingxSecret)
	bybitSpot := bybit.NewSpotClient(bybitKey, bybitSecret)
	bybitFutures := bybit.NewFuturesClient(bybitKey, bybitSecret)
	gateioSpot := gateio.NewSpotClient(gateKey, gateSecret)
	gateioFutures := gateio.NewFuturesClient(gateKey, gateSecret)
	mexcSpot := mexc.NewSpotClient(mexcKey, mexcSecret)
	// mexcFutures := mexc.NewFuturesClient(mexcKey, mexcSecret)
	bitgetSpot := bitget.NewSpotClient(bitgetKey, bitgetSecret, bitgetMemo)
	bitgetFutures := bitget.NewFuturesClient(bitgetKey, bitgetSecret, bitgetMemo)
	okxSpot := okx.NewSpotClient(okxKey, okxSecret, okxMemo)
	okxFutures := okx.NewFuturesClient(okxKey, okxSecret, okxMemo)
	huobiSpot := huobi.NewSpotClient(huobiKey, huobiSecret)
	huobiFutures := huobi.NewFuturesClient(huobiKey, huobiSecret)
	bullishFutures := bullish.NewFuturesClient(bullishKey, bullishSecret, bullishJWT)

	binanceSpot.Debug = false
	binanceFutures.Debug = false
	bingxSpot.Debug = false
	bingxFutures.Debug = false
	bybitSpot.Debug = false
	bybitFutures.Debug = false
	gateioSpot.Debug = false
	gateioFutures.Debug = false
	mexcSpot.Debug = false
	// mexcFutures.Debug = false
	bitgetSpot.Debug = false
	bitgetFutures.Debug = false
	okxSpot.Debug = false
	okxFutures.Debug = false
	huobiSpot.Debug = false
	huobiFutures.Debug = false
	bullishFutures.Debug = false
	bullishFutures.Proxy = "http://localhost:1080"
	//=======================AccountInfo
	n = "AccountInfo"
	//SPOT
	// printAnswers(binanceSpot.NewGetAccountInfo().Do(ctx))
	// printAnswers(bingxSpot.NewGetAccountInfo().Do(ctx))
	// printAnswers(bybitSpot.NewGetAccountInfo().Do(ctx))
	// printAnswers(gateioSpot.NewGetAccountInfo().Do(ctx))
	// printAnswers(mexcSpot.NewGetAccountInfo().Do(ctx))
	// printAnswers(bitgetSpot.NewGetAccountInfo().Do(ctx))
	// printAnswers(okxSpot.NewGetAccountInfo().Do(ctx))
	// printAnswers(huobiSpot.NewGetAccountInfo().Do(ctx))

	//FUTURES
	// printAnswers(bingxFutures.NewGetAccountInfo().Do(ctx))
	// printAnswers(bybitFutures.NewGetAccountInfo().Do(ctx))
	// printAnswers(gateioFutures.NewGetAccountInfo().Do(ctx))
	// printAnswers(bitgetFutures.NewGetAccountInfo().Do(ctx))
	// printAnswers(okxFutures.NewGetAccountInfo().Do(ctx))
	// printAnswers(huobiFutures.NewGetAccountInfo().Do(ctx))

	// printAnswers(binanceFutures.NewGetAccountInfo().Do(ctx))

	//=======================GetBalance
	n = "GetBalance"
	//SPOT
	// printAnswers(binanceSpot.NewGetBalance().Do(ctx))
	// printAnswers(bingxSpot.NewGetBalance().Do(ctx))
	// printAnswers(bybitSpot.NewGetBalance().Do(ctx))
	// printAnswers(gateioSpot.NewGetBalance().Do(ctx))
	// printAnswers(mexcSpot.NewGetBalance().Do(ctx))
	// printAnswers(bitgetSpot.NewGetBalance().Do(ctx))
	// printAnswers(okxSpot.NewGetBalance().Do(ctx))
	// printAnswers(huobiSpot.NewGetBalance().UID("53799773").Do(ctx))
	// printAnswers(huobiSpot.NewGetBalance().UID("69069265").Do(ctx))

	//FUTURES
	// printAnswers(binanceFutures.NewGetBalance().Do(ctx))
	// printAnswers(bingxFutures.NewGetBalance().Do(ctx))
	// printAnswers(bybitFutures.NewGetBalance().Do(ctx))
	// printAnswers(gateioFutures.NewGetBalance().Do(ctx))
	// printAnswers(bitgetFutures.NewGetBalance().Do(ctx))
	// printAnswers(okxFutures.NewGetBalance().Do(ctx))
	// printAnswers(huobiFutures.NewGetBalance().UID("53799773").Do(ctx))
	printAnswers(bullishFutures.NewGetBalance().UID("111872616831896").Do(ctx))

	//=======================InstrumentsInfo
	n = "InstrumentsInfo"
	//SPOT
	// printAnswers(binanceSpot.NewGetInstrumentsInfo().Symbol("BTCUSDT").Do(ctx))
	// printAnswers(bybitSpot.NewGetInstrumentsInfo().Symbol("DOGEUSDT").Do(ctx))
	// printAnswers(bingxSpot.NewGetInstrumentsInfo().Symbol("TRX-USDT").Do(ctx))
	// printAnswers(gateioSpot.NewGetInstrumentsInfo().Symbol("DOGE_USDT").Do(ctx))
	// printAnswers(mexcSpot.NewGetInstrumentsInfo().Symbol("BTCUSDT").Do(ctx))
	// printAnswers(bitgetSpot.NewGetInstrumentsInfo().Symbol("DOGEUSDT").Do(ctx))
	// printAnswers(okxSpot.NewGetInstrumentsInfo().Symbol("DOGE-USDT").Do(ctx))
	// printAnswers(huobiSpot.NewGetInstrumentsInfo().Do(ctx))
	// huobiSpot.NewGetInstrumentsInfo().Do(ctx)
	//FUTURES
	// printAnswers(binanceFutures.NewGetInstrumentsInfo().Do(ctx))
	// printAnswers(bybitFutures.NewGetInstrumentsInfo().Symbol("DOGEUSDT").Do(ctx))
	// printAnswers(bingxFutures.NewGetInstrumentsInfo().Symbol("TRX-USDT").Do(ctx))
	// printAnswers(gateioFutures.NewGetInstrumentsInfo().Symbol("DOGE_USDT").Do(ctx))
	// printAnswers(bitgetFutures.NewGetInstrumentsInfo().Symbol("DOGEUSDT").Do(ctx))
	// printAnswers(okxFutures.NewGetInstrumentsInfo().Symbol("ETH-USD-SWAP").Do(ctx))
	// printAnswers(huobiFutures.NewGetInstrumentsInfo().Symbol("BTC-USDT").Do(ctx))
	// printAnswers(bullishFutures.NewGetInstrumentsInfo().Symbol("BTC-USDC-PERP").Do(ctx))

	//=======================PlaceOrder
	n = "PlaceOrder"
	//SPOT
	// printAnswers(binanceSpot.NewPlaceOrder().Symbol("DOGEUSDT").Side(entity.SideTypeSell).OrderType(entity.OrderTypeLimit).Price("0.25050").Size("10").Do(ctx))
	// printAnswers(bingxSpot.NewPlaceOrder().Symbol("TRX-USDT").Side(entity.SideTypeSell).OrderType(entity.OrderTypeMarket).Price("0.2892").Size("10").Do(ctx))
	// printAnswers(bybitSpot.NewPlaceOrder().Symbol("DOGEUSDT").Side(entity.SideTypeSell).OrderType(entity.OrderTypeLimit).Price("0.19874").Size("30").Do(ctx))
	// printAnswers(gateioSpot.NewPlaceOrder().Symbol("DOGE_USDT").Side(entity.SideTypeSell).OrderType(entity.OrderTypeLimit).Size("19").Price("0.18250").Do(ctx))
	// printAnswers(mexcSpot.NewPlaceOrder().Symbol("DOGEUSDT").Side(entity.SideTypeSell).OrderType(entity.OrderTypeLimit).Price("0.17288").Size("10").Do(ctx))
	// printAnswers(bitgetSpot.NewPlaceOrder().Symbol("DOGEUSDT").Side(entity.SideTypeSell).OrderType(entity.OrderTypeMarket).Price("0.17250").Size("11").Do(ctx))
	// printAnswers(okxSpot.NewPlaceOrder().Symbol("DOGE-USDT").Side(entity.SideTypeSell).OrderType(entity.OrderTypeMarket).Price("0.17111").Size("10").Do(ctx))
	// printAnswers(huobiSpot.NewPlaceOrder().Symbol("DOGEUSDT").Side(entity.SideTypeSell).OrderType(entity.OrderTypeLimit).Size("56").Price("0.20899").UID("69069265").Do(ctx))

	//FUTURES
	// printAnswers(binanceFutures.NewPlaceOrder().Symbol("DOGEUSDT").Side(entity.SideTypeSell).OrderType(entity.OrderTypeMarket).Size("35").Do(ctx))
	// printAnswers(bingxFutures.NewPlaceOrder().Symbol("TRX-USDT").Side(entity.SideTypeBuy).OrderType(entity.OrderTypeMarket).Price("0.28550").Size("10").HedgeMode(false).Do(ctx))
	// printAnswers(bybitFutures.NewPlaceOrder().Symbol("DOGEUSDT").Side(entity.SideTypeSell).OrderType(entity.OrderTypeLimit).Price("0.19880").Size("35").Do(ctx))
	// printAnswers(gateioFutures.NewPlaceOrder().Symbol("DOGE_USDT").Side(entity.SideTypeSell).Size("2").Price("0.18780").OrderType(entity.OrderTypeLimit).Do(ctx))
	// printAnswers(bitgetFutures.NewPlaceOrder().Symbol("DOGEUSDT").Side(entity.SideTypeSell).OrderType(entity.OrderTypeMarket).Price("0.17000").Size("120").HedgeMode(false).MarginMode(entity.MarginModeTypeCross).Do(ctx))
	// printAnswers(okxFutures.NewPlaceOrder().Symbol("DOGE-USDT-SWAP").Side(entity.SideTypeBuy).OrderType(entity.OrderTypeLimit).Price("0.17950").Size("0.01").MarginMode(entity.MarginModeTypeCross).Do(ctx))
	// printAnswers(bullishFutures.NewPlaceOrder().UID("111872616831896").Symbol("BTC-USDC-PERP").Side(entity.SideTypeBuy).OrderType(entity.OrderTypeMarket).Size("0.001").Do(ctx))
	// printAnswers(bullishFutures.NewPlaceOrder().UID("111872616831896").Symbol("BTC-USDC-PERP").Side(entity.SideTypeSell).OrderType(entity.OrderTypeLimit).Price("118750").Size("0.001").Do(ctx))

	//=======================GetPositions
	n = "GetPositions"
	//FUTURES
	// printAnswers(binanceFutures.NewGetPositions().Do(ctx))
	// printAnswers(bingxFutures.NewGetPositions().Do(ctx))
	// printAnswers(bybitFutures.NewGetPositions().Symbol("BTCUSDT").Do(ctx))
	// printAnswers(gateioFutures.NewGetPositions().Do(ctx))
	// printAnswers(bitgetFutures.NewGetPositions().Do(ctx))
	// printAnswers(okxFutures.NewGetPositions().Do(ctx))
	printAnswers(bullishFutures.NewGetPositions().UID("111872616831896").Do(ctx))

	//=======================CancelOrder
	n = "CancelOrder"
	//SPOT
	// printAnswers(binanceSpot.NewCancelOrder().Symbol("DOGEUSDT").OrderID("10968968313").Do(ctx))
	// printAnswers(bingxSpot.NewCancelOrder().Symbol("TRX-USDT").OrderID("1942168112220504064").Do(ctx))
	// printAnswers(bybitSpot.NewCancelOrder().OrderID("1993102065488195328").Do(ctx))
	// printAnswers(gateioSpot.NewCancelOrder().Symbol("DOGE_USDT").OrderID("872370092892").Do(ctx))
	// printAnswers(mexcSpot.NewCancelOrder().Symbol("DOGEUSDT").OrderID("C02__571234536180977664120").Do(ctx))
	// printAnswers(bitgetSpot.NewCancelOrder().Symbol("DOGEUSDT").OrderID("1326467275447353344").Do(ctx))
	// printAnswers(okxSpot.NewCancelOrder().Symbol("DOGE-USDT").OrderID("2669952276160045056").Do(ctx))
	// printAnswers(huobiSpot.NewCancelOrder().OrderID("1371171052466066").Do(ctx))

	//FUTURES
	// printAnswers(binanceFutures.NewCancelOrder().Symbol("DOGEUSDT").OrderID("76780426418").Do(ctx))
	// printAnswers(bingxFutures.NewCancelOrder().Symbol("TRX-USDT").OrderID("1942486320061566976").Do(ctx))
	// printAnswers(bybitFutures.NewCancelOrder().Symbol("DOGEUSDT").OrderID("0c88fc3c-7896-4b62-aacc-b5c9c52122cd").Do(ctx))
	// printAnswers(gateioFutures.NewCancelOrder().OrderID("56013522187409589").Do(ctx))
	// printAnswers(bitgetFutures.NewCancelOrder().Symbol("DOGEUSDT").OrderID("1326805187463618574").Do(ctx))
	// printAnswers(okxFutures.NewCancelOrder().Symbol("DOGE-USDT-SWAP").OrderID("2672372971099906048").Do(ctx))
	// printAnswers(bullishFutures.NewCancelOrder().UID("111872616831896").Symbol("BTC-USDC-PERP").OrderID("869492653091717121").Do(ctx))

	//=======================AmendOrder
	n = "AmendOrder"
	//SPOT
	// printAnswers(binanceSpot.NewAmendOrder().Symbol("DOGEUSDT").OrderID("10968968313").NewSize("9").Do(ctx))
	// printAnswers(bybitSpot.NewAmendOrder().Symbol("DOGEUSDT").OrderID("1993102065488195328").NewPrice("0.19855").NewSize("31").Do(ctx))
	// printAnswers(gateioSpot.NewAmendOrder().Symbol("TRX_USDT").OrderID("186336434877910305").NewSize("3").NewPrice("0.2314").Do(ctx))
	// printAnswers(okxSpot.NewAmendOrder().Symbol("DOGE-USDT").OrderID("2669952276160045056").NewPrice("0.17101").NewSize("11.11").Do(ctx))

	//FUTURES
	// printAnswers(binanceFutures.NewAmendOrder().Symbol("DOGEUSDT").OrderID("76780426418").NewSize("39").NewPrice("0.23522").Side(entity.SideTypeBuy).Do(ctx))
	// printAnswers(gateioFutures.NewAmendOrder().OrderID("186336434877910305").NewSize("3").NewPrice("0.2314").Do(ctx))
	// printAnswers(bybitFutures.NewAmendOrder().Symbol("DOGEUSDT").OrderID("0c88fc3c-7896-4b62-aacc-b5c9c52122cd").NewPrice("0.19897").NewSize("31").Do(ctx))
	// printAnswers(bitgetFutures.NewAmendOrder().Symbol("DOGEUSDT").OrderID("1326804028506120199").NewSize("35").NewPrice("0.17011").NewСlientOrderID("12345").Do(ctx))
	// printAnswers(okxFutures.NewAmendOrder().Symbol("DOGE-USDT-SWAP").OrderID("2672372971099906048").NewSize("0.02").NewPrice("0.1787").Do(ctx))
	// printAnswers(bullishFutures.NewAmendOrder().UID("111872616831896").Symbol("BTC-USDC-PERP").OrderID("869492653091717121").NewSize("0.0012").NewPrice("118820").Do(ctx))
	//=======================OrderList
	n = "OrderList"
	//SPOT
	// printAnswers(binanceSpot.NewGetOrderList().Do(ctx))
	// printAnswers(bingxSpot.NewGetOrderList().Do(ctx))
	// printAnswers(bybitSpot.NewGetOrderList().Do(ctx))
	// printAnswers(gateioSpot.NewGetOrderList().Do(ctx))
	// printAnswers(mexcSpot.NewGetOrderList().Do(ctx))
	// printAnswers(bitgetSpot.NewGetOrderList().Do(ctx))
	// printAnswers(okxSpot.NewGetOrderList().Do(ctx))
	// printAnswers(huobiSpot.NewGetOrderList().Do(ctx))

	//FUTURES
	// printAnswers(binanceFutures.NewGetOrderList().Do(ctx))
	// printAnswers(bingxFutures.NewGetOrderList().Do(ctx))
	// printAnswers(bybitFutures.NewGetOrderList().Symbol("BTCUSDT").Do(ctx))
	// printAnswers(gateioFutures.NewGetOrderList().Do(ctx))
	// printAnswers(mexcSpot.NewGetOrderList().Symbol("MXUSDT").Do(ctx))
	// printAnswers(bitgetFutures.NewGetOrderList().Do(ctx))
	// printAnswers(okxFutures.NewGetOrderList().Do(ctx))
	printAnswers(bullishFutures.NewGetOrderList().Do(ctx))
	//=======================OrdersHistory
	n = "OrdersHistory"
	//SPOT
	// printAnswers(binanceSpot.NewOrdersHistory().Symbol("DOGEUSDT").Do(ctx))
	// printAnswers(bingxSpot.NewOrdersHistory().Do(ctx))
	// printAnswers(bybitSpot.NewOrdersHistory().Do(ctx))
	// printAnswers(gateioSpot.NewOrdersHistory().Do(ctx))
	// printAnswers(mexcSpot.NewOrdersHistory().Symbol("DOGEUSDT").Do(ctx))
	// printAnswers(bitgetSpot.NewOrdersHistory().Do(ctx))
	// printAnswers(okxSpot.NewOrdersHistory().Do(ctx))
	// printAnswers(huobiSpot.NewOrdersHistory().Do(ctx))

	//FUTURES
	// printAnswers(binanceFutures.NewOrdersHistory().Do(ctx))
	// printAnswers(bingxFutures.NewOrdersHistory().Do(ctx))
	// printAnswers(bybitFutures.NewOrdersHistory().Do(ctx))
	// printAnswers(gateioFutures.NewOrdersHistory().Do(ctx))
	// printAnswers(bitgetFutures.NewOrdersHistory().Do(ctx))
	// printAnswers(okxFutures.NewOrdersHistory().Do(ctx))

	//=======================PositionsHistory
	n = "PositionsHistory"
	//FUTURES
	// printAnswers(binanceFutures.NewPositionsHistory().Do(ctx))
	// printAnswers(bingxFutures.NewPositionsHistory().Symbol("TRX-USDT").StartTime(time.Now().UnixMilli() - (60 * 60 * 24 * 1000)).EndTime(time.Now().UnixMilli()).Do(ctx))
	// printAnswers(bybitFutures.NewPositionsHistory().Do(ctx))
	// printAnswers(gateioFutures.NewPositionsHistory().Do(ctx))
	// printAnswers(bitgetFutures.NewPositionsHistory().Do(ctx))
	// printAnswers(okxFutures.NewPositionsHistory().Symbol("ATOM-USDT-SWAP").Do(ctx))

	//=======================GetPositionMode
	n = "GetPositionMode"
	//FUTURES
	// printAnswers(binanceFutures.NewGetPositionMode().Do(ctx))
	// printAnswers(bingxFutures.NewGetPositionMode().Do(ctx))
	// printAnswers(bybitFutures.NewGetPositionMode().Do(ctx)) //processing
	// printAnswers(gateioFutures.NewGetPositionMode().Do(ctx))
	// printAnswers(bitgetFutures.NewGetPositionMode().Symbol("DOGEUSDT").Do(ctx))
	// printAnswers(okxFutures.NewGetPositionMode().Do(ctx))

	//=======================SetPositionMode
	n = "SetPositionMode"
	//FUTURES
	// printAnswers(binanceFutures.NewSetPositionMode().Mode(entity.PositionModeTypeOneWay).Do(ctx))
	// printAnswers(bingxFutures.NewSetPositionMode().Mode(entity.PositionModeTypeOneWay).Do(ctx))
	// printAnswers(bybitFutures.NewSetPositionMode().Symbol("DOGEUSDT").Mode(entity.PositionModeTypeOneWay).Do(ctx))
	// printAnswers(gateioFutures.NewSetPositionMode().Mode(entity.PositionModeTypeHedge).Do(ctx))
	// printAnswers(bitgetFutures.NewSetPositionMode().Mode(entity.PositionModeTypeOneWay).Do(ctx))
	// printAnswers(okxFutures.NewSetPositionMode().Mode(entity.PositionModeTypeHedge).Do(ctx))

	//=======================GetLeverage
	n = "GetLeverage"
	//FUTURES
	// printAnswers(binanceFutures.NewGetLeverage().Symbol("TRX-USDT").Do(ctx)) // process
	// printAnswers(bingxFutures.NewGetLeverage().Symbol("TRX-USDT").Do(ctx))
	// printAnswers(bybitFutures.NewGetLeverage().Do(ctx))  //processing
	// printAnswers(gateioFutures.NewGetLeverage().Symbol("DOGE_USDT").Do(ctx))
	// printAnswers(bitgetFutures.NewGetLeverage().Symbol("DOGEUSDT").Do(ctx))
	// printAnswers(okxFutures.NewGetLeverage().Symbol("BTC-USDT-SWAP").MarginMode(entity.MarginModeTypeCross).Do(ctx))

	//=======================SetLeverage
	n = "SetLeverage"
	//FUTURES
	// printAnswers(binanceFutures.NewSetLeverage().Symbol("DOGEUSDT").Leverage("50").Do(ctx))
	// printAnswers(bingxFutures.NewSetLeverage().Symbol("TRX-USDT").Leverage("50").PositionSide(entity.PositionSideTypeBoth).Do(ctx))
	// printAnswers(bybitFutures.NewSetLeverage().Symbol("DOGEUSDT").Leverage("10").Do(ctx))
	// printAnswers(gateioFutures.NewSetLeverage().Symbol("DOGE_USDT").Leverage("20").Do(ctx))
	// printAnswers(bitgetFutures.NewSetLeverage().Symbol("DOGEUSDT").Leverage("30").PositionSide(entity.PositionSideTypeLong).Do(ctx))
	// printAnswers(okxFutures.NewSetLeverage().Symbol("BTC-USDT-SWAP").Leverage("100").PositionSide(entity.PositionSideTypeShort).MarginMode(entity.MarginModeTypeIsolated).Do(ctx))

	//=======================GetMarginMode
	n = "GetMarginMode"
	//FUTURES
	// printAnswers(bingxFutures.NewGetMarginMode().Symbol("DOGE-USDT").Do(ctx))
	// printAnswers(gateioFutures.NewGetMarginMode().Symbol("BTC-USDT").Do(ctx))  //processing
	// printAnswers(bitgetFutures.NewGetMarginMode().Symbol("DOGEUSDT").Do(ctx))

	//=======================SetMarginMode
	n = "SetMarginMode"
	//FUTURES
	// printAnswers(binanceFutures.NewSetMarginMode().Symbol("TRX-USDT").MarginMode(entity.MarginModeTypeIsolated).Do(ctx))
	// printAnswers(bingxFutures.NewSetMarginMode().Symbol("TRX-USDT").MarginMode(entity.MarginModeTypeIsolated).Do(ctx))
	// printAnswers(bybitFutures.NewSetMarginMode().Symbol("DOGEUSDT").Leverage("20").MarginMode(entity.MarginModeTypeCross).Do(ctx)) //processing
	// printAnswers(gateioFutures.NewSetMarginMode().Symbol("BTC-USDT").Do(ctx))  //processing
	// printAnswers(bitgetFutures.NewSetMarginMode().Symbol("DOGEUSDT").MarginMode(entity.MarginModeTypeIsolated).Do(ctx))

	//=======================GetListenKey
	n = "GetListenKey"
	//SPOT
	// printAnswers(bingxSpot.NewGetListenKey().Do(ctx))
	// printAnswers(mexcSpot.NewGetListenKey().Do(ctx))

	//FUTURES
	// printAnswers(bingxFutures.NewGetListenKey().Do(ctx))

	//=======================ExtendListenKey
	n = "ExtendListenKey"
	//SPOT
	// printAnswers(bingxSpot.NewExtendListenKey().ListenKey("0e81716a1c67128521a48f066800b24420a14194b6637e43968d75e3a7293fce").Do(ctx))
	// printAnswers(mexcSpot.NewExtendListenKey().ListenKey("802935b2c04071cff608424de8cb0416585ab71783e9c932ea748f73c6efe223").Do(ctx))

	//FUTURES
	// printAnswers(bingxFutures.NewExtendListenKey().ListenKey("0e81716a1c67128521a48f066800b24420a14194b6637e43968d75e3a7293fce").Do(ctx))
	// printAnswers(bingxFutures.NewExtendListenKey().Do(ctx))

	//=======================SignAuthStream
	n = "SignAuthStream"
	// printAnswers(bybitSpot.NewSignAuthStream().TimeStamp(1753350824193).Do(ctx))
	// printAnswers(gateioSpot.NewSignAuthStream().TimeStamp(1753352964).Channel("futures.orders").Event("subscribe").Do(ctx))
	// printAnswers(bitgetSpot.NewSignAuthStream().TimeStamp(1753356762807).Do(ctx))
	// printAnswers(okxSpot.NewSignAuthStream().TimeStamp(1753356762807).Do(ctx))

	//=======================GenerateJWT
	n = "GenerateJWT"
	// printAnswers(bullishFutures.NewGenerateJWT().Do(ctx))

	//===========Not Exit
	// c := make(chan os.Signal, 1)
	// signal.Notify(c, os.Interrupt)
	// for {
	// 	<-c
	// 	return
	// }
}
