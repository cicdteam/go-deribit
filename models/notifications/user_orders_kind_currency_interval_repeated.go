package notifications

type UserOrdersKindCurrencyIntervalRepeated struct {
	Jsonrpc string `json:"jsonrpc" mapstructure:"jsonrpc"`
	Method  string `json:"method" mapstructure:"method"`
	Params  struct {
		Channel string `json:"channel" mapstructure:"channel"`
		Data    []struct {
			Amount              int64   `json:"amount" mapstructure:"amount"`
			API                 bool    `json:"api" mapstructure:"api"`
			AveragePrice        int64   `json:"average_price" mapstructure:"average_price"`
			Commission          int64   `json:"commission" mapstructure:"commission"`
			CreationTimestamp   int64   `json:"creation_timestamp" mapstructure:"creation_timestamp"`
			Direction           string  `json:"direction" mapstructure:"direction"`
			FilledAmount        int64   `json:"filled_amount" mapstructure:"filled_amount"`
			InstrumentName      string  `json:"instrument_name" mapstructure:"instrument_name"`
			IsLiquidation       bool    `json:"is_liquidation" mapstructure:"is_liquidation"`
			Label               string  `json:"label" mapstructure:"label"`
			LastUpdateTimestamp int64   `json:"last_update_timestamp" mapstructure:"last_update_timestamp"`
			MaxShow             int64   `json:"max_show" mapstructure:"max_show"`
			OrderID             string  `json:"order_id" mapstructure:"order_id"`
			OrderState          string  `json:"order_state" mapstructure:"order_state"`
			OrderType           string  `json:"order_type" mapstructure:"order_type"`
			PostOnly            bool    `json:"post_only" mapstructure:"post_only"`
			Price               float64 `json:"price" mapstructure:"price"`
			ProfitLoss          int64   `json:"profit_loss" mapstructure:"profit_loss"`
			ReduceOnly          bool    `json:"reduce_only" mapstructure:"reduce_only"`
			TimeInForce         string  `json:"time_in_force" mapstructure:"time_in_force"`
		} `json:"data" mapstructure:"data"`
	} `json:"params" mapstructure:"params"`
}