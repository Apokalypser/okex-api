package okex

import (
	"github.com/spf13/viper"
	"log"
	"testing"
)

func newSwapWSForTest() *SwapWS {
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
	ws := NewSwapWS(wsURL, accessKey, secretKey, passphrase, true)
	err = ws.SetProxy("socks5://127.0.0.1:1080")
	////ws.SetProxy("http://127.0.0.1:1080")
	if err != nil {
		log.Fatal(err)
	}
	return ws
}

func TestSwapWS_AllInOne(t *testing.T) {
	ws := newSwapWSForTest()
	ws.SetTickerCallback(func(tickers []WSTicker) {
		log.Printf("%#v", tickers)
	})
	ws.SetAccountCallback(func(accounts []WSAccount) {
		log.Printf("%#v", accounts)
	})
	//ws.SubscribeTicker("ticker_1", "BTC-USD-SWAP")
	//ws.SubscribeTrade("trade_1", "BTC-USD-SWAP")
	//ws.SubscribeDepthL2Tbt("depthL2_1", "BTC-USD-SWAP")
	//ws.SubscribeOrder("order_1", "BTC-USD-SWAP")
	ws.SubscribePosition("position_1", "BTC-USD-SWAP")
	ws.SubscribeAccount("account_1", "BTC-USD-SWAP")
	ws.Start()

	select {}
}
