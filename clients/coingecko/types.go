package coingecko

import (
	"time"
)

type Assets []Asset

type Asset struct {
	Id              string `json:"id"`
	ChainIdentifier *int   `json:"chain_identifier"`
	Name            string `json:"name"`
	Shortname       string `json:"shortname"`
	NativeCoinId    string `json:"native_coin_id"`
	Image           Image  `json:"image"`
}

type Image struct {
	Thumb *string `json:"thumb"`
	Small *string `json:"small"`
	Large *string `json:"large"`
}
type Coins []Coin

type Platforms map[string]string

type DetailPlatforms map[string]Detail

type Detail struct {
	DecimalPlace    int64  `json:"decimal_place"`
	ContractAddress string `json:"contract_address"`
}

type Coin struct {
	Id        string    `json:"id"`
	Symbol    string    `json:"symbol"`
	Name      string    `json:"name"`
	Platforms Platforms `json:"platforms"`
}

type FullCoin struct {
	Id                 string            `json:"id"`
	Symbol             string            `json:"symbol"`
	Name               string            `json:"name"`
	WebSlug            string            `json:"web_slug"`
	AssetPlatformId    string            `json:"asset_platform_id"`
	Platforms          Platforms         `json:"platforms"`
	DetailPlatforms    DetailPlatforms   `json:"detail_platforms"`
	BlockTimeInMinutes int               `json:"block_time_in_minutes"`
	HashingAlgorithm   string            `json:"hashing_algorithm"`
	Categories         []string          `json:"categories"`
	PreviewListing     bool              `json:"preview_listing"`
	PublicNotice       interface{}       `json:"public_notice"`
	AdditionalNotices  []interface{}     `json:"additional_notices"`
	Localization       map[string]string `json:"localization"`
	Description        map[string]string `json:"description"`
	Links              struct {
		Homepage                    []string            `json:"homepage"`
		Whitepaper                  string              `json:"whitepaper"`
		BlockchainSite              []string            `json:"blockchain_site"`
		OfficialForumUrl            []string            `json:"official_forum_url"`
		ChatUrl                     []string            `json:"chat_url"`
		AnnouncementUrl             []string            `json:"announcement_url"`
		TwitterScreenName           string              `json:"twitter_screen_name"`
		FacebookUsername            string              `json:"facebook_username"`
		BitcointalkThreadIdentifier int                 `json:"bitcointalk_thread_identifier"`
		TelegramChannelIdentifier   string              `json:"telegram_channel_identifier"`
		SubredditUrl                string              `json:"subreddit_url"`
		ReposUrl                    map[string][]string `json:"repos_url"`
	} `json:"links"`
	Image struct {
		Thumb string `json:"thumb"`
		Small string `json:"small"`
		Large string `json:"large"`
	} `json:"image"`
	CountryOrigin                string      `json:"country_origin"`
	GenesisDate                  interface{} `json:"genesis_date"`
	ContractAddress              string      `json:"contract_address"`
	SentimentVotesUpPercentage   interface{} `json:"sentiment_votes_up_percentage"`
	SentimentVotesDownPercentage interface{} `json:"sentiment_votes_down_percentage"`
	WatchlistPortfolioUsers      int         `json:"watchlist_portfolio_users"`
	MarketCapRank                int         `json:"market_cap_rank"`
	MarketData                   struct {
		CurrentPrice     map[string]float64 `json:"current_price"`
		TotalValueLocked map[string]float64 `json:"total_value_locked"`
		McapToTvlRatio   float64            `json:"mcap_to_tvl_ratio"`
		FdvToTvlRatio    float64            `json:"fdv_to_tvl_ratio"`
		Roi              struct {
			Times      float64 `json:"times"`
			Currency   string  `json:"currency"`
			Percentage float64 `json:"percentage"`
		} `json:"roi"`
		Ath                                    map[string]float64   `json:"ath"`
		AthChangePercentage                    map[string]float64   `json:"ath_change_percentage"`
		AthDate                                map[string]time.Time `json:"ath_date"`
		Atl                                    map[string]float64   `json:"atl"`
		AtlChangePercentage                    map[string]float64   `json:"atl_change_percentage"`
		AtlDate                                map[string]time.Time `json:"atl_date"`
		MarketCap                              map[string]float64   `json:"market_cap"`
		MarketCapRank                          int                  `json:"market_cap_rank"`
		FullyDilutedValuation                  map[string]float64   `json:"fully_diluted_valuation"`
		MarketCapFdvRatio                      float64              `json:"market_cap_fdv_ratio"`
		TotalVolume                            map[string]float64   `json:"total_volume"`
		High24H                                map[string]float64   `json:"high_24h"`
		Low24H                                 map[string]float64   `json:"low_24h"`
		PriceChange24H                         float64              `json:"price_change_24h"`
		PriceChangePercentage24H               float64              `json:"price_change_percentage_24h"`
		PriceChangePercentage7D                float64              `json:"price_change_percentage_7d"`
		PriceChangePercentage14D               float64              `json:"price_change_percentage_14d"`
		PriceChangePercentage30D               float64              `json:"price_change_percentage_30d"`
		PriceChangePercentage60D               float64              `json:"price_change_percentage_60d"`
		PriceChangePercentage200D              float64              `json:"price_change_percentage_200d"`
		PriceChangePercentage1Y                float64              `json:"price_change_percentage_1y"`
		MarketCapChange24H                     float64              `json:"market_cap_change_24h"`
		MarketCapChangePercentage24H           float64              `json:"market_cap_change_percentage_24h"`
		PriceChange24HInCurrency               map[string]float64   `json:"price_change_24h_in_currency"`
		PriceChangePercentage1HInCurrency      map[string]float64   `json:"price_change_percentage_1h_in_currency"`
		PriceChangePercentage24HInCurrency     map[string]float64   `json:"price_change_percentage_24h_in_currency"`
		PriceChangePercentage7DInCurrency      map[string]float64   `json:"price_change_percentage_7d_in_currency"`
		PriceChangePercentage14DInCurrency     map[string]float64   `json:"price_change_percentage_14d_in_currency"`
		PriceChangePercentage30DInCurrency     map[string]float64   `json:"price_change_percentage_30d_in_currency"`
		PriceChangePercentage60DInCurrency     map[string]float64   `json:"price_change_percentage_60d_in_currency"`
		PriceChangePercentage200DInCurrency    map[string]float64   `json:"price_change_percentage_200d_in_currency"`
		PriceChangePercentage1YInCurrency      map[string]float64   `json:"price_change_percentage_1y_in_currency"`
		MarketCapChange24HInCurrency           map[string]float64   `json:"market_cap_change_24h_in_currency"`
		MarketCapChangePercentage24HInCurrency map[string]float64   `json:"market_cap_change_percentage_24h_in_currency"`
		TotalSupply                            float64              `json:"total_supply"`
		MaxSupply                              float64              `json:"max_supply"`
		CirculatingSupply                      float64              `json:"circulating_supply"`
		LastUpdated                            time.Time            `json:"last_updated"`
	} `json:"market_data"`
	CommunityData struct {
		FacebookLikes            int     `json:"facebook_likes"`
		TwitterFollowers         int     `json:"twitter_followers"`
		RedditAveragePosts48H    float64 `json:"reddit_average_posts_48h"`
		RedditAverageComments48H float64 `json:"reddit_average_comments_48h"`
		RedditSubscribers        int     `json:"reddit_subscribers"`
		RedditAccountsActive48H  int     `json:"reddit_accounts_active_48h"`
		TelegramChannelUserCount int     `json:"telegram_channel_user_count"`
	} `json:"community_data"`
	DeveloperData struct {
		Forks                        int `json:"forks"`
		Stars                        int `json:"stars"`
		Subscribers                  int `json:"subscribers"`
		TotalIssues                  int `json:"total_issues"`
		ClosedIssues                 int `json:"closed_issues"`
		PullRequestsMerged           int `json:"pull_requests_merged"`
		PullRequestContributors      int `json:"pull_request_contributors"`
		CodeAdditionsDeletions4Weeks struct {
			Additions int `json:"additions"`
			Deletions int `json:"deletions"`
		} `json:"code_additions_deletions_4_weeks"`
		CommitCount4Weeks              int           `json:"commit_count_4_weeks"`
		Last4WeeksCommitActivitySeries []interface{} `json:"last_4_weeks_commit_activity_series"`
	} `json:"developer_data"`
	StatusUpdates []interface{} `json:"status_updates"`
	LastUpdated   time.Time     `json:"last_updated"`
	Tickers       []struct {
		Base   string `json:"base"`
		Target string `json:"target"`
		Market struct {
			Name                string `json:"name"`
			Identifier          string `json:"identifier"`
			HasTradingIncentive bool   `json:"has_trading_incentive"`
		} `json:"market"`
		Last                   float64            `json:"last"`
		Volume                 float64            `json:"volume"`
		ConvertedLast          map[string]float64 `json:"converted_last"`
		ConvertedVolume        map[string]float64 `json:"converted_volume"`
		TrustScore             *string            `json:"trust_score"`
		BidAskSpreadPercentage float64            `json:"bid_ask_spread_percentage"`
		Timestamp              time.Time          `json:"timestamp"`
		LastTradedAt           time.Time          `json:"last_traded_at"`
		LastFetchAt            time.Time          `json:"last_fetch_at"`
		IsAnomaly              bool               `json:"is_anomaly"`
		IsStale                bool               `json:"is_stale"`
		TradeUrl               string             `json:"trade_url"`
		TokenInfoUrl           *string            `json:"token_info_url"`
		CoinId                 string             `json:"coin_id"`
		TargetCoinId           string             `json:"target_coin_id"`
	} `json:"tickers"`
}

type Price map[string]float64

type ContractPrice map[string]Price
