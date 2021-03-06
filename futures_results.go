package okex

/*
 OKEX futures contract api response results
 @author Tony Tian
 @date 2018-03-17
 @version 1.0.0
*/

type ServerTime struct {
	Iso   string `json:"iso"`
	Epoch string `json:"epoch"`
}

type ExchangeRate struct {
	InstrumentId string  `json:"instrument_id"`
	Rate         float64 `json:"rate,string"`
	Timestamp    string  `json:"timestamp"`
}

type BizWarmTips struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Msg     string `json:"msg"`
}

type Result struct {
	Result bool `json:"result"`
}

type PageResult struct {
	From  int
	To    int
	Limit int
}

type FuturesPosition struct {
	BizWarmTips
	Result
	MarginMode    string
	CrossPosition []FuturesCrossPositionHolding
	FixedPosition []FuturesFixedPositionHolding
}

type FuturesCrossPosition struct {
	Result
	MarginMode    string                        `json:"margin_mode"`
	CrossPosition []FuturesCrossPositionHolding `json:"holding"`
}

type FuturesFixedPosition struct {
	Result
	MarginMode    string                        `json:"margin_mode"`
	FixedPosition []FuturesFixedPositionHolding `json:"holding"`
}

type FuturesCrossPositionHolding struct {
	FuturesPositionBase
	LiquidationPrice float64 `json:"liquidation_price,string"`
	Leverage         float64 `json:"leverage,string"`
}

type FuturesFixedPositionHolding struct {
	FuturesPositionBase
	LongMargin      float64 `json:"long_margin,string"`
	LongLiquiPrice  float64 `json:"long_liqui_price,string"`
	LongPnlRatio    float64 `json:"long_pnl_ratio,string"`
	LongLeverage    float64 `json:"long_leverage,string"`
	ShortMargin     float64 `json:"short_margin,string"`
	ShortLiquiPrice float64 `json:"short_liqui_price,string"`
	ShortPnlRatio   float64 `json:"short_pnl_ratio,string"`
	ShortLeverage   float64 `json:"short_leverage,string"`
}

type FuturesPositionBase struct {
	LongQty              float64 `json:"long_qty,string"`
	LongAvailQty         float64 `json:"long_avail_qty,string"`
	LongAvgCost          float64 `json:"long_avg_cost,string"`
	LongSettlementPrice  float64 `json:"long_settlement_price,string"`
	RealizedPnl          float64 `json:"realized_pnl,string"`
	ShortQty             float64 `json:"short_qty,string"`
	ShortAvailQty        float64 `json:"short_avail_qty,string"`
	ShortAvgCost         float64 `json:"short_avg_cost,string"`
	ShortSettlementPrice float64 `json:"short_settlement_price,string"`
	InstrumentId         string  `json:"instrument_id"`
	CreatedAt            string  `json:"created_at"`
	UpdatedAt            string  `json:"updated_at"`
}

type FuturesAccount struct {
	BizWarmTips
	Result
	MarginMode   string
	CrossAccount map[string]FuturesCrossAccount
	FixedAccount map[string]FuturesFixedAccount
}

type FuturesAccountsContract struct {
	AvailableQty      string  `json:"available_qty"`
	FixedBalance      string  `json:"fixed_balance"`
	InstrumentID      string  `json:"instrument_id"`
	MarginForUnfilled string  `json:"margin_for_unfilled"`
	MarginFrozen      string  `json:"margin_frozen"`
	RealizedPnl       string  `json:"realized_pnl"`
	UnrealizedPnl     string  `json:"unrealized_pnl"`
	MarginRatio       string  `json:"margin_ratio"`
	MaintMarginRatio  string  `json:"maint_margin_ratio"`
	CanWithdraw       string  `json:"can_withdraw"`
	Equity            string  `json:"equity"`
	MarginMode        string  `json:"margin_mode"`
	TotalAvailBalance float64 `json:"total_avail_balance,string"`
	AutoMargin        string  `json:"auto_margin"`
	Underlying        string  `json:"underlying"`
}

type GetFuturesAccountsResult struct {
	Info struct {
		Btc     FuturesAccountsContract `json:"btc"`
		Eth     FuturesAccountsContract `json:"eth"`
		Etc     FuturesAccountsContract `json:"etc"`
		Ltc     FuturesAccountsContract `json:"ltc"`
		Xrp     FuturesAccountsContract `json:"xrp"`
		Bsv     FuturesAccountsContract `json:"bsv"`
		Trx     FuturesAccountsContract `json:"trx"`
		Bch     FuturesAccountsContract `json:"bch"`
		Eos     FuturesAccountsContract `json:"eos"`
		BtcUSDT FuturesAccountsContract `json:"btc-usdt"`
		EthUSDT FuturesAccountsContract `json:"eth-usdt"`
		EtcUSDT FuturesAccountsContract `json:"etc-usdt"`
		LtcUSDT FuturesAccountsContract `json:"ltc-usdt"`
		XrpUSDT FuturesAccountsContract `json:"xrp-usdt"`
		BsvUSDT FuturesAccountsContract `json:"bsv-usdt"`
		TrxUSDT FuturesAccountsContract `json:"trx-usdt"`
		BchUSDT FuturesAccountsContract `json:"bch-usdt"`
		EosUSDT FuturesAccountsContract `json:"eos-usdt"`
	} `json:"info"`
}

type FuturesMarkdown struct {
	BizWarmTips
	InstrumentId string  `json:"instrument_id"`
	Timestamp    string  `json:"timestamp"`
	MarkPrice    float32 `json:"mark_price"`
}

type FuturesFixedAccountInfo struct {
	Result
	Info map[string]FuturesFixedAccount `json:"info"`
}

type FuturesCrossAccountInfo struct {
	Result
	Info map[string]FuturesCrossAccount `json:"info"`
}

type FuturesFixedAccount struct {
	MarginMode        string                         `json:"margin_mode"`
	Equity            float64                        `json:"equity,string"`
	TotalAvailBalance float64                        `json:"total_avail_balance,string"`
	Contracts         []FuturesFixedAccountContracts `json:"contracts"`
}

type FuturesFixedAccountContracts struct {
	AvailableQty      float64 `json:"available_qty,string"`
	FixedBalance      float64 `json:"fixed_balance,string"`
	InstrumentId      string  `json:"instrument_id"`
	MarginFixed       float64 `json:"margin_fixed,string"`
	MarginForUnfilled float64 `json:"margin_for_unfilled,string"`
	MarginFrozen      float64 `json:"margin_frozen,string"`
	RealizedPnl       float64 `json:"realized_pnl,string"`
	UnrealizedPnl     float64 `json:"unrealizedPnl,string"`
}

type FuturesCrossAccount struct {
	Equity            float64 `json:"equity,string"`
	Margin            float64 `json:"margin,string"`
	MarginMode        string  `json:"margin_mode"`
	MarginRatio       float64 `json:"margin_ratio,string"`
	RealizedPnl       float64 `json:"realized_pnl,string"`
	UnrealizedPnl     float64 `json:"unrealized_pnl,string"`
	TotalAvailBalance float64 `json:"total_avail_balance,string"`
}

type FuturesCurrencyAccount struct {
	TotalAvailBalance float64                        `json:"total_avail_balance,string"` // 账户余额（账户静态权益）
	Contracts         []FuturesFixedAccountContracts `json:"contracts"`
	Equity            float64                        `json:"equity,string"`         // 账户权益（账户动态权益）
	MarginMode        string                         `json:"margin_mode"`           // 账户类型 全仓：crossed 逐仓: fixed
	AutoMargin        int                            `json:"auto_margin,string"`    // 是否自动追加保证金 1: 自动追加已开启 0: 自动追加未开启
	LiquiMode         string                         `json:"liqui_mode"`            // 强平模式：tier（梯度强平）
	CanWithdraw       float64                        `json:"can_withdraw,string"`   // 可划转数量
	RealizedPnl       float64                        `json:"realized_pnl,string"`   // 全仓模式 已实现盈亏
	UnRealizedPnl     float64                        `json:"unrealized_pnl,string"` // 全仓模式 未实现盈亏
	Margin            float64                        `json:"margin,string"`         // 保证金（挂单冻结+持仓已用）
}

type FuturesCurrencyAccountV0 struct {
	BizWarmTips
	Result
	MarginMode   string
	CrossAccount FuturesCrossAccount
	FixedAccount FuturesFixedAccount
}

type FuturesCurrencyLedger struct {
	LedgerId  int64                        `json:"ledger_id,string"`
	Amount    float64                      `json:"amount,string"`
	Balance   float64                      `json:"balance,string"`
	Currency  string                       `json:"currency"`
	Type      string                       `json:"type"`
	Timestamp string                       `json:"timestamp"`
	Details   FuturesCurrencyLedgerDetails `json:"details"`
}

type FuturesCurrencyLedgerDetails struct {
	OrderId      string `json:"order_id"`
	InstrumentId string `json:"instrument_id"`
}

type FuturesAccountsHolds struct {
	InstrumentId string  `json:"instrument_id"`
	Amount       float64 `json:"amount,string"`
	Timestamp    string  `json:"timestamp"`
}

type FuturesNewOrderResult struct {
	BizWarmTips
	Result
	ClientOid string `json:"client_oid"`
	OrderId   string `json:"order_id"`
}

type FuturesBatchNewOrderResult struct {
	Result
	OrderInfo []OrderInfo `json:"order_info"`
}

type CodeMessage struct {
	ErrorCode    int64  `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}

/*
  If OrderId = -1, ErrorCode > 0, error order
*/
type OrderInfo struct {
	ClientOid string `json:"client_oid"`
	OrderId   string `json:"order_id"`
	CodeMessage
}

type FuturesCancelInstrumentOrderResult struct {
	Result
	ErrorCode    int    `json:"error_code,string"`
	ErrorMessage string `json:"error_message"`
	OrderId      string `json:"order_id"`
	ClientOid    string `json:"client_oid"`
	InstrumentId string `json:"instrument_id"`
}

type FuturesBatchCancelInstrumentOrdersResult struct {
	Result
	OrderIds     []string `json:"order_ids"`
	InstrumentId string   `json:"instrument_id"`
}

type FuturesClosePositionResult struct {
	Result
	ClosePositionInfo []ClosePositionInfo `json:"close_position_info"`
}

type ClosePositionInfo struct {
	InstrumentId string `json:"instrument_id"`
	CodeMessage
}

type FuturesGetOrdersResult struct {
	Result
	Orders []FuturesGetOrderResult `json:"order_info"`
}

type FuturesGetOrderResult struct {
	InstrumentId string  `json:"instrument_id"`
	Size         int64   `json:"size,string"`
	Timestamp    string  `json:"timestamp"`
	FilledQty    float64 `json:"filled_qty,string"`
	Fee          float64 `json:"fee,string"`
	OrderId      string  `json:"order_id"`
	ClientOId    string  `json:"client_oid"`
	Price        float64 `json:"price,string"`
	PriceAvg     float64 `json:"price_avg,string"`
	Status       string  `json:"status"`
	State        int     `json:"state,string"`
	Type         int     `json:"type,string"`
	OrderType    int     `json:"order_type,string"`
	ContractVal  float64 `json:"contract_val,string"`
	Leverage     float64 `json:"leverage,string"`
}

type FuturesFillResult struct {
	TradeId      int64   `json:"trade_id,string"`
	InstrumentId string  `json:"instrument_id"`
	Price        float64 `json:"price,string"`
	OrderQty     float64 `json:"order_qty,string"`
	OrderId      string  `json:"order_id"`
	CreatedAt    string  `json:"created_at"`
	ExecType     string  `json:"exec_type"`
	Fee          float64 `json:"fee,string"`
	Side         string  `json:"side"`
}

type FuturesUsersSelfTrailingVolumesResult struct {
	FuturesUsersSelfTrailingVolumeResult []FuturesUsersSelfTrailingVolumeResult
}

type FuturesUsersSelfTrailingVolumeResult struct {
	InstrumentId   string  `json:"instrument_id"`
	ExchangeVolume float64 `json:"exchange_volume,string"`
	Volume         float64 `json:"volume,string"`
	RecordedAt     string  `json:"recorded_at"`
}

type FuturesInstrumentsResult struct {
	InstrumentId        string  `json:"instrument_id"`
	UnderlyingIndex     string  `json:"underlying_index"`
	QuoteCurrency       string  `json:"quote_currency"`
	TickSize            float64 `json:"tick_size,string"`
	ContractVal         float64 `json:"contract_val,string"`
	Listing             string  `json:"listing"`
	Delivery            string  `json:"delivery"`
	TradeIncrement      float64 `json:"trade_increment,string"`
	Alias               string  `json:"alias"`
	Underlying          string  `json:"underlying"`
	BaseCurrency        string  `json:"base_currency"`
	SettlementCurrency  string  `json:"settlement_currency"`
	IsInverse           bool    `json:"is_inverse,string"`
	ContractValCurrency string  `json:"contract_val_currency"`
}

type FuturesInstrumentCurrenciesResult struct {
	Id      int64   `json:"id,string"`
	Name    string  `json:"name"`
	MinSize float64 `json:"min_size,string"`
}

type FuturesInstrumentBookResult struct {
	Asks      [][]string `json:"asks"`
	Bids      [][]string `json:"bids"`
	Timestamp string     `json:"timestamp"`
}

type FuturesInstrumentTickerResult struct {
	InstrumentId string  `json:"instrument_id"`
	BestBid      float64 `json:"best_bid,string"`
	BestAsk      float64 `json:"best_ask,string"`
	High24h      float64 `json:"high_24h,string"`
	Low24h       float64 `json:"low_24h,string"`
	Last         float64 `json:"last,string"`
	Volume24h    float64 `json:"volume_24h,string"`
	Timestamp    string  `json:"timestamp"`
}

type FuturesInstrumentTradesResult struct {
	TradeId   string  `json:"trade_id"`
	Side      string  `json:"side"`
	Price     float64 `json:"price,string"`
	Qty       float64 `json:"qty,string"`
	Timestamp string  `json:"timestamp"`
}

type FuturesInstrumentIndexResult struct {
	InstrumentId string  `json:"instrument_id"`
	Index        float64 `json:"index,string"`
	Timestamp    string  `json:"timestamp"`
}

type FuturesInstrumentEstimatedPriceResult struct {
	InstrumentId    string  `json:"instrument_id"`
	SettlementPrice float64 `json:"settlement_price,string"`
	Timestamp       string  `json:"timestamp"`
}

type FuturesInstrumentOpenInterestResult struct {
	InstrumentId string `json:"instrument_id"`
	Amount       int64  `json:"amount,string"`
	Timestamp    string `json:"timestamp"`
}

type FuturesInstrumentPriceLimitResult struct {
	InstrumentId string  `json:"instrument_id"`
	Highest      float64 `json:"highest,string"`
	Lowest       float64 `json:"lowest,string"`
	Timestamp    string  `json:"timestamp"`
}

type FuturesInstrumentLiquidationListResult struct {
	Page            PageResult
	LiquidationList []FuturesInstrumentLiquidationResult
}

type FuturesInstrumentLiquidationResult struct {
	InstrumentId string  `json:"instrument_id"`
	Price        float64 `json:"price"`
	Size         int64   `json:"size"`
	Loss         float64 `json:"loss"`
	CreatedAt    string  `json:"created_at"`
}
