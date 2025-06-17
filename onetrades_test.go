package onetrades

import (
	"context"
	"log"
	"os"
	"os/signal"
	"testing"

	"github.com/Betarost/onetrades/bingx"
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
	// binanceKey := viper.GetString("BINANCE_API")
	// binanceSecret := viper.GetString("BINANCE_SECRET")
	bingxKey := viper.GetString("BINGX_API")
	bingxSecret := viper.GetString("BINGX_SECRET")
	// okxKey := viper.GetString("OKX_API")
	// okxSecret := viper.GetString("OKX_SECRET")
	// okxMemo := viper.GetString("OKX_MEMO")
	// bybitKey := viper.GetString("BYBIT_API")
	// bybitSecret := viper.GetString("BYBIT_SECRET")
	// mexcKey := viper.GetString("MEX_API")
	// mexcSecret := viper.GetString("MEX_SECRET")
	// mexcMemo := viper.GetString("MEX_MEMO")
	// gateKey := viper.GetString("GATE_API")
	// gateSecret := viper.GetString("GATE_SECRET")

	//==========================CLIENTS==========================
	// binanceSpot := binance.NewSpotClient(binanceKey, binanceSecret)
	// binanceFutures := binance.NewFuturesClient(binanceKey, binanceSecret)
	bingxSpot := bingx.NewSpotClient(bingxKey, bingxSecret)
	bingxFutures := bingx.NewFuturesClient(bingxKey, bingxSecret)
	// okxSpot := okx.NewSpotClient(okxKey, okxSecret, okxMemo)
	// okxFutures := okx.NewFuturesClient(okxKey, okxSecret, okxMemo)
	// bybitSpot := bybit.NewSpotClient(bybitKey, bybitSecret)
	// bybitFutures := bybit.NewFuturesClient(bybitKey, bybitSecret)
	//==========================BINGX CLIENT==========================
	//==========================GATE CLIENT==========================
	//==========================MEXC CLIENT==========================
	//==========================BITGET CLIENT==========================
	//==========================HUOBI CLIENT==========================
	// binanceSpot.Debug = true
	// binanceFutures.Debug = true
	//=======================InstrumentsInfo
	n = "InstrumentsInfo"
	// printAnswers(binanceSpot.NewGetInstrumentsInfo().Symbol("BTCUSDT").Do(ctx))
	// printAnswers(binanceFutures.NewGetInstrumentsInfo().Do(ctx))
	printAnswers(bingxSpot.NewGetInstrumentsInfo().Symbol("BTC-USDT").Do(ctx))
	printAnswers(bingxFutures.NewGetInstrumentsInfo().Symbol("BTC-USDT").Do(ctx))
	// printAnswers(okxSpot.NewGetInstrumentsInfo().Symbol("BTC-USDT").Do(ctx))
	// printAnswers(okxFutures.NewGetInstrumentsInfo().Symbol("BTC-USDT-SWAP").Do(ctx))
	// printAnswers(bybitSpot.NewGetInstrumentsInfo().Symbol("BTCUSDT").Do(ctx))
	// printAnswers(bybitFutures.NewGetInstrumentsInfo().Symbol("BTCUSDT").Do(ctx))

	//==========================BYBIT==========================
	//==========================Bybit SPOT==========================

	//======================= GET AccountInfo
	// res, err := client.NewGetAccountInfo().Do(context.Background())
	// log.Println("=Error=", err)
	// log.Printf("=res= %+v", res)

	//======================= GET Trading Balance
	// res, err := client.NewGetTradingAccountBalance().Do(context.Background())
	// log.Println("=Error=", err)
	// log.Printf("=res= %+v", res)

	//======================= GET Funding Balance
	// res, err := client.NewGetFundingAccountBalance().Do(context.Background())
	// log.Println("=Error=", err)
	// log.Printf("=res= %+v", res)

	//======================= PlaceOrder
	// res, err := client.NewPlaceOrder().Symbol("TRXUSDT").Side(entity.SideTypeSell).Size("5").OrderType(entity.OrderTypeMarket).Do(context.Background())
	// res, err := client.NewPlaceOrder().Symbol("TRX-USDT").Side(entity.SideTypeSell).Size("1").OrderType(entity.OrderTypeLimit).Price("0.28730").Do(context.Background())
	// res, err := client.NewPlaceOrder().Symbol("TRX-USDT").Side(entity.SideTypeSell).Size("1").OrderType(entity.OrderTypeLimit).Price("0.28730").TpPrice("0.28670").SlPrice("0.28850").Do(context.Background())
	// res, err := client.NewPlaceOrder().Symbol("TRX-USDT").Side(entity.SideTypeSell).Size("1").OrderType(entity.OrderTypeLimit).Price("0.28510").Do(context.Background())
	// log.Println("=Error=", err)
	// log.Printf("=res= %+v", res)

	//======================END Bybit==========================

	//==========================OKX SPOT==========================
	// okxKey := viper.GetString("OKX_API")
	// okxSecret := viper.GetString("OKX_SECRET")
	// okxMemo := viper.GetString("OKX_MEMO")
	// client := okx.NewSpotClient(okxKey, okxSecret, okxMemo)

	//======================= GET AccountInfo
	// res, err := client.NewGetAccountInfo().Do(context.Background())
	// log.Println("=Error=", err)
	// log.Printf("=res= %+v", res)

	//======================= GET Trading Balance
	// res, err := client.NewGetTradingAccountBalance().Do(context.Background())
	// log.Println("=Error=", err)
	// log.Printf("=res= %+v", res)

	//======================= GET Funding Balance
	// res, err := client.NewGetFundingAccountBalance().Do(context.Background())
	// log.Println("=Error=", err)
	// log.Printf("=res= %+v", res)

	//=======================Get InstrumentsInfo
	// res, err := client.NewGetInstrumentsInfo().Symbol("BTC-USDT").Do(context.Background())
	// log.Println("=Error=", err)
	// log.Printf("=res= %+v", res)

	//=======================Get OrderList
	// res, err := client.NewGetOrderList().Do(context.Background())
	// log.Println("=Error=", err)
	// log.Printf("=res= %+v", res)

	//=======================Amend Order
	// res, err := client.NewAmendOrder().Symbol("TRX-USDT").OrderID("2582962054382215168").NewSize("8").Do(context.Background())
	// log.Println("=Error=", err)
	// log.Printf("=res= %+v", res)

	//=======================Cancel Order
	// res, err := client.NewCancelOrder().Symbol("TRX-USDT").OrderID("2581988433413267456").Do(context.Background())
	// log.Println("=Error=", err)
	// log.Printf("=res= %+v", res)

	//=======================MultiCancelOrders
	// res, err := client.NewMultiCancelOrders().Symbol("TRX-USDT").OrderIDs([]string{"2582956581687910400", "2582956136789696512"}).Do(context.Background())
	// log.Println("=Error=", err)
	// log.Printf("=res= %+v", res)

	//======================= PlaceOrder
	// res, err := client.NewPlaceOrder().Symbol("TRX-USDT").Side(entity.SideTypeSell).Size("1").OrderType(entity.OrderTypeLimit).Price("0.28730").Do(context.Background())
	// res, err := client.NewPlaceOrder().Symbol("TRX-USDT").Side(entity.SideTypeSell).Size("1").OrderType(entity.OrderTypeLimit).Price("0.28730").TpPrice("0.28670").SlPrice("0.28850").Do(context.Background())
	// // res, err := client.NewPlaceOrder().Symbol("TRX-USDT").Side(entity.SideTypeSell).Size("1").OrderType(entity.OrderTypeLimit).Price("0.28510").Do(context.Background())
	// log.Println("=Error=", err)
	// log.Printf("=res= %+v", res)

	//==========================OKX FUTURES==========================
	// okxKey := viper.GetString("OKX_API")
	// okxSecret := viper.GetString("OKX_SECRET")
	// okxMemo := viper.GetString("OKX_MEMO")
	// client := okx.NewFuturesClient(okxKey, okxSecret, okxMemo)
	//=======================Get InstrumentsInfo
	// res, err := client.NewGetInstrumentsInfo().Symbol("BTC-USDT-SWAP").Do(context.Background())
	// log.Println("=Error=", err)
	// log.Printf("=res= %+v", res)

	//======================= SET PositionMode
	// res, err := client.NewSetPositionMode().Mode(entity.PositionModeTypeHedge).Do(context.Background())
	// log.Println("=Error=", err)
	// log.Printf("=res= %+v", res)
	//======================= SET Leverage
	// res, err := client.NewSetLeverage().Symbol("DOGE-USDT-SWAP").Leverage("20").MarginMode(entity.MarginModeTypeIsolated).PositionSide(entity.PositionSideTypeLong).Do(context.Background())
	// log.Println("=Error=", err)
	// log.Printf("=res= %+v", res)
	//======================= GET Position
	// res, err := client.NewGetPositions().Do(context.Background())
	// log.Println("=Error=", err)
	// log.Printf("=res= %+v", res)

	//=======================Get OrderList
	// res, err := client.NewGetOrderList().Do(context.Background())
	// log.Println("=Error=", err)
	// log.Printf("=res= %+v", res)

	//=======================Cancel Order
	// res, err := client.NewCancelOrder().Symbol("BTC-USDT-SWAP").OrderID("2583772267918123008").Do(context.Background())
	// log.Println("=Error=", err)
	// log.Printf("=res= %+v", res)

	//=======================Amend Order
	// res, err := client.NewAmendOrder().Symbol("TRX-USDT").OrderID("2582962054382215168").NewSize("8").Do(context.Background())
	// log.Println("=Error=", err)
	// log.Printf("=res= %+v", res)

	//======================= PlaceOrder
	// res, err := client.NewPlaceOrder().Symbol("BTC-USDT-SWAP").PositionSide(entity.PositionSideTypeShort).Side(entity.SideTypeSell).Size("0.01").OrderType(entity.OrderTypeLimit).Price("109850").TpPrice("109500").TradeMode(entity.MarginModeTypeCross).Do(context.Background())
	// // res, err := client.NewPlaceOrder().Symbol("TRX-USDT").Side(entity.SideTypeSell).Size("1").OrderType(entity.OrderTypeLimit).Price("0.28510").Do(context.Background())
	// log.Println("=Error=", err)
	// log.Printf("=res= %+v", res)

	//======================END OKX==========================

	//===========================================================
	//======================END NEW==============================
	//===========================================================

	//===========================================================
	//======================OPTION==============================
	//===========================================================

	//==========================OKX==========================
	// okxKey := viper.GetString("OKX_API")
	// okxSecret := viper.GetString("OKX_SECRET")
	// okxMemo := viper.GetString("OKX_MEMO")
	// client := NewOptionOKXClient(okxKey, okxSecret, okxMemo)
	//======================= GET ContractsInfo
	// res, err := client.NewGetContractsInfo().Symbol("BTC-USD-250627-105000-C").Do(context.Background())
	// t.Logf("Results: %d  %v %+v", len(res), err, res)
	// ======================= GET MarketData
	// res, err := client.NewGetMarketData().ExpTime("250725").Do(context.Background())
	// t.Logf("Results: %d  %v", len(res), err)
	// ======================= GET OrderBook
	// res, err := client.NewGetOrderBook().Symbol("BTC-USD-250926-110000-C").Do(context.Background())
	// t.Logf("Results: %+v  %v", res, err)
	//=======================Trade PlaceOrder
	// res, err := client.NewTradePlaceOrder().Symbol("BTC-USD-250627-105000-C").Side(entity.SideTypeBuy).Size("1").Price("0.0455").OrderType(entity.OrderTypeLimit).Isolated(true).Do(context.Background())
	// t.Logf("Results: %+v  %v", res, err)
	//======================= GET Position
	// res, err := client.NewGetPositions().Do(context.Background())
	// t.Logf("Results: %+v  %v", res, err)
	//======================= GET Balance
	// res, err := client.NewGetAccountBalance().Do(context.Background())
	// t.Logf("Results: %+v  %v", res, err)
	//======================END OKX==========================

	//===========================================================
	//======================FUTURES==============================
	//===========================================================

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

	//===========Not Exit
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	for {
		<-c
		return
	}

}
