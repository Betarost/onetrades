package onetrades

import (
	"context"
	"log"
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
	//======================= GET Balance
	// res, err := client.NewGetAccountBalance().Do(context.Background())
	// t.Logf("Results: %+v  %v", res, err)
	//======================= GET Position
	// res, err := client.NewGetPositions().Do(context.Background())
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
	res, err := client.NewGetContractsInfo().Do(context.Background())
	t.Logf("Results: %+v  %v", res, err)
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

}
