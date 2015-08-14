package gobet

import (
	"time"
)

// custom struct that encompasses all possible API input values
type AccountsInputParams struct {
	AppName            string             `json:"appName,omitempty"`
	From               Wallet             `json:"from,omitempty"`
	To                 Wallet             `json:"to,omitempty"`
	Amount             float64            `json:"amount,omitempty"`
	SubscriptionLength int                `json:"subscriptionLength,omitempty"`
	SubscriptionToken  string             `json:"subscriptionToken,omitempty"`
	VendorClientId     string             `json:"vendorClientId,omitempty"`
	SubscriptionStatus SubscriptionStatus `json:"subscriptionStatus,omitempty"`
	Locale             string             `json:"locale,omitempty"`
	FromRecord         int                `json:"fromRecord,omitempty"`
	RecordCount        int                `json:"recordCount,omitempty"`
	ItemDateRange      *TimeRange         `json:"itemDateRange,omitempty"`
	IncludeItem        *IncludeItem       `json:"includeItem,omitempty"`
	Wallet             Wallet             `json:"wallet,omitempty"`
	FromCurrency       string             `json:"fromCurrency,omitempty"`
}

type TransferResponse struct {
	TransactionId string `json:"transactionId,omitempty"`
}

type ApplicationSubscription struct {
	SubscriptionToken    string             `json:"subscriptionToken,omitempty"`
	ExpiryDateTime       time.Time          `json:"expiryDateTime,omitempty"`
	ExpiredDateTime      time.Time          `json:"expiredDateTime,omitempty"`
	CreatedDateTime      time.Time          `json:"createdDateTime,omitempty"`
	ActivationDateTime   time.Time          `json:"activationDateTime,omitempty"`
	CancellationDateTime time.Time          `json:"cancellationDateTime,omitempty"`
	SubscriptionStatus   SubscriptionStatus `json:"subscriptionStatus,omitempty"`
	ClientReference      string             `json:"clientReference,omitempty"`
	VendorClientId       string             `json:"vendorClientId,omitempty"`
}

type SubscriptionHistory struct {
	SubscriptionToken    string             `json:"subscriptionToken,omitempty"`
	ExpiryDateTime       time.Time          `json:"expiryDateTime,omitempty"`
	ExpiredDateTime      time.Time          `json:"expiredDateTime,omitempty"`
	CreatedDateTime      time.Time          `json:"createdDateTime,omitempty"`
	ActivationDateTime   time.Time          `json:"activationDateTime,omitempty"`
	CancellationDateTime time.Time          `json:"cancellationDateTime,omitempty"`
	SubscriptionStatus   SubscriptionStatus `json:"subscriptionStatus,omitempty"`
	ClientReference      string             `json:"clientReference,omitempty"`
}

type AccountSubscription struct {
	SubscriptionTokens   []SubscriptionTokenInfo `json:"subscriptionTokens,omitempty"`
	ApplicationName      string                  `json:"applicationName,omitempty"`
	ApplicationVersionId string                  `json:"applicationVersionId,omitempty"`
}

type SubscriptionTokenInfo struct {
	SubscriptionToken    string             `json:"subscriptionToken,omitempty"`
	ActivatedDateTime    time.Time          `json:"activatedDateTime,omitempty"`
	ExpiryDateTime       time.Time          `json:"expiryDateTime,omitempty"`
	ExpiredDateTime      time.Time          `json:"expiredDateTime,omitempty"`
	CancellationDateTime time.Time          `json:"cancellationDateTime,omitempty"`
	SubscriptionStatus   SubscriptionStatus `json:"subscriptionStatus,omitempty"`
}

type DeveloperApp struct {
	AppName     string                `json:"appName,omitempty"`
	AppId       int64                 `json:"appId,omitempty"`
	AppVersions []DeveloperAppVersion `json:"appVersions,omitempty"`
}

type DeveloperAppVersion struct {
	Owner                string `json:"owner,omitempty"`
	VersionId            int64  `json:"versionId,omitempty"`
	Version              string `json:"version,omitempty"`
	ApplicationKey       string `json:"applicationKey,omitempty"`
	DelayData            bool   `json:"delayData,omitempty"`
	SubscriptionRequired bool   `json:"subscriptionRequired,omitempty"`
	OwnerManaged         bool   `json:"ownerManaged,omitempty"`
	Active               bool   `json:"active,omitempty"`
}

type AccountFundsResponse struct {
	AvailableToBetBalance float64 `json:"availableToBetBalance,omitempty"`
	Exposure              float64 `json:"exposure,omitempty"`
	RetainedCommission    float64 `json:"retainedCommission,omitempty"`
	ExposureLimit         float64 `json:"exposureLimit,omitempty"`
	DiscountRate          float64 `json:"discountRate,omitempty"`
	PointsBalance         int     `json:"pointsBalance,omitempty"`
}

type AccountDetailsResponse struct {
	CurrencyCode  string  `json:"currencyCode,omitempty"`
	FirstName     string  `json:"firstName,omitempty"`
	LastName      string  `json:"lastName,omitempty"`
	LocaleCode    string  `json:"localeCode,omitempty"`
	Region        string  `json:"region,omitempty"`
	Timezone      string  `json:"timezone,omitempty"`
	DiscountRate  float64 `json:"discountRate,omitempty"`
	PointsBalance int     `json:"pointsBalance,omitempty"`
	CountryCode   string  `json:"countryCode,omitempty"`
}

type AccountStatementReport struct {
	AccountStatement []StatementItem `json:"accountStatement,omitempty"`
	MoreAvailable    bool            `json:"moreAvailable,omitempty"`
}

type StatementItem struct {
	RefId         string              `json:"refId,omitempty"`
	ItemDate      time.Time           `json:"itemDate,omitempty"`
	Amount        float64             `json:"amount,omitempty"`
	Balance       float64             `json:"balance,omitempty"`
	ItemClass     ItemClass           `json:"itemClass,omitempty"`
	ItemClassData map[string]string   `json:"itemClassData,omitempty"`
	LegacyData    StatementLegacyData `json:"legacyData,omitempty"`
}

type StatementLegacyData struct {
	AvgPrice        float64   `json:"avgPrice,omitempty"`
	BetSize         float64   `json:"betSize,omitempty"`
	BetType         string    `json:"betType,omitempty"`
	BetCategoryType string    `json:"betCategoryType,omitempty"`
	CommissionRate  string    `json:"commissionRate,omitempty"`
	EventId         int64     `json:"eventId,omitempty"`
	EventTypeId     int64     `json:"eventTypeId,omitempty"`
	FullMarketName  string    `json:"fullMarketName,omitempty"`
	GrossBetAmount  float64   `json:"grossBetAmount,omitempty"`
	MarketName      string    `json:"marketName,omitempty"`
	MarketType      string    `json:"marketType,omitempty"`
	PlacedDate      time.Time `json:"placedDate,omitempty"`
	SelectionId     int64     `json:"selectionId,omitempty"`
	SelectionName   string    `json:"selectionName,omitempty"`
	StartDate       time.Time `json:"startDate,omitempty"`
	TransactionType string    `json:"transactionType,omitempty"`
	TransactionId   int64     `json:"transactionId,omitempty"`
	WinLose         string    `json:"winLose,omitempty"`
}

// below exists from Betting_type_definitions
//type TimeRange

type CurrencyRate struct {
	CurrencyCode string  `json:"currencyCode,omitempty"`
	Rate         float64 `json:"rate,omitempty"`
}
