package okex

import (
	"github.com/spf13/viper"
	"log"
	"testing"
)

func newFuturesWSForTest() *FuturesWS {
	viper.SetConfigName("test_config")
	viper.AddConfigPath("./conf")
	err := viper.ReadInConfig()
	if err != nil {
		log.Panic(err)
	}

	accessKey := viper.GetString("access_key")
	secretKey := viper.GetString("secret_key")
	passphrase := viper.GetString("passphrase")
	wsURL := "wss://real.okex.com:8443/ws/v3"

	ws := NewFuturesWS(wsURL,
		accessKey, secretKey, passphrase, true)
	err = ws.SetProxy("socks5://127.0.0.1:1080")
	//ws.SetProxy("http://127.0.0.1:1080")
	if err != nil {
		log.Fatal(err)
	}
	return ws
}

func TestFuturesWS_AllInOne(t *testing.T) {
	ws := newFuturesWSForTest()
	ws.SetTickerCallback(func(tickers []WSTicker) {
		log.Printf("%#v", tickers)
	})
	ws.SetAccountCallback(func(accounts []WSAccount) {
		log.Printf("%#v", accounts)
	})
	//ws.SubscribeTicker("ticker_1", "BTC-USD-200626")
	//ws.SubscribeTrade("trade_1", "BTC-USD-200626")
	//ws.SubscribeDepthL2Tbt("depthL2_1", "BTC-USD-200626")
	//ws.SubscribeOrder("order_1", "BTC-USD-200626")
	ws.SubscribePosition("position_1", "BTC-USD-200626")
	ws.SubscribeAccount("account_1", "BTC") // BTC/BTC-USDT
	ws.Start()

	select {}
}

func TestFuturesWS_Depth20(t *testing.T) {
	ws := newFuturesWSForTest()
	ws.SetDepth20SnapshotCallback(func(ob *OrderBook) {
		log.Printf("%#v", ob)
	})
	ws.SubscribeDepthL2Tbt("depthL2_1", "BTC-USD-200626")
	ws.Start()

	select {}
}

func TestFuturesWS_SubscribeOrder(t *testing.T) {
	ws := newFuturesWSForTest()
	ws.SetOrderCallback(func(orders []WSOrder) {
		log.Printf("%#v", orders)
	})
	ws.SubscribeOrder("order_1", "BTC-USD-200626")
	ws.Start()

	select {}
}
