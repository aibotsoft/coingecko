package coingecko

import (
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/time/rate"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// Client struct
type Client struct {
	cfg         Config
	client      *http.Client
	rateLimiter *rate.Limiter
}
type Config struct {
	BaseUrl     string
	Debug       bool
	HttpClient  *http.Client
	RateLimiter *rate.Limiter
}

// NewClient create new client object
func NewClient(cfg Config) *Client {
	if cfg.BaseUrl == "" {
		cfg.BaseUrl = "https://api.coingecko.com/api/v3"
	}
	if cfg.RateLimiter == nil {
		//Our Free API* has a rate limit of 50 calls/minute.
		cfg.RateLimiter = rate.NewLimiter(rate.Every(time.Millisecond*800), 1)
	}
	c := &Client{cfg: cfg, rateLimiter: cfg.RateLimiter}
	if cfg.HttpClient != nil {
		c.client = cfg.HttpClient
	} else {
		t := http.DefaultTransport.(*http.Transport).Clone()
		t.MaxIdleConns = 1
		t.MaxConnsPerHost = 1
		t.MaxIdleConnsPerHost = 1
		t.IdleConnTimeout = 0
		c.client = &http.Client{Timeout: 10 * time.Second, Transport: t}
	}
	return c
}

// MakeReq HTTP request helper
func (c *Client) MakeReq(ctx context.Context, url string, data interface{}) error {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)

	if err != nil {
		return err
	}
	start := time.Now()
	err = c.rateLimiter.Wait(ctx) // This is a blocking call. Honors the rate limit
	if err != nil {
		return err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if c.cfg.Debug {
		dump, err := httputil.DumpResponse(resp, true)
		//_ = ioutil.WriteFile("test.json", dump, os.ModePerm)
		if err != nil {
			return fmt.Errorf("httputil.DumpResponse error: %w", err)
		}
		log.Printf("\nTotalTime: %s\n%s", time.Since(start), string(dump))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if 200 != resp.StatusCode {
		return fmt.Errorf("%s", body)
	}
	err = json.Unmarshal(body, data)
	return err
}

// Ping /ping endpoint
func (c *Client) Ping(ctx context.Context) (data *Ping, err error) {
	err = c.MakeReq(ctx, fmt.Sprintf("%s/ping", c.cfg.BaseUrl), &data)
	return
}

// SimplePrice /simple/price Multiple ID and Currency (ids, vs_currencies)
func (c *Client) SimplePrice(ctx context.Context, ids []string, vsCurrencies []string) (*map[string]map[string]float32, error) {
	params := url.Values{}
	idsParam := strings.Join(ids[:], ",")
	vsCurrenciesParam := strings.Join(vsCurrencies[:], ",")
	params.Add("ids", idsParam)
	params.Add("vs_currencies", vsCurrenciesParam)
	t := make(map[string]map[string]float32)
	err := c.MakeReq(ctx, fmt.Sprintf("%s/simple/price?%s", c.cfg.BaseUrl, params.Encode()), &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

// SimpleSinglePrice /simple/price  Single ID and Currency (ids, vs_currency)
func (c *Client) SimpleSinglePrice(ctx context.Context, id string, vsCurrency string) (*SimpleSinglePrice, error) {
	idParam := []string{strings.ToLower(id)}
	vcParam := []string{strings.ToLower(vsCurrency)}
	t, err := c.SimplePrice(ctx, idParam, vcParam)
	if err != nil {
		return nil, err
	}
	curr := (*t)[id]
	if len(curr) == 0 {
		return nil, fmt.Errorf("id or vsCurrency not existed")
	}
	data := &SimpleSinglePrice{ID: id, Currency: vsCurrency, MarketPrice: curr[vsCurrency]}
	return data, nil
}

// SimpleSupportedVSCurrencies /simple/supported_vs_currencies
func (c *Client) SimpleSupportedVSCurrencies(ctx context.Context) (data *SimpleSupportedVSCurrencies, err error) {
	err = c.MakeReq(ctx, fmt.Sprintf("%s/simple/supported_vs_currencies", c.cfg.BaseUrl), &data)
	return
}

// CoinsList /coins/list
func (c *Client) CoinsList(ctx context.Context) (data []CoinBaseStruct, err error) {
	err = c.MakeReq(ctx, fmt.Sprintf("%s/coins/list", c.cfg.BaseUrl), &data)
	return
}

// CoinsMarket /coins/market
func (c *Client) CoinsMarket(ctx context.Context, req CoinsMarketRequest) (data []CoinsMarketItem, err error) {
	if len(req.VsCurrency) == 0 {
		return nil, fmt.Errorf("vs_currency is required")
	}
	params := url.Values{}
	// vs_currency
	params.Add("vs_currency", req.VsCurrency)
	// order
	if len(req.Order) == 0 {
		req.Order = OrderTypeObject.MarketCapDesc
	}
	params.Add("order", req.Order)
	// ids
	if len(req.Ids) != 0 {
		idsParam := strings.Join(req.Ids[:], ",")
		params.Add("ids", idsParam)
	}
	// per_page
	if req.PerPage <= 0 || req.PerPage > 250 {
		req.PerPage = 100
	}
	params.Add("per_page", Int2String(req.PerPage))
	params.Add("page", Int2String(req.Page))
	// sparkline
	params.Add("sparkline", Bool2String(req.Sparkline))
	// price_change_percentage
	if len(req.PriceChangePercentage) != 0 {
		priceChangePercentageParam := strings.Join(req.PriceChangePercentage[:], ",")
		params.Add("price_change_percentage", priceChangePercentageParam)
	}
	err = c.MakeReq(ctx, fmt.Sprintf("%s/coins/markets?%s", c.cfg.BaseUrl, params.Encode()), &data)
	return
}

type CoinsIDRequest struct {
	ID            string `json:"id"`
	Localization  bool   `json:"localization"`
	Tickers       bool   `json:"tickers"`
	MarketData    bool   `json:"market_data"`
	CommunityData bool   `json:"community_data"`
	DeveloperData bool   `json:"developer_data"`
	Sparkline     bool   `json:"sparkline"`
}

// CoinsID /coins/{id}
func (c *Client) CoinsID(ctx context.Context, r CoinsIDRequest) (data *CoinsID, err error) {
	params := url.Values{}
	params.Add("localization", Bool2String(r.Localization))
	params.Add("tickers", Bool2String(r.Tickers))
	params.Add("market_data", Bool2String(r.MarketData))
	params.Add("community_data", Bool2String(r.CommunityData))
	params.Add("developer_data", Bool2String(r.DeveloperData))
	params.Add("sparkline", Bool2String(r.Sparkline))
	err = c.MakeReq(ctx, fmt.Sprintf("%s/coins/%s?%s", c.cfg.BaseUrl, r.ID, params.Encode()), &data)
	return data, nil
}

type CoinsIDTickersOrder string

var (
	TrustScoreDesc CoinsIDTickersOrder = "trust_score_desc"
	TrustScoreAsc  CoinsIDTickersOrder = "trust_score_asc"
	VolumeDesc     CoinsIDTickersOrder = "volume_desc"
)

type CoinsIDTickersRequest struct {
	ID          string              `json:"id"`
	ExchangeIds []string            `json:"exchange_ids"`
	Page        int                 `json:"page"`
	Order       CoinsIDTickersOrder `json:"order"`
	Depth       bool                `json:"depth"`
}

// CoinsIDTickers /coins/{id}/tickers
func (c *Client) CoinsIDTickers(ctx context.Context, r CoinsIDTickersRequest) (data *CoinsIDTickers, err error) {
	params := url.Values{}
	params.Add("page", strconv.Itoa(r.Page))
	params.Add("order", string(r.Order))
	params.Add("depth", Bool2String(r.Depth))
	//ExchangeIdsParam := strings.Join(r.ExchangeIds[:], ",")

	params.Add("exchange_ids", strings.Join(r.ExchangeIds[:], ","))
	err = c.MakeReq(ctx, fmt.Sprintf("%s/coins/%s/tickers?%s", c.cfg.BaseUrl, r.ID, params.Encode()), &data)
	return
}

// CoinsIDHistory /coins/{id}/history?date={date}&localization=false
func (c *Client) CoinsIDHistory(ctx context.Context, id string, date string, localization bool) (data *CoinsIDHistory, err error) {
	if len(id) == 0 || len(date) == 0 {
		return nil, fmt.Errorf("id and date is required")
	}
	params := url.Values{}
	params.Add("date", date)
	params.Add("localization", Bool2String(localization))
	err = c.MakeReq(ctx, fmt.Sprintf("%s/coins/%s/history?%s", c.cfg.BaseUrl, id, params.Encode()), &data)
	return
}

// CoinsIDMarketChart /coins/{id}/market_chart?vs_currency={usd, eur, jpy, etc.}&days={1,14,30,max}
func (c *Client) CoinsIDMarketChart(ctx context.Context, req CoinsIDMarketChartRequest) (data *CoinsIDMarketChart, err error) {
	if len(req.ID) == 0 || len(req.VsCurrency) == 0 || len(req.Days) == 0 {
		return nil, fmt.Errorf("id, vs_currency, and days is required")
	}
	params := url.Values{}
	params.Add("vs_currency", req.VsCurrency)
	params.Add("days", req.Days)
	params.Add("interval", req.Interval)
	err = c.MakeReq(ctx, fmt.Sprintf("%s/coins/%s/market_chart?%s", c.cfg.BaseUrl, req.ID, params.Encode()), &data)
	return
}

func (c *Client) CategoriesList(ctx context.Context) (data []CategoriesListItem, err error) {
	err = c.MakeReq(ctx, fmt.Sprintf("%s/coins/categories/list", c.cfg.BaseUrl), &data)
	return
}
func (c *Client) Categories(ctx context.Context) (data []CategoriesItem, err error) {
	err = c.MakeReq(ctx, fmt.Sprintf("%s/coins/categories", c.cfg.BaseUrl), &data)
	return
}
func (c *Client) Exchanges(ctx context.Context, perPage int, page int) (data []ExchangesItem, err error) {
	params := url.Values{}
	params.Add("per_page", strconv.Itoa(perPage))
	params.Add("page", strconv.Itoa(page))
	err = c.MakeReq(ctx, fmt.Sprintf("%s/exchanges?%s", c.cfg.BaseUrl, params.Encode()), &data)
	return
}
func (c *Client) ExchangesID(ctx context.Context, id string) (data []ExchangesItem, err error) {
	err = c.MakeReq(ctx, fmt.Sprintf("%s/exchanges/%s", c.cfg.BaseUrl, id), &data)
	return
}

// ExchangeRates https://api.coingecko.com/api/v3/exchange_rates
func (c *Client) ExchangeRates(ctx context.Context) (data *ExchangeRatesItem, err error) {
	err = c.MakeReq(ctx, fmt.Sprintf("%s/exchange_rates", c.cfg.BaseUrl), &data)
	return
}

// Search https://api.coingecko.com/api/v3/search
func (c *Client) Search(ctx context.Context, query string) (data *SearchResponse, err error) {
	params := url.Values{}
	params.Add("query", query)
	err = c.MakeReq(ctx, fmt.Sprintf("%s/search?%s", c.cfg.BaseUrl, params.Encode()), &data)
	return
}

// Global https://api.coingecko.com/api/v3/global
func (c *Client) Global(ctx context.Context) (data *Global, err error) {
	err = c.MakeReq(ctx, fmt.Sprintf("%s/global", c.cfg.BaseUrl), &data)
	return
}

// Bool2String boolean to string
func Bool2String(b bool) string {
	return strconv.FormatBool(b)
}

// Int2String Integer to string
func Int2String(i int) string {
	return strconv.Itoa(i)
}
