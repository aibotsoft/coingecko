package coingecko

import (
	"context"
	"github.com/davecgh/go-spew/spew"
	"testing"
)

var c = NewClient(Config{
	BaseUrl:    "",
	Debug:      false,
	HttpClient: nil,
})

func TestPing(t *testing.T) {
	got, err := c.Ping(context.Background())
	if err != nil {
		t.FailNow()
	}
	if got.GeckoSays != "(V3) To the Moon!" {
		t.FailNow()
	}
	spew.Dump(got)
	t.Log(got)
}

func TestSimpleSinglePrice(t *testing.T) {
	simplePrice, err := c.SimpleSinglePrice(context.Background(), "bitcoin", "usd")
	if err != nil {
		t.FailNow()
	}
	t.Log(simplePrice)
	if simplePrice.ID != "bitcoin" || simplePrice.Currency != "usd" || simplePrice.MarketPrice != float32(5013.61) {
		t.FailNow()
	}
}
func TestSimplePrice(t *testing.T) {
	ids := []string{"bitcoin", "ethereum"}
	vc := []string{"usd", "myr"}
	sp, err := c.SimplePrice(context.Background(), ids, vc)
	if err != nil {
		t.FailNow()
	}
	t.Log(sp)
}

func TestSimpleSupportedVSCurrencies(t *testing.T) {
	s, err := c.SimpleSupportedVSCurrencies(context.Background())
	if err != nil {
		t.FailNow()
	}
	t.Log(s)
}

func TestCoinsList(t *testing.T) {
	list, err := c.CoinsList(context.Background())
	if err != nil {
		t.FailNow()
	}
	t.Log(list)
}

//func TestGlobal(t *testing.T) {
//	list, err := c.Global()
//	if err != nil {
//		t.FailNow()
//	}
//	t.Logf("%+v", list)
//	spew.Dump(list)
//}
//
//func TestSearch(t *testing.T) {
//	got, err := c.Search("Ethereum ")
//	if err != nil {
//		t.Log(err)
//		t.FailNow()
//	}
//	t.Logf("%+v", got)
//
//	spew.Dump(got)
//}
//func TestCategoriesList(t *testing.T) {
//	got, err := c.CategoriesList()
//	if err != nil {
//		t.Log(err)
//		t.FailNow()
//	}
//	//t.Logf("%+v", got)
//	spew.Dump(got)
//}
//func TestCategories(t *testing.T) {
//	got, err := c.Categories()
//	if err != nil {
//		t.Log(err)
//		t.FailNow()
//	}
//	//t.Logf("%+v", got)
//	spew.Dump(got)
//}
//func TestExchanges(t *testing.T) {
//	got, err := c.Exchanges()
//	if err != nil {
//		t.Log(err)
//		t.FailNow()
//	}
//	//t.Logf("%+v", got)
//	spew.Dump(got)
//}
//func TestExchangesID(t *testing.T) {
//	got, err := c.ExchangesID("binance")
//	if err != nil {
//		t.Log(err)
//		t.FailNow()
//	}
//	//t.Logf("%+v", got)
//	spew.Dump(got)
//}
