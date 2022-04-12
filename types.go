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

// CoinsID https://api.coingecko.com/api/v3/coins/bitcoin
type CoinsID struct {
	ID                           string              `json:"id"`
	Symbol                       string              `json:"symbol"`
	Name                         string              `json:"name"`
	BlockTimeInMin               int32               `json:"block_time_in_minutes"`
	HashingAlgorithm             string              `json:"hashing_algorithm"`
	Categories                   []string            `json:"categories"`
	Localization                 *LocalizationItem   `json:"localization,omitempty"`
	Description                  DescriptionItem     `json:"description"`
	Links                        *LinksItem          `json:"links"`
	Image                        ImageItem           `json:"image"`
	CountryOrigin                string              `json:"country_origin"`
	GenesisDate                  string              `json:"genesis_date"`
	SentimentVotesUpPercentage   float64             `json:"sentiment_votes_up_percentage"`
	SentimentVotesDownPercentage float64             `json:"sentiment_votes_down_percentage"`
	MarketCapRank                uint16              `json:"market_cap_rank"`
	CoinGeckoRank                uint16              `json:"coingecko_rank"`
	CoinGeckoScore               float32             `json:"coingecko_score"`
	DeveloperScore               float32             `json:"developer_score"`
	CommunityScore               float32             `json:"community_score"`
	LiquidityScore               float32             `json:"liquidity_score"`
	PublicInterestScore          float32             `json:"public_interest_score"`
	MarketData                   *MarketDataItem     `json:"market_data"`
	CommunityData                *CommunityDataItem  `json:"community_data,omitempty"`
	DeveloperData                *DeveloperDataItem  `json:"developer_data,omitempty"`
	PublicInterestStats          *PublicInterestItem `json:"public_interest_stats"`
	StatusUpdates                *[]StatusUpdateItem `json:"status_updates,omitempty"`
	LastUpdated                  time.Time           `json:"last_updated"`
	Tickers                      *[]TickerItem       `json:"tickers,omitempty"`
}

// CoinsIDTickers https://api.coingecko.com/api/v3/coins/steem/tickers?page=1
type CoinsIDTickers struct {
	Name    string       `json:"name"`
	Tickers []TickerItem `json:"tickers"`
}

// CoinsIDHistory https://api.coingecko.com/api/v3/coins/steem/history?date=30-12-2018
type CoinsIDHistory struct {
	CoinBaseStruct
	Localization   LocalizationItem    `json:"localization"`
	Image          ImageItem           `json:"image"`
	MarketData     *MarketDataItem     `json:"market_data"`
	CommunityData  *CommunityDataItem  `json:"community_data"`
	DeveloperData  *DeveloperDataItem  `json:"developer_data"`
	PublicInterest *PublicInterestItem `json:"public_interest_stats"`
}

// CoinsIDMarketChart https://api.coingecko.com/api/v3/coins/bitcoin/market_chart?vs_currency=usd&days=1
type CoinsIDMarketChartRequest struct {
	ID         string `json:"id"`
	VsCurrency string `json:"vs_currency"`
	Days       string `json:"days"`
	Interval   string `json:"interval"`
}

// ChartItem
//type ChartItem [2]float32

type CoinsIDMarketChart struct {
	Prices       [][2]float64 `json:"prices,omitempty"`
	MarketCaps   [][2]float64 `json:"market_caps,omitempty"`
	TotalVolumes [][2]float64 `json:"total_volumes,omitempty"`
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

//type CategoriesListResponse []CategoriesListItem

type CategoriesListItem struct {
	ID   string `json:"category_id"`
	Name string `json:"name"`
}

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

// OrderType in CoinGecko
type OrderType struct {
	MarketCapDesc string
	MarketCapAsc  string
	GeckoDesc     string
	GeckoAsc      string
	VolumeAsc     string
	VolumeDesc    string
}

// OrderTypeObject for certain order
var OrderTypeObject = &OrderType{
	MarketCapDesc: "market_cap_desc",
	MarketCapAsc:  "market_cap_asc",
	GeckoDesc:     "gecko_desc",
	GeckoAsc:      "gecko_asc",
	VolumeAsc:     "volume_asc",
	VolumeDesc:    "volume_desc",
}

// PriceChangePercentage

// PriceChangePercentage in different amount of time
type PriceChangePercentage struct {
	PCP1h   string
	PCP24h  string
	PCP7d   string
	PCP14d  string
	PCP30d  string
	PCP200d string
	PCP1y   string
}

// PriceChangePercentageObject for different amount of time
var PriceChangePercentageObject = &PriceChangePercentage{
	PCP1h:   "1h",
	PCP24h:  "24h",
	PCP7d:   "7d",
	PCP14d:  "14d",
	PCP30d:  "30d",
	PCP200d: "200d",
	PCP1y:   "1y",
}

// CoinBaseStruct [private]
type CoinBaseStruct struct {
	ID     string `json:"id"`
	Symbol string `json:"symbol"`
	Name   string `json:"name"`
}

// AllCurrencies map all currencies (USD, BTC) to float64
type AllCurrencies map[string]float64

// LocalizationItem map all locale (en, zh) into respective string
type LocalizationItem map[string]string

// TYPES
// DescriptionItem map all description (in locale) into respective string
type DescriptionItem map[string]string

// LinksItem map all links
type LinksItem map[string]interface{}

// MarketDataItem map all market data item
type MarketDataItem struct {
	CurrentPrice                           AllCurrencies     `json:"current_price"`
	ROI                                    *ROIItem          `json:"roi"`
	ATH                                    AllCurrencies     `json:"ath"`
	ATHChangePercentage                    AllCurrencies     `json:"ath_change_percentage"`
	ATHDate                                map[string]string `json:"ath_date"`
	ATL                                    AllCurrencies     `json:"atl"`
	ATLChangePercentage                    AllCurrencies     `json:"atl_change_percentage"`
	ATLDate                                map[string]string `json:"atl_date"`
	MarketCap                              AllCurrencies     `json:"market_cap"`
	MarketCapRank                          uint16            `json:"market_cap_rank"`
	TotalVolume                            AllCurrencies     `json:"total_volume"`
	High24                                 AllCurrencies     `json:"high_24h"`
	Low24                                  AllCurrencies     `json:"low_24h"`
	PriceChange24h                         float64           `json:"price_change_24h"`
	PriceChangePercentage24h               float64           `json:"price_change_percentage_24h"`
	PriceChangePercentage7d                float64           `json:"price_change_percentage_7d"`
	PriceChangePercentage14d               float64           `json:"price_change_percentage_14d"`
	PriceChangePercentage30d               float64           `json:"price_change_percentage_30d"`
	PriceChangePercentage60d               float64           `json:"price_change_percentage_60d"`
	PriceChangePercentage200d              float64           `json:"price_change_percentage_200d"`
	PriceChangePercentage1y                float64           `json:"price_change_percentage_1y"`
	MarketCapChange24h                     float64           `json:"market_cap_change_24h"`
	MarketCapChangePercentage24h           float64           `json:"market_cap_change_percentage_24h"`
	PriceChange24hInCurrency               AllCurrencies     `json:"price_change_24h_in_currency"`
	PriceChangePercentage1hInCurrency      AllCurrencies     `json:"price_change_percentage_1h_in_currency"`
	PriceChangePercentage24hInCurrency     AllCurrencies     `json:"price_change_percentage_24h_in_currency"`
	PriceChangePercentage7dInCurrency      AllCurrencies     `json:"price_change_percentage_7d_in_currency"`
	PriceChangePercentage14dInCurrency     AllCurrencies     `json:"price_change_percentage_14d_in_currency"`
	PriceChangePercentage30dInCurrency     AllCurrencies     `json:"price_change_percentage_30d_in_currency"`
	PriceChangePercentage60dInCurrency     AllCurrencies     `json:"price_change_percentage_60d_in_currency"`
	PriceChangePercentage200dInCurrency    AllCurrencies     `json:"price_change_percentage_200d_in_currency"`
	PriceChangePercentage1yInCurrency      AllCurrencies     `json:"price_change_percentage_1y_in_currency"`
	MarketCapChange24hInCurrency           AllCurrencies     `json:"market_cap_change_24h_in_currency"`
	MarketCapChangePercentage24hInCurrency AllCurrencies     `json:"market_cap_change_percentage_24h_in_currency"`
	TotalSupply                            *float64          `json:"total_supply"`
	CirculatingSupply                      float64           `json:"circulating_supply"`
	Sparkline                              *SparklineItem    `json:"sparkline_7d"`
	LastUpdated                            string            `json:"last_updated"`
}

// CommunityDataItem map all community data item
type CommunityDataItem struct {
	FacebookLikes            *uint    `json:"facebook_likes,omitempty"`
	TwitterFollowers         *uint    `json:"twitter_followers,omitempty"`
	RedditAveragePosts48h    *float64 `json:"reddit_average_posts_48h,omitempty"`
	RedditAverageComments48h *float64 `json:"reddit_average_comments_48h,omitempty"`
	RedditSubscribers        *uint    `json:"reddit_subscribers,omitempty"`
	RedditAccountsActive48h  *uint    `json:"reddit_accounts_active_48h,omitempty"`
	TelegramChannelUserCount *uint    `json:"telegram_channel_user_count,omitempty"`
}

// DeveloperDataItem map all developer data item
type DeveloperDataItem struct {
	Forks              *uint `json:"forks"`
	Stars              *uint `json:"stars"`
	Subscribers        *uint `json:"subscribers"`
	TotalIssues        *uint `json:"total_issues"`
	ClosedIssues       *uint `json:"closed_issues"`
	PRMerged           *uint `json:"pull_requests_merged"`
	PRContributors     *uint `json:"pull_request_contributors"`
	CommitsCount4Weeks *uint `json:"commit_count_4_weeks"`
}

// PublicInterestItem map all public interest item
type PublicInterestItem struct {
	AlexaRank   *uint `json:"alexa_rank,omitempty"`
	BingMatches *uint `json:"bing_matches,omitempty"`
}

// ImageItem struct for all sizes of image
type ImageItem struct {
	Thumb string `json:"thumb"`
	Small string `json:"small"`
	Large string `json:"large"`
}

// ROIItem ROI Item
type ROIItem struct {
	Times      float64 `json:"times"`
	Currency   string  `json:"currency"`
	Percentage float64 `json:"percentage"`
}

// SparklineItem for sparkline
type SparklineItem struct {
	Price []float64 `json:"price"`
}

// TickerItem for ticker
type TickerItem struct {
	Base   string `json:"base"`
	Target string `json:"target"`
	Market struct {
		Name             string `json:"name"`
		Identifier       string `json:"identifier"`
		TradingIncentive bool   `json:"has_trading_incentive"`
	} `json:"market"`
	Last   float64 `json:"last"`
	Volume float64 `json:"volume"`

	//ConvertedLast   map[string]float64 `json:"converted_last"`
	ConvertedLast struct {
		BTC float64 `json:"btc"`
		ETH float64 `json:"eth"`
		USD float64 `json:"usd"`
	} `json:"converted_last"`
	//ConvertedVolume map[string]float64 `json:"converted_volume"`
	ConvertedVolume struct {
		BTC float64 `json:"btc"`
		ETH float64 `json:"eth"`
		USD float64 `json:"usd"`
	} `json:"converted_volume"`
	TrustScore             string    `json:"trust_score"`
	BidAskSpreadPercentage float64   `json:"bid_ask_spread_percentage"`
	LastTradedAt           time.Time `json:"last_traded_at"`
	LastFetchAt            time.Time `json:"last_fetch_at"`

	Timestamp    time.Time `json:"timestamp"`
	IsAnomaly    bool      `json:"is_anomaly"`
	IsStale      bool      `json:"is_stale"`
	CoinID       string    `json:"coin_id"`
	TargetCoinId string    `json:"target_coin_id"`
}

// StatusUpdateItem for BEAM
type StatusUpdateItem struct {
	Description string `json:"description"`
	Category    string `json:"category"`
	CreatedAt   string `json:"created_at"`
	User        string `json:"user"`
	UserTitle   string `json:"user_title"`
	Pin         bool   `json:"pin"`
	Project     struct {
		CoinBaseStruct
		Type  string    `json:"type"`
		Image ImageItem `json:"image"`
	} `json:"project"`
}

type CoinsMarketRequest struct {
	VsCurrency            string
	Ids                   []string
	Order                 string
	PerPage               int
	Page                  int
	Sparkline             bool
	PriceChangePercentage []string
}

// CoinsMarketItem item in CoinMarket
type CoinsMarketItem struct {
	ID                    string  `json:"id"`
	Symbol                string  `json:"symbol"`
	Name                  string  `json:"name"`
	Image                 string  `json:"image"`
	CurrentPrice          float64 `json:"current_price"`
	MarketCap             float64 `json:"market_cap"`
	MarketCapRank         int64   `json:"market_cap_rank"`
	FullyDilutedValuation *int64  `json:"fully_diluted_valuation,omitempty"`

	TotalVolume              float64 `json:"total_volume"`
	High24                   float64 `json:"high_24h"`
	Low24                    float64 `json:"low_24h"`
	PriceChange24h           float64 `json:"price_change_24h"`
	PriceChangePercentage24h float64 `json:"price_change_percentage_24h"`

	MarketCapChange24h           float64 `json:"market_cap_change_24h"`
	MarketCapChangePercentage24h float64 `json:"market_cap_change_percentage_24h"`

	CirculatingSupply *float64 `json:"circulating_supply,omitempty"`
	TotalSupply       *float64 `json:"total_supply,omitempty"`
	MaxSupply         *float64 `json:"max_supply,omitempty"`

	ATH float64 `json:"ath"`
	//the drop in percents of the price of a cryptocurrency compared to its maximum price (ATH) of all time
	ATHChangePercentage float64   `json:"ath_change_percentage"`
	ATHDate             time.Time `json:"ath_date"`
	ATL                 float64   `json:"atl"`
	ATLChangePercentage float64   `json:"atl_change_percentage"`
	ATLDate             time.Time `json:"atl_date"`

	LastUpdated                         time.Time      `json:"last_updated"`
	PriceChangePercentage1hInCurrency   *float64       `json:"price_change_percentage_1h_in_currency,omitempty"`
	PriceChangePercentage24hInCurrency  *float64       `json:"price_change_percentage_24h_in_currency,omitempty"`
	PriceChangePercentage7dInCurrency   *float64       `json:"price_change_percentage_7d_in_currency,omitempty"`
	PriceChangePercentage14dInCurrency  *float64       `json:"price_change_percentage_14d_in_currency,omitempty"`
	PriceChangePercentage30dInCurrency  *float64       `json:"price_change_percentage_30d_in_currency,omitempty"`
	PriceChangePercentage200dInCurrency *float64       `json:"price_change_percentage_200d_in_currency,omitempty"`
	PriceChangePercentage1yInCurrency   *float64       `json:"price_change_percentage_1y_in_currency,omitempty"`
	ROI                                 *ROIItem       `json:"roi,omitempty"`
	SparklineIn7d                       *SparklineItem `json:"sparkline_in_7d,omitempty"`
}

// EventCountryItem item in EventsCountries
type EventCountryItem struct {
	Country string `json:"country"`
	Code    string `json:"code"`
}

// ExchangeRatesItem item in ExchangeRate
type ExchangeRatesItem map[string]ExchangeRatesItemStruct

// ExchangeRatesItemStruct struct in ExchangeRateItem
type ExchangeRatesItemStruct struct {
	Name  string  `json:"name"`
	Unit  string  `json:"unit"`
	Value float64 `json:"value"`
	Type  string  `json:"type"`
}

// Global for data of /global
type Global struct {
	ActiveCryptocurrencies          uint16        `json:"active_cryptocurrencies"`
	UpcomingICOs                    uint16        `json:"upcoming_icos"`
	OngoingICOs                     uint16        `json:"ongoing_icos"`
	EndedICOs                       uint16        `json:"ended_icos"`
	Markets                         uint16        `json:"markets"`
	MarketCapChangePercentage24hUSD float32       `json:"market_cap_change_percentage_24h_usd"`
	TotalMarketCap                  AllCurrencies `json:"total_market_cap"`
	TotalVolume                     AllCurrencies `json:"total_volume"`
	MarketCapPercentage             AllCurrencies `json:"market_cap_percentage"`
	UpdatedAt                       int64         `json:"updated_at"`
}
