package main

import (
	"fmt"
	"github.com/marcos-gonalons/tradingview-scraper/v2"
	"os"
	"os/signal"
	"syscall"
)

func main() { //(t *testing.T) {
	tradingviewsocket, err := tradingview.Connect(
		func(symbol string, data *tradingview.QuoteData) {
			//fmt.Printf("symbol:%s\n", symbol)
			//resp1, _ := json.Marshal(data)
			//fmt.Printf("respose:%s\n", string(resp1))

			if data.Price != nil {
				fmt.Printf("price=%f\n", *data.Price)
			}

			if data.Time != nil {
				fmt.Printf("time=%d\n", *data.Time)
			}

			/*
				if data.Volume != nil {
					fmt.Printf("volume=%f\n", *data.Volume)
				}
				//如果没有数据,证明没有任何change
				if data.Bid != nil {
					fmt.Printf("bid=%f\n", *data.Bid)
				}
				if data.Ask != nil {
					fmt.Printf("ask=%f\n", *data.Ask)
				}
			*/
		},
		func(err error, context string) {
			fmt.Printf("%#v", "error -> "+err.Error())
			fmt.Printf("%#v", "context -> "+context)
		},
	)
	if err != nil {
		panic("Error while initializing the trading view socket -> " + err.Error())
	}

	tradingviewsocket.AddSymbol("KRX:KOSPI200")

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGKILL)
	<-quit
}
