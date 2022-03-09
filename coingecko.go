package coingecko

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

var baseURL = "https://api.coingecko.com/api/v3"

// Client struct
type Client struct {
	httpClient *http.Client
}

// NewClient create new client object
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	return &Client{httpClient: httpClient}
}

// helper
// doReq HTTP client
func doReq(req *http.Request, client *http.Client) ([]byte, error) {
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if 200 != resp.StatusCode {
		return nil, fmt.Errorf("%s", body)
	}
	return body, nil
}

// MakeReq HTTP request helper
func (c *Client) MakeReq(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}
	resp, err := doReq(req, c.httpClient)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// Bool2String boolean to string
func Bool2String(b bool) string {
	return strconv.FormatBool(b)
}

// Int2String Integer to string
func Int2String(i int) string {
	return strconv.Itoa(i)
}

// API

// Ping /ping endpoint
func (c *Client) Ping() (*Ping, error) {
	url := fmt.Sprintf("%s/ping", baseURL)
	resp, err := c.MakeReq(url)
	if err != nil {
		return nil, err
	}
	var data *Ping
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// SimpleSinglePrice /simple/price  Single ID and Currency (ids, vs_currency)
func (c *Client) SimpleSinglePrice(id string, vsCurrency string) (*SimpleSinglePrice, error) {
	idParam := []string{strings.ToLower(id)}
	vcParam := []string{strings.ToLower(vsCurrency)}

	t, err := c.SimplePrice(idParam, vcParam)
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

// SimplePrice /simple/price Multiple ID and Currency (ids, vs_currencies)
func (c *Client) SimplePrice(ids []string, vsCurrencies []string) (*map[string]map[string]float32, error) {
	params := url.Values{}
	idsParam := strings.Join(ids[:], ",")
	vsCurrenciesParam := strings.Join(vsCurrencies[:], ",")

	params.Add("ids", idsParam)
	params.Add("vs_currencies", vsCurrenciesParam)

	url := fmt.Sprintf("%s/simple/price?%s", baseURL, params.Encode())
	resp, err := c.MakeReq(url)
	if err != nil {
		return nil, err
	}

	t := make(map[string]map[string]float32)
	err = json.Unmarshal(resp, &t)
	if err != nil {
		return nil, err
	}

	return &t, nil
}

// SimpleSupportedVSCurrencies /simple/supported_vs_currencies
func (c *Client) SimpleSupportedVSCurrencies() (*SimpleSupportedVSCurrencies, error) {
	url := fmt.Sprintf("%s/simple/supported_vs_currencies", baseURL)
	resp, err := c.MakeReq(url)
	if err != nil {
		return nil, err
	}
	var data *SimpleSupportedVSCurrencies
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// CoinsList /coins/list
func (c *Client) CoinsList() (*CoinList, error) {
	url := fmt.Sprintf("%s/coins/list", baseURL)
	resp, err := c.MakeReq(url)
	if err != nil {
		return nil, err
	}

	var data *CoinList
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// CoinsMarket /coins/market
func (c *Client) CoinsMarket(vsCurrency string, ids []string, order string, perPage int, page int, sparkline bool, priceChangePercentage []string) (*CoinsMarket, error) {
	if len(vsCurrency) == 0 {
		return nil, fmt.Errorf("vs_currency is required")
	}
	params := url.Values{}
	// vs_currency
	params.Add("vs_currency", vsCurrency)
	// order
	if len(order) == 0 {
		order = OrderTypeObject.MarketCapDesc
	}
	params.Add("order", order)
	// ids
	if len(ids) != 0 {
		idsParam := strings.Join(ids[:], ",")
		params.Add("ids", idsParam)
	}
	// per_page
	if perPage <= 0 || perPage > 250 {
		perPage = 100
	}
	params.Add("per_page", Int2String(perPage))
	params.Add("page", Int2String(page))
	// sparkline
	params.Add("sparkline", Bool2String(sparkline))
	// price_change_percentage
	if len(priceChangePercentage) != 0 {
		priceChangePercentageParam := strings.Join(priceChangePercentage[:], ",")
		params.Add("price_change_percentage", priceChangePercentageParam)
	}
	url := fmt.Sprintf("%s/coins/markets?%s", baseURL, params.Encode())
	resp, err := c.MakeReq(url)
	if err != nil {
		return nil, err
	}
	var data *CoinsMarket
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// CoinsID /coins/{id}
func (c *Client) CoinsID(id string, localization bool, tickers bool, marketData bool, communityData bool, developerData bool, sparkline bool) (*CoinsID, error) {

	if len(id) == 0 {
		return nil, fmt.Errorf("id is required")
	}
	params := url.Values{}
	params.Add("localization", Bool2String(sparkline))
	params.Add("tickers", Bool2String(tickers))
	params.Add("market_data", Bool2String(marketData))
	params.Add("community_data", Bool2String(communityData))
	params.Add("developer_data", Bool2String(developerData))
	params.Add("sparkline", Bool2String(sparkline))
	url := fmt.Sprintf("%s/coins/%s?%s", baseURL, id, params.Encode())
	resp, err := c.MakeReq(url)
	if err != nil {
		return nil, err
	}

	var data *CoinsID
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// CoinsIDTickers /coins/{id}/tickers
func (c *Client) CoinsIDTickers(id string, page int) (*CoinsIDTickers, error) {
	if len(id) == 0 {
		return nil, fmt.Errorf("id is required")
	}
	params := url.Values{}
	if page > 0 {
		params.Add("page", Int2String(page))
	}
	url := fmt.Sprintf("%s/coins/%s/tickers?%s", baseURL, id, params.Encode())
	resp, err := c.MakeReq(url)
	if err != nil {
		return nil, err
	}
	var data *CoinsIDTickers
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// CoinsIDHistory /coins/{id}/history?date={date}&localization=false
func (c *Client) CoinsIDHistory(id string, date string, localization bool) (*CoinsIDHistory, error) {
	if len(id) == 0 || len(date) == 0 {
		return nil, fmt.Errorf("id and date is required")
	}
	params := url.Values{}
	params.Add("date", date)
	params.Add("localization", Bool2String(localization))

	url := fmt.Sprintf("%s/coins/%s/history?%s", baseURL, id, params.Encode())
	resp, err := c.MakeReq(url)
	if err != nil {
		return nil, err
	}
	var data *CoinsIDHistory
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// CoinsIDMarketChart /coins/{id}/market_chart?vs_currency={usd, eur, jpy, etc.}&days={1,14,30,max}
func (c *Client) CoinsIDMarketChart(id string, vs_currency string, days string) (*CoinsIDMarketChart, error) {
	if len(id) == 0 || len(vs_currency) == 0 || len(days) == 0 {
		return nil, fmt.Errorf("id, vs_currency, and days is required")
	}

	params := url.Values{}
	params.Add("vs_currency", vs_currency)
	params.Add("days", days)

	url := fmt.Sprintf("%s/coins/%s/market_chart?%s", baseURL, id, params.Encode())
	resp, err := c.MakeReq(url)
	if err != nil {
		return nil, err
	}

	m := CoinsIDMarketChart{}
	err = json.Unmarshal(resp, &m)
	if err != nil {
		return &m, err
	}

	return &m, nil
}

func (c *Client) CategoriesList() (*CategoriesListResponse, error) {
	url := fmt.Sprintf("%s/coins/categories/list", baseURL)
	resp, err := c.MakeReq(url)
	if err != nil {
		return nil, err
	}

	var data *CategoriesListResponse
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
func (c *Client) Categories() (*CategoriesResponse, error) {
	url := fmt.Sprintf("%s/coins/categories", baseURL)
	resp, err := c.MakeReq(url)
	if err != nil {
		return nil, err
	}
	//fmt.Println(string(resp))
	var data *CategoriesResponse
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
func (c *Client) Exchanges() (*ExchangesResponse, error) {
	url := fmt.Sprintf("%s/exchanges", baseURL)
	resp, err := c.MakeReq(url)
	if err != nil {
		return nil, err
	}
	//fmt.Println(string(resp))
	var data *ExchangesResponse
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
func (c *Client) ExchangesID(id string) (*ExchangesResponse, error) {
	url := fmt.Sprintf("%s/exchanges/%s", baseURL, id)
	resp, err := c.MakeReq(url)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(resp))
	var data *ExchangesResponse
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// ExchangeRates https://api.coingecko.com/api/v3/exchange_rates
func (c *Client) ExchangeRates() (*ExchangeRatesItem, error) {
	url := fmt.Sprintf("%s/exchange_rates", baseURL)
	resp, err := c.MakeReq(url)
	if err != nil {
		return nil, err
	}
	var data *ExchangeRatesResponse
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return &data.Rates, nil
}

// Search https://api.coingecko.com/api/v3/search
func (c *Client) Search(query string) (*SearchResponse, error) {
	params := url.Values{}
	params.Add("query", query)
	u := fmt.Sprintf("%s/search?%s", baseURL, params.Encode())

	resp, err := c.MakeReq(u)
	if err != nil {
		return nil, err
	}
	var data *SearchResponse
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Global https://api.coingecko.com/api/v3/global
func (c *Client) Global() (*Global, error) {
	url := fmt.Sprintf("%s/global", baseURL)
	resp, err := c.MakeReq(url)
	if err != nil {
		return nil, err
	}
	var data *GlobalResponse
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return &data.Data, nil
}
