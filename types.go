package coingecko

import (
	"time"
)

// Ping https://api.coingecko.com/api/v3/ping
type Ping struct {
	GeckoSays string `json:"gecko_says"`
}

// SimpleSinglePrice https://api.coingecko.com/api/v3/simple/price?ids=bitcoin&vs_currencies=usd
type SimpleSinglePrice struct {
	ID          string
	Currency    string
	MarketPrice float32
}

// SimpleSupportedVSCurrencies https://api.coingecko.com/api/v3/simple/supported_vs_currencies
type SimpleSupportedVSCurrencies []string

// CoinList https://api.coingecko.com/api/v3/coins/list
type CoinList []CoinsListItem

// CoinsMarket https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&order=market_cap_desc&per_page=100&page=1&sparkline=false
type CoinsMarket []CoinsMarketItem

// CoinsID https://api.coingecko.com/api/v3/coins/bitcoin
type CoinsID struct {
	coinBaseStruct
	BlockTimeInMin      int32               `json:"block_time_in_minutes"`
	Categories          []string            `json:"categories"`
	Localization        LocalizationItem    `json:"localization"`
	Description         DescriptionItem     `json:"description"`
	Links               *LinksItem          `json:"links"`
	Image               ImageItem           `json:"image"`
	CountryOrigin       string              `json:"country_origin"`
	GenesisDate         string              `json:"genesis_date"`
	MarketCapRank       uint16              `json:"market_cap_rank"`
	CoinGeckoRank       uint16              `json:"coingecko_rank"`
	CoinGeckoScore      float32             `json:"coingecko_score"`
	DeveloperScore      float32             `json:"developer_score"`
	CommunityScore      float32             `json:"community_score"`
	LiquidityScore      float32             `json:"liquidity_score"`
	PublicInterestScore float32             `json:"public_interest_score"`
	MarketData          *MarketDataItem     `json:"market_data"`
	CommunityData       *CommunityDataItem  `json:"community_data"`
	DeveloperData       *DeveloperDataItem  `json:"developer_data"`
	PublicInterestStats *PublicInterestItem `json:"public_interest_stats"`
	StatusUpdates       *[]StatusUpdateItem `json:"status_updates"`
	LastUpdated         string              `json:"last_updated"`
	Tickers             *[]TickerItem       `json:"tickers"`
}

// CoinsIDTickers https://api.coingecko.com/api/v3/coins/steem/tickers?page=1
type CoinsIDTickers struct {
	Name    string       `json:"name"`
	Tickers []TickerItem `json:"tickers"`
}

// CoinsIDHistory https://api.coingecko.com/api/v3/coins/steem/history?date=30-12-2018
type CoinsIDHistory struct {
	coinBaseStruct
	Localization   LocalizationItem    `json:"localization"`
	Image          ImageItem           `json:"image"`
	MarketData     *MarketDataItem     `json:"market_data"`
	CommunityData  *CommunityDataItem  `json:"community_data"`
	DeveloperData  *DeveloperDataItem  `json:"developer_data"`
	PublicInterest *PublicInterestItem `json:"public_interest_stats"`
}

// CoinsIDMarketChart https://api.coingecko.com/api/v3/coins/bitcoin/market_chart?vs_currency=usd&days=1
type CoinsIDMarketChart struct {
	coinBaseStruct
	Prices       *[]ChartItem `json:"prices"`
	MarketCaps   *[]ChartItem `json:"market_caps"`
	TotalVolumes *[]ChartItem `json:"total_volumes"`
}

// CoinsIDStatusUpdates

// CoinsIDContractAddress https://api.coingecko.com/api/v3/coins/{id}/contract/{contract_address}
// type CoinsIDContractAddress struct {
// 	ID                  string           `json:"id"`
// 	Symbol              string           `json:"symbol"`
// 	Name                string           `json:"name"`
// 	BlockTimeInMin      uint16           `json:"block_time_in_minutes"`
// 	Categories          []string         `json:"categories"`
// 	Localization        LocalizationItem `json:"localization"`
// 	Description         DescriptionItem  `json:"description"`
// 	Links               LinksItem        `json:"links"`
// 	Image               ImageItem        `json:"image"`
// 	CountryOrigin       string           `json:"country_origin"`
// 	GenesisDate         string           `json:"genesis_date"`
// 	ContractAddress     string           `json:"contract_address"`
// 	MarketCapRank       uint16           `json:"market_cap_rank"`
// 	CoinGeckoRank       uint16           `json:"coingecko_rank"`
// 	CoinGeckoScore      float32          `json:"coingecko_score"`
// 	DeveloperScore      float32          `json:"developer_score"`
// 	CommunityScore      float32          `json:"community_score"`
// 	LiquidityScore      float32          `json:"liquidity_score"`
// 	PublicInterestScore float32          `json:"public_interest_score"`
// 	MarketData          `json:"market_data"`
// }

// EventsCountries https://api.coingecko.com/api/v3/events/countries
type EventsCountries struct {
	Data []EventCountryItem `json:"data"`
}

// EventsTypes https://api.coingecko.com/api/v3/events/types
type EventsTypes struct {
	Data  []string `json:"data"`
	Count uint16   `json:"count"`
}

// ExchangeRatesResponse https://api.coingecko.com/api/v3/exchange_rates
type ExchangeRatesResponse struct {
	Rates ExchangeRatesItem `json:"rates"`
}

type CategoriesListResponse []CategoriesListItem

type CategoriesListItem struct {
	ID   string `json:"category_id"`
	Name string `json:"name"`
}

type CategoriesResponse []CategoriesItem
type ExchangesResponse []ExchangesItem

type CategoriesItem struct {
	ID                 string    `json:"id"`
	Name               string    `json:"name"`
	MarketCap          float64   `json:"market_cap"`
	MarketCapChange24H float64   `json:"market_cap_change_24h"`
	Content            string    `json:"content"`
	Top3Coins          []string  `json:"top_3_coins"`
	Volume24H          float64   `json:"volume_24h"`
	UpdatedAt          time.Time `json:"updated_at"`
}
type ExchangesItem struct {
	ID                          string  `json:"id"`
	Name                        string  `json:"name"`
	YearEstablished             int64   `json:"year_established"`
	Country                     string  `json:"country"`
	Description                 string  `json:"description"`
	Url                         string  `json:"url"`
	Image                       string  `json:"image"`
	HasTradingIncentive         bool    `json:"has_trading_incentive"`
	TrustScore                  int64   `json:"trust_score"`
	TrustScoreRank              int64   `json:"trust_score_rank"`
	TradeVolume24HBtc           float64 `json:"trade_volume_24h_btc"`
	TradeVolume24HBtcNormalized float64 `json:"trade_volume_24h_btc_normalized"`
}

// GlobalResponse https://api.coingecko.com/api/v3/global
type GlobalResponse struct {
	Data Global `json:"data"`
}

type Coin struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	Symbol        string `json:"symbol"`
	MarketCapRank int64  `json:"market_cap_rank"`
	Thumb         string `json:"thumb"`
	Large         string `json:"large"`
}

// GlobalResponse https://api.coingecko.com/api/v3/global
type SearchResponse struct {
	Coins      []Coin        `json:"coins"`
	Exchanges  []interface{} `json:"exchanges"`
	Icos       []interface{} `json:"icos"`
	Categories []interface{} `json:"categories"`
}
