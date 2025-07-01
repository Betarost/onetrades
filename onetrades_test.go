package onetrades

import (
	"context"
	"log"
	"testing"

	"github.com/Betarost/onetrades/binance"
	"github.com/Betarost/onetrades/bingx"
	"github.com/Betarost/onetrades/bitget"
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
	mexcFutures := mexc.NewFuturesClient(mexcKey, mexcSecret)
	bitgetSpot := bitget.NewSpotClient(bitgetKey, bitgetSecret, bitgetMemo)
	bitgetFutures := bitget.NewFuturesClient(bitgetKey, bitgetSecret, bitgetMemo)
	okxSpot := okx.NewSpotClient(okxKey, okxSecret, okxMemo)
	okxFutures := okx.NewFuturesClient(okxKey, okxSecret, okxMemo)
	huobiSpot := huobi.NewSpotClient(huobiKey, huobiSecret)
	huobiFutures := huobi.NewFuturesClient(huobiKey, huobiSecret)

	binanceSpot.Debug = false
	binanceFutures.Debug = false
	bingxSpot.Debug = false
	bingxFutures.Debug = false
	bybitSpot.Debug = false
	bybitFutures.Debug = false
	gateioSpot.Debug = false
	gateioFutures.Debug = false
	mexcSpot.Debug = false
	mexcFutures.Debug = false
	bitgetSpot.Debug = false
	bitgetFutures.Debug = false
	okxSpot.Debug = false
	okxFutures.Debug = false
	huobiSpot.Debug = false
	huobiFutures.Debug = false

	//=======================InstrumentsInfo
	n = "InstrumentsInfo"
	//SPOT
	// printAnswers(bingxSpot.NewGetInstrumentsInfo().Symbol("BTC-USDT").Do(ctx))
	// printAnswers(bybitSpot.NewGetInstrumentsInfo().Symbol("DOGEUSDT").Do(ctx))
	// printAnswers(gateioSpot.NewGetInstrumentsInfo().Symbol("DOGE_USDT").Do(ctx))
	// printAnswers(mexcSpot.NewGetInstrumentsInfo().Symbol("MXUSDT").Do(ctx))
	// printAnswers(bitgetSpot.NewGetInstrumentsInfo().Symbol("BTCUSDT").Do(ctx))
	// printAnswers(okxSpot.NewGetInstrumentsInfo().Symbol("DOGE-USDT").Do(ctx))
	// printAnswers(huobiSpot.NewGetInstrumentsInfo().Do(ctx))

	//FUTURES
	// printAnswers(bingxFutures.NewGetInstrumentsInfo().Symbol("BTC-USDT").Do(ctx))
	// printAnswers(bybitFutures.NewGetInstrumentsInfo().Symbol("BTCUSDT").Do(ctx))
	// printAnswers(gateioFutures.NewGetInstrumentsInfo().Symbol("BTC_USDT").Do(ctx))
	// printAnswers(bitgetFutures.NewGetInstrumentsInfo().Symbol("BTCUSDT").Do(ctx))
	// printAnswers(okxFutures.NewGetInstrumentsInfo().Symbol("BTC-USDT-SWAP").Do(ctx))
	printAnswers(huobiFutures.NewGetInstrumentsInfo().Symbol("BTC-USDT").Do(ctx))

	// printAnswers(binanceSpot.NewGetInstrumentsInfo().Symbol("BTCUSDT").Do(ctx))
	// printAnswers(binanceFutures.NewGetInstrumentsInfo().Do(ctx))
	// printAnswers(huobiFutures.NewGetInstrumentsInfo().Symbol("BTC-USDT").Do(ctx))

	//=======================AccountInfo
	n = "AccountInfo"
	//SPOT
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

	// printAnswers(binanceSpot.NewGetAccountInfo().Do(ctx))
	// printAnswers(binanceFutures.NewGetAccountInfo().Do(ctx))

	//=======================GetBalance
	n = "GetBalance"
	//SPOT
	// printAnswers(bingxSpot.NewGetBalance().Do(ctx))
	// printAnswers(bybitSpot.NewGetBalance().Do(ctx))
	// printAnswers(gateioSpot.NewGetBalance().Do(ctx))
	// printAnswers(mexcSpot.NewGetBalance().Do(ctx))
	// printAnswers(bitgetSpot.NewGetBalance().Do(ctx))
	// printAnswers(okxSpot.NewGetBalance().Do(ctx))
	// printAnswers(huobiSpot.NewGetBalance().UID("53799773").Do(ctx))
	// printAnswers(huobiSpot.NewGetBalance().UID("69069265").Do(ctx))

	//FUTURES
	// printAnswers(bingxFutures.NewGetBalance().Do(ctx))
	// printAnswers(bybitFutures.NewGetBalance().Do(ctx))
	// printAnswers(gateioFutures.NewGetBalance().Do(ctx))
	// printAnswers(bitgetFutures.NewGetBalance().Do(ctx))
	// printAnswers(okxFutures.NewGetBalance().Do(ctx))
	// printAnswers(huobiFutures.NewGetBalance().Do(ctx))

	//=======================PlaceOrder
	n = "PlaceOrder"
	//SPOT
	// printAnswers(bingxSpot.NewPlaceOrder().Symbol("PEPE-USDT").Side(entity.SideTypeBuy).OrderType(entity.OrderTypeLimit).Price("0.00000910").Size("180000").Do(ctx))
	// printAnswers(bybitSpot.NewPlaceOrder().Symbol("DOGEUSDT").Side(entity.SideTypeBuy).OrderType(entity.OrderTypeLimit).Price("0.15870").Size("35").Do(ctx))
	// printAnswers(gateioSpot.NewPlaceOrder().Symbol("DOGE_USDT").Side(entity.SideTypeBuy).OrderType(entity.OrderTypeLimit).Size("35").Price("0.15870").Do(ctx))
	// printAnswers(mexcSpot.NewPlaceOrder().Symbol("DOGEUSDT").Side(entity.SideTypeBuy).OrderType(entity.OrderTypeLimit).Price("0.15870").Size("10").Do(ctx))
	// printAnswers(bitgetSpot.NewPlaceOrder().Symbol("DOGEUSDT").Side(entity.SideTypeSell).OrderType(entity.OrderTypeLimit).Price("0.16250").Size("10").Do(ctx))
	// printAnswers(okxSpot.NewPlaceOrder().Symbol("DOGE-USDT").Side(entity.SideTypeBuy).OrderType(entity.OrderTypeLimit).Price("0.15870").Size("10").Do(ctx))
	// printAnswers(huobiSpot.NewPlaceOrder().Symbol("DOGEUSDT").Side(entity.SideTypeBuy).OrderType(entity.OrderTypeLimit).Size("100").Price("0.15666").UID("69069265").Do(ctx))

	//FUTURES

	// printAnswers(gateioFutures.NewPlaceOrder().Symbol("TRX_USDT").Side(entity.SideTypeBuy).Size("1").OrderType(entity.OrderTypeMarket).ClientOrderID("t-235k").Do(ctx))
	// printAnswers(gateioFutures.NewPlaceOrder().Symbol("TRX_USDT").Side(entity.SideTypeBuy).Size("1").Price("0.2687").OrderType(entity.OrderTypeLimit).ClientOrderID("t-235k2").Do(ctx))

	//=======================AmendOrder
	n = "AmendOrder"
	//SPOT
	// printAnswers(bybitSpot.NewAmendOrder().Symbol("DOGEUSDT").OrderID("1981265133053505280").NewPrice("0.15766").Do(ctx))
	// printAnswers(gateioSpot.NewAmendOrder().Symbol("DOGE_USDT").OrderID("865644853947").NewSize("37").Do(ctx))
	// printAnswers(okxSpot.NewAmendOrder().Symbol("DOGE-USDT").OrderID("2631421072347750400").NewPrice("0.15766").Do(ctx))

	//FUTURES

	// printAnswers(gateioFutures.NewAmendOrder().OrderID("186336434877910305").NewSize("2").Do(ctx))

	//=======================CancelOrder
	n = "CancelOrder"
	//SPOT
	// printAnswers(bingxSpot.NewCancelOrder().Symbol("PEPE-USDT").OrderID("1937962543553085440").Do(ctx))
	// printAnswers(bybitSpot.NewCancelOrder().Symbol("DOGEUSDT").OrderID("1980932944252202752").Do(ctx))
	// printAnswers(gateioSpot.NewCancelOrder().Symbol("DOGE_USDT").OrderID("865408051601").Do(ctx))
	// printAnswers(mexcSpot.NewCancelOrder().Symbol("DOGEUSDT").OrderID("C02__566648763716751360120").Do(ctx))
	// printAnswers(bitgetSpot.NewCancelOrder().Symbol("DOGEUSDT").OrderID("1322436121694060548").Do(ctx))
	// printAnswers(okxSpot.NewCancelOrder().Symbol("DOGE-USDT").OrderID("2629979625991954432").Do(ctx))
	// printAnswers(huobiSpot.NewCancelOrder().Symbol("DOGEUSDT").OrderID("1367428516065452").Do(ctx))

	//FUTURES

	// printAnswers(gateioFutures.NewCancelOrder().OrderID("186336434877885773").Do(ctx))

	//=======================OrderList
	n = "OrderList"
	//SPOT
	// printAnswers(bingxSpot.NewGetOrderList().Do(ctx))
	// printAnswers(bybitSpot.NewGetOrderList().Do(ctx))
	// printAnswers(gateioSpot.NewGetOrderList().Do(ctx))
	// printAnswers(mexcSpot.NewGetOrderList().Symbol("DOGEUSDT").Do(ctx))
	// printAnswers(bitgetSpot.NewGetOrderList().Do(ctx))
	// printAnswers(okxSpot.NewGetOrderList().Do(ctx))
	// printAnswers(huobiSpot.NewGetOrderList().Do(ctx))

	//FUTURES

	// printAnswers(mexcSpot.NewGetOrderList().Symbol("MXUSDT").Do(ctx))

	//=======================GetPositions
	n = "GetPositions"
	//FUTURES
	// printAnswers(gateioFutures.NewGetPositions().Do(ctx))
	//=======================SetLeverage
	n = "SetLeverage"
	//FUTURES
	//=======================SetPositionMode
	n = "SetPositionMode"
	//FUTURES
	// printAnswers(gateioFutures.NewSetPositionMode().Mode(entity.PositionModeTypeOneWay).Do(ctx))

	//=======================FundingAccountBalance
	n = "FundingAccountBalance"
	// printAnswers(bingxSpot.NewGetFundingAccountBalance().Do(ctx))
	// printAnswers(bingxFutures.NewGetFundingAccountBalance().Do(ctx))
	// printAnswers(mexcSpot.NewGetFundingAccountBalance().Do(ctx))

	//=======================TradingAccountBalance
	n = "TradingAccountBalance"
	// printAnswers(bybitSpot.NewGetTradingAccountBalance().Do(ctx))
	// printAnswers(bingxSpot.NewGetTradingAccountBalance().Do(ctx))

	//===========Not Exit
	// c := make(chan os.Signal, 1)
	// signal.Notify(c, os.Interrupt)
	// for {
	// 	<-c
	// 	return
	// }
}
