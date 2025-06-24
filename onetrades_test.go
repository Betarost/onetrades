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
	// printAnswers(binanceSpot.NewGetInstrumentsInfo().Symbol("BTCUSDT").Do(ctx))
	// printAnswers(binanceFutures.NewGetInstrumentsInfo().Do(ctx))
	// printAnswers(bingxSpot.NewGetInstrumentsInfo().Symbol("BTC-USDT").Do(ctx))
	// printAnswers(bingxFutures.NewGetInstrumentsInfo().Symbol("BTC-USDT").Do(ctx))
	// printAnswers(bybitSpot.NewGetInstrumentsInfo().Symbol("BTCUSDT").Do(ctx))
	// printAnswers(bybitFutures.NewGetInstrumentsInfo().Symbol("BTCUSDT").Do(ctx))
	// printAnswers(gateioSpot.NewGetInstrumentsInfo().Symbol("BTC_USDT").Do(ctx))
	// printAnswers(gateioFutures.NewGetInstrumentsInfo().Symbol("BTC_USDT").Do(ctx))
	printAnswers(mexcSpot.NewGetInstrumentsInfo().Symbol("MXUSDT").Do(ctx))
	// printAnswers(mexcFutures.NewGetInstrumentsInfo().Symbol("BTC_USDT").Do(ctx))
	// printAnswers(bitgetSpot.NewGetInstrumentsInfo().Symbol("BTCUSDT").Do(ctx))
	// printAnswers(bitgetFutures.NewGetInstrumentsInfo().Symbol("BTCUSDT").Do(ctx))
	// printAnswers(okxSpot.NewGetInstrumentsInfo().Symbol("BTC-USDT").Do(ctx))
	// printAnswers(okxFutures.NewGetInstrumentsInfo().Symbol("BTC-USDT-SWAP").Do(ctx))
	// printAnswers(huobiSpot.NewGetInstrumentsInfo().Symbol("BTCUSDT").Do(ctx))
	// printAnswers(huobiFutures.NewGetInstrumentsInfo().Symbol("BTC-USDT").Do(ctx))

	//=======================AccountInfo
	n = "AccountInfo"
	// printAnswers(binanceSpot.NewGetAccountInfo().Do(ctx))
	// printAnswers(binanceFutures.NewGetAccountInfo().Do(ctx))
	// printAnswers(bingxSpot.NewGetAccountInfo().Do(ctx))
	// printAnswers(bingxFutures.NewGetAccountInfo().Do(ctx))

	// printAnswers(gateioSpot.NewGetAccountInfo().Do(ctx))
	// printAnswers(gateioFutures.NewGetAccountInfo().Do(ctx))
	printAnswers(mexcSpot.NewGetAccountInfo().Do(ctx))
	// printAnswers(mexcFutures.NewGetAccountInfo().Do(ctx))
	// printAnswers(mexcFutures.NewGetAccountInfo().Do(ctx))

	//=======================FundingAccountBalance
	n = "FundingAccountBalance"
	// printAnswers(bingxSpot.NewGetFundingAccountBalance().Do(ctx))
	// printAnswers(bingxFutures.NewGetFundingAccountBalance().Do(ctx))
	// printAnswers(mexcSpot.NewGetFundingAccountBalance().Do(ctx))

	//=======================TradingAccountBalance
	n = "TradingAccountBalance"
	// printAnswers(bingxSpot.NewGetTradingAccountBalance().Do(ctx))

	//=======================GetBalance
	n = "GetBalance"
	printAnswers(mexcSpot.NewGetBalance().Do(ctx))

	//=======================PlaceOrder
	n = "PlaceOrder"
	// printAnswers(mexcSpot.NewPlaceOrder().Symbol("MXUSDT").Side(entity.SideTypeBuy).OrderType(entity.OrderTypeLimit).Size("0.5").Price("2.35").Do(ctx))

	//=======================AmendOrder
	n = "AmendOrder"

	//=======================CancelOrder
	n = "CancelOrder"
	// printAnswers(mexcSpot.NewCancelOrder().Symbol("MXUSDT").OrderID("C02__566124671469281280120").Do(ctx))

	//=======================OrderList
	n = "OrderList"
	printAnswers(mexcSpot.NewGetOrderList().Symbol("MXUSDT").Do(ctx))
	//=======================GetPositions
	n = "GetPositions"

	//=======================SetLeverage
	n = "SetLeverage"

	//=======================SetPositionMode
	n = "SetPositionMode"

	//===========Not Exit
	// c := make(chan os.Signal, 1)
	// signal.Notify(c, os.Interrupt)
	// for {
	// 	<-c
	// 	return
	// }
}
