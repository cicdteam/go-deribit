package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"text/template"
)

type data struct {
	FuncName, Args, Format, Type, Channel string
}

var subscriptions = map[string]struct {
	funcName, notificationType string
}{
	"book.{instrument_name}.{group}.{depth}.{interval}": {"BookGroup", "BookNotification"},
	"book.{instrument_name}.{interval}":                 {"BookInterval", "BookNotificationRaw"},
	"deribit_price_index.{index_name}":                  {"DeribitPriceIndex", "DeribitPriceIndexNotification"},
	"deribit_price_ranking.{index_name}":                {"DeribitPriceRanking", "DeribitPriceRankingNotification"},
	"estimated_expiration_price.{index_name}":           {"EstimatedExpirationPrice", "EstimatedExpirationPriceNotification"},
	"markprice.options.{index_name}":                    {"MarkPriceOptions", "MarkpriceOptionsNotification"},
	"perpetual.{instrument_name}.{interval}":            {"Perpetual", "PerpetualNotification"},
	"quote.{instrument_name}":                           {"Quote", "QuoteNotification"},
	"ticker.{instrument_name}.{interval}":               {"Ticker", "TickerNotification"},
	"trades.{instrument_name}.{interval}":               {"Trades", "PublicTrade"},
	"user.orders.{instrument_name}.{interval}":          {"UserOrdersInstrumentName", "Order"},
	"user.orders.{kind}.{currency}.{interval}":          {"UserOrdersKind", "Order"},
	"user.portfolio.{currency}":                         {"UserPortfolio", "UserPortfolioNotification"},
	"user.trades.{instrument_name}.{interval}":          {"UserTradesInstrument", "UserTrade"},
	"user.trades.{kind}.{currency}.{interval}":          {"UserTradesKind", "UserTrade"},
	"announcements":                                     {"Announcements", "AnnouncementNotification"},
}

func main() {
	var d data
	var t *template.Template

	fmt.Println("// Code generated by make generate-methods DO NOT EDIT.\n\npackage deribit")

	for c, params := range subscriptions {
		d.Channel = c
		d.FuncName = "Subscribe" + params.funcName
		d.Type = params.notificationType
		re := regexp.MustCompile(`\{(.*?)\}`)
		match := re.FindAllStringSubmatch(c, -1)
		if len(match) > 0 {
			args := make([]string, len(match))
			for i, m := range match {
				args[i] = m[1]
			}
			d.Args = strings.Join(args, ", ")
			d.Format = re.ReplaceAllString(c, "%s")
		} else {
			d.Format = c
		}
		t = template.Must(template.New("subMethod").Parse(subscriptionTemplate))
		err := t.Execute(os.Stdout, d)
		if err != nil {
			log.Fatal(err)
		}
	}
}

var subscriptionTemplate = `
// {{.FuncName}} subscribes to the {{.Channel}} channel
func (e *Exchange) {{.FuncName}}({{.Args}}{{if .Args}} string{{end}}) (chan *models.{{.Type}}, error) {
	chans := []string{ {{if .Args}}fmt.Sprintf("{{.Format}}", {{.Args}} ) {{else}}"{{.Format}}"{{end}} }
	
	c := make(chan *RPCNotification)
	out := make(chan *models.{{.Type}})
	sub := &RPCSubscription{Data: c, Channel: chans[0]}
	e.subscriptions[chans[0]] = sub

    client := e.Client()
	if _, err := client.GetPublicSubscribe(&operations.GetPublicSubscribeParams{Channels: chans}); err != nil {
		delete(e.subscriptions, chans[0])
		return nil, fmt.Errorf("error subscribing to channel: %s", err)
	}

	go func() {
	Loop:
		for {
			select {
			case n := <-c:
				var ret models.{{.Type}}
				if err := json.Unmarshal(n.Params.Data, &ret); err != nil {
					e.errors <- fmt.Errorf("error decoding notification: %s", err)
				}
				out <- &ret
			case <-e.stop:
				break Loop
			}
		}
	}()
	return out, nil
}
`
