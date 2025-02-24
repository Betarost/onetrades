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

	//=====================BINANCE GET BALANCE==========================
	// binanceKey := viper.GetString("BINANCE_API")
	// binanceSecret := viper.GetString("BINANCE_SECRET")
	// client := onetrades.NewFutureBinanceClient(binanceKey, binanceSecret)
	// res, err := client.NewGetAccountBalance().Do(context.Background())
	// t.Logf("Results: %+v  %v", res, err)

	//=====================BYBIT GET BALANCE======================-=====
	bybitKey := viper.GetString("BYBIT_API")
	bybitSecret := viper.GetString("BYBIT_SECRET")
	client := NewFutureBybitClient(bybitKey, bybitSecret)
	res, err := client.NewGetAccountBalance().Do(context.Background())
	t.Logf("Results: %+v  %v", res, err)
}
