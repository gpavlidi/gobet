package gobet

import (
	"time"
)

// custom struct that encompasses all possible API input values
type BettingInputParams struct {
	Filter                 *MarketFilter      `json:"filter,omitempty"`
	Locale                 string             `json:"locale,omitempty"`
	MarketIds              []string           `json:"marketIds,omitempty"`
	PriceProjection        *PriceProjection   `json:"priceProjection,omitempty"`
	OrderProjection        *OrderProjection   `json:"orderProjection,omitempty"`
	MatchProjection        *MatchProjection   `json:"matchProjection,omitempty"`
	CurrencyCode           string             `json:"currencyCode,omitempty"`
	MarketProjection       []MarketProjection `json:"marketProjection,omitempty"`
	PlacedDateRange        *TimeRange         `json:"placedDateRange,omitempty"`
	Sort                   *MarketSort        `json:"sort,omitempty"`
	MaxResults             int                `json:"maxResults,omitempty"`
	IncludeSettledBets     bool               `json:"includeSettledBets,omitempty"`
	IncludeBspBets         bool               `json:"includeBspBets,omitempty"`
	NetOfCommission        bool               `json:"netOfCommission,omitempty"`
	Granularity            *TimeGranularity   `json:"granularity,omitempty"`
	BetIds                 []BetId            `json:"betIds,omitempty"`
	DateRange              *TimeRange         `json:"dateRange,omitempty"`
	OrderBy                *OrderBy           `json:"orderBy,omitempty"`
	SortDir                *SortDir           `json:"sortDir,omitempty"`
	FromRecord             int                `json:"fromRecord,omitempty"`
	RecordCount            int                `json:"recordCount,omitempty"`
	BetStatus              *BetStatus         `json:"betStatus,omitempty"`
	EventTypeIds           []EventTypeId      `json:"eventTypeIds,omitempty"`
	EventIds               []EventId          `json:"eventIds,omitempty"`
	RunnerIds              []RunnerId         `json:"runnerIds,omitempty"`
	Side                   *Side              `json:"side,omitempty"`
	SettledDateRange       *TimeRange         `json:"settledDateRange,omitempty"`
	GroupBy                *GroupBy           `json:"groupBy,omitempty"`
	IncludeItemDescription bool               `json:"includeItemDescription,omitempty"`
	MarketId               string             `json:"marketId,omitempty"`
	/*is instructions in API*/ PlaceInstructions []PlaceInstruction `json:"instructions,omitempty"`
	/*is instructions in API*/ CancelInstructions []CancelInstruction `json:"instructions,omitempty"`
	/*is instructions in API*/ UpdateInstructions []UpdateInstruction `json:"instructions,omitempty"`
	/*is instructions in API*/ ReplaceInstructions []ReplaceInstruction `json:"instructions,omitempty"`
	CustomerRef                                    string               `json:"customerRef,omitempty"`
}

// Below are Betfair API structs

type MarketFilter struct {
	TextQuery          string              `json:"textQuery,omitempty"`
	ExchangeIds        []string            `json:"exchangeIds,omitempty"`
	EventTypeIds       []string            `json:"eventTypeIds,omitempty"`
	EventIds           []string            `json:"eventIds,omitempty"`
	CompetitionIds     []string            `json:"competitionIds,omitempty"`
	MarketIds          []string            `json:"marketIds,omitempty"`
	Venues             []string            `json:"venues,omitempty"`
	BspOnly            bool                `json:"bspOnly,omitempty"`
	TurnInPlayEnabled  bool                `json:"turnInPlayEnabled,omitempty"`
	InPlayOnly         bool                `json:"inPlayOnly,omitempty"`
	MarketBettingTypes []MarketBettingType `json:"marketBettingTypes,omitempty"`
	MarketCountries    []string            `json:"marketCountries,omitempty"`
	MarketTypeCodes    []string            `json:"marketTypeCodes,omitempty"`
	MarketStartTime    *TimeRange          `json:"marketStartTime,omitempty"`
	WithOrders         []OrderStatus       `json:"withOrders,omitempty"`
}

type MarketCatalogue struct {
	MarketId        string            `json:"marketId,omitempty"`
	MarketName      string            `json:"marketName,omitempty"`
	MarketStartTime time.Time         `json:"marketStartTime,omitempty"`
	Description     MarketDescription `json:"description,omitempty"`
	TotalMatched    float64           `json:"totalMatched,omitempty"`
	Runners         []RunnerCatalog   `json:"runners,omitempty"`
	EventType       EventType         `json:"eventType,omitempty"`
	Competition     Competition       `json:"competition,omitempty"`
	Event           Event             `json:"event,omitempty"`
}

type MarketBook struct {
	MarketId              string       `json:"marketId,omitempty"`
	IsMarketDataDelayed   bool         `json:"isMarketDataDelayed,omitempty"`
	Status                MarketStatus `json:"status,omitempty"`
	BetDelay              int          `json:"betDelay,omitempty"`
	BspReconciled         bool         `json:"bspReconciled,omitempty"`
	Complete              bool         `json:"complete,omitempty"`
	Inplay                bool         `json:"inplay,omitempty"`
	NumberOfWinners       int          `json:"numberOfWinners,omitempty"`
	NumberOfRunners       int          `json:"numberOfRunners,omitempty"`
	NumberOfActiveRunners int          `json:"numberOfActiveRunners,omitempty"`
	LastMatchTime         time.Time    `json:"lastMatchTime,omitempty"`
	TotalMatched          float64      `json:"totalMatched,omitempty"`
	TotalAvailable        float64      `json:"totalAvailable,omitempty"`
	CrossMatching         bool         `json:"crossMatching,omitempty"`
	RunnersVoidable       bool         `json:"runnersVoidable,omitempty"`
	Version               int          `json:"version,omitempty"`
	Runners               []Runner     `json:"runners,omitempty"`
}

type RunnerCatalog struct {
	SelectionId  int               `json:"selectionId,omitempty"`
	RunnerName   string            `json:"runnerName,omitempty"`
	Handicap     float64           `json:"handicap,omitempty"`
	SortPriority int               `json:"sortPriority,omitempty"`
	Metadata     map[string]string `json:"metadata,omitempty"`
}

type Runner struct {
	SelectionId      int            `json:"selectionId,omitempty"`
	Handicap         float64        `json:"handicap,omitempty"`
	Status           RunnerStatus   `json:"status,omitempty"`
	AdjustmentFactor float64        `json:"adjustmentFactor,omitempty"`
	LastPriceTraded  float64        `json:"lastPriceTraded,omitempty"`
	TotalMatched     float64        `json:"totalMatched,omitempty"`
	RemovalDate      time.Time      `json:"removalDate,omitempty"`
	Sp               StartingPrices `json:"sp,omitempty"`
	Ex               ExchangePrices `json:"ex,omitempty"`
	Orders           []Order        `json:"orders,omitempty"`
	Matches          []Match        `json:"matches,omitempty"`
}

type StartingPrices struct {
	NearPrice         float64     `json:"nearPrice,omitempty"`
	FarPrice          float64     `json:"farPrice,omitempty"`
	BackStakeTaken    []PriceSize `json:"backStakeTaken,omitempty"`
	LayLiabilityTaken []PriceSize `json:"layLiabilityTaken,omitempty"`
	ActualSP          float64     `json:"actualSP,omitempty"`
}

type ExchangePrices struct {
	AvailableToBack []PriceSize `json:"availableToBack,omitempty"`
	AvailableToLay  []PriceSize `json:"availableToLay,omitempty"`
	TradedVolume    []PriceSize `json:"tradedVolume,omitempty"`
}

type Event struct {
	Id          string    `json:"id,omitempty"`
	Name        string    `json:"name,omitempty"`
	CountryCode string    `json:"countryCode,omitempty"`
	Timezone    string    `json:"timezone,omitempty"`
	Venue       string    `json:"venue,omitempty"`
	OpenDate    time.Time `json:"openDate,omitempty"` //Europe/London GMT by default
}

type EventResult struct {
	Event       Event `json:"event,omitempty"`
	MarketCount int   `json:"marketCount,omitempty"`
}

type Competition struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type CompetitionResult struct {
	Competition       Competition `json:"competition,omitempty"`
	MarketCount       int         `json:"marketCount,omitempty"`
	CompetitionRegion string      `json:"competitionRegion,omitempty"`
}

type EventType struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type EventTypeResult struct {
	EventType   EventType `json:"eventType,omitempty"`
	MarketCount int       `json:"marketCount,omitempty"`
}

type MarketTypeResult struct {
	MarketType  string `json:"marketType,omitempty"`
	MarketCount int    `json:"marketCount,omitempty"`
}

type CountryCodeResult struct {
	CountryCode string `json:"countryCode,omitempty"`
	MarketCount int    `json:"marketCount,omitempty"`
}

type VenueResult struct {
	Venue       string `json:"venue,omitempty"`
	MarketCount int    `json:"marketCount,omitempty"`
}

type TimeRange struct {
	From time.Time `json:"from,omitempty"`
	To   time.Time `json:"to,omitempty"`
}

type TimeRangeResult struct {
	TimeRange   TimeRange `json:"timeRange,omitempty"`
	MarketCount int       `json:"marketCount,omitempty"`
}

type Order struct {
	BetId           string          `json:"betId,omitempty"`
	OrderType       OrderType       `json:"orderType,omitempty"`
	Status          OrderStatus     `json:"status,omitempty"`
	PersistenceType PersistenceType `json:"persistenceType,omitempty"`
	Side            Side            `json:"side,omitempty"`
	Price           float64         `json:"price,omitempty"`
	Size            float64         `json:"size,omitempty"`
	BspLiability    float64         `json:"bspLiability,omitempty"`
	PlacedDate      time.Time       `json:"placedDate,omitempty"`
	avgPriceMatched float64         `json:"avgPriceMatched,omitempty"`
	SizeMatched     float64         `json:"sizeMatched,omitempty"`
	SizeRemaining   float64         `json:"sizeRemaining,omitempty"`
	SizeLapsed      float64         `json:"sizeLapsed,omitempty"`
	SizeCancelled   float64         `json:"sizeCancelled,omitempty"`
	SizeVoided      float64         `json:"sizeVoided,omitempty"`
}

type Match struct {
	BetId     string    `json:"betId,omitempty"`
	MatchId   string    `json:"matchId,omitempty"`
	Side      Side      `json:"side,omitempty"`
	Price     float64   `json:"price,omitempty"`
	Size      float64   `json:"size,omitempty"`
	MatchDate time.Time `json:"matchDate,omitempty"`
}

type MarketDescription struct {
	PersistenceEnabled bool              `json:"persistenceEnabled,omitempty"`
	BspMarket          bool              `json:"bspMarket,omitempty"`
	MarketTime         time.Time         `json:"marketTime,omitempty"`
	SuspendTime        time.Time         `json:"suspendTime,omitempty"`
	SettleTime         time.Time         `json:"settleTime,omitempty"`
	BettingType        MarketBettingType `json:"bettingType,omitempty"`
	TurnInPlayEnabled  bool              `json:"turnInPlayEnabled,omitempty"`
	MarketType         string            `json:"marketType,omitempty"`
	Regulator          string            `json:"regulator,omitempty"`
	MarketBaseRate     float64           `json:"marketBaseRate,omitempty"`
	DiscountAllowed    bool              `json:"discountAllowed,omitempty"`
	Wallet             string            `json:"wallet,omitempty"`
	Rules              string            `json:"rules,omitempty"`
	RulesHasDate       bool              `json:"rulesHasDate,omitempty"`
	EachWayDivisor     float64           `json:"eachWayDivisor,omitempty"`
	Clarifications     string            `json:"clarifications,omitempty"`
}

type MarketRates struct {
	MarketBaseRate  float64 `json:"marketBaseRate,omitempty"`
	DiscountAllowed bool    `json:"discountAllowed,omitempty"`
}

type MarketLicence struct {
	Wallet         string `json:"wallet,omitempty"`
	Rules          string `json:"rules,omitempty"`
	RulesHasDate   bool   `json:"rulesHasDate,omitempty"`
	Clarifications string `json:"clarifications,omitempty"`
}

type MarketLineRangeInfo struct {
	MaxUnitValue float64 `json:"maxUnitValue,omitempty"`
	MinUnitValue float64 `json:"minUnitValue,omitempty"`
	Interval     float64 `json:"interval,omitempty"`
	MarketUnit   string  `json:"marketUnit,omitempty"`
}

type PriceSize struct {
	Price float64 `json:"price,omitempty"`
	Size  float64 `json:"size,omitempty"`
}

type ClearedOrderSummary struct {
	EventTypeId     EventTypeId     `json:"eventTypeId,omitempty"`
	EventId         EventId         `json:"eventId,omitempty"`
	MarketId        MarketId        `json:"marketId,omitempty"`
	SelectionId     SelectionId     `json:"selectionId,omitempty"`
	Handicap        Handicap        `json:"handicap,omitempty"`
	BetId           BetId           `json:"betId,omitempty"`
	PlacedDate      time.Time       `json:"placedDate,omitempty"`
	PersistenceType PersistenceType `json:"persistenceType,omitempty"`
	OrderType       OrderType       `json:"orderType,omitempty"`
	Side            Side            `json:"side,omitempty"`
	ItemDescription ItemDescription `json:"itemDescription,omitempty"`
	PriceRequested  Price           `json:"priceRequested,omitempty"`
	SettledDate     time.Time       `json:"settledDate,omitempty"`
	BetCount        int             `json:"betCount,omitempty"`
	Commission      Size            `json:"commission,omitempty"`
	PriceMatched    Price           `json:"priceMatched,omitempty"`
	PriceReduced    bool            `json:"priceReduced,omitempty"`
	SizeSettled     Size            `json:"sizeSettled,omitempty"`
	Profit          Size            `json:"profit,omitempty"`
	SizeCancelled   Size            `json:"sizeCancelled,omitempty"`
}

type ClearedOrderSummaryReport struct {
	ClearedOrders []ClearedOrderSummary `json:"clearedOrders,omitempty"`
	MoreAvailable bool                  `json:"moreAvailable,omitempty"`
}

type ItemDescription struct {
	EventTypeDesc   string    `json:"eventTypeDesc,omitempty"`
	EventDesc       string    `json:"eventDesc,omitempty"`
	MarketDesc      string    `json:"marketDesc,omitempty"`
	MarketStartTime time.Time `json:"marketStartTime,omitempty"`
	RunnerDesc      string    `json:"runnerDesc,omitempty"`
	NumberOfWinners int       `json:"numberOfWinners,omitempty"`
}

type RunnerId struct {
	MarketId    MarketId    `json:"marketId,omitempty"`
	SelectionId SelectionId `json:"selectionId,omitempty"`
	Handicap    Handicap    `json:"handicap,omitempty"`
}

type CurrentOrderSummaryReport struct {
	CurrentOrders []CurrentOrderSummary `json:"currentOrders,omitempty"`
	MoreAvailable bool                  `json:"moreAvailable,omitempty"`
}

type CurrentOrderSummary struct {
	BetId               string          `json:"betId,omitempty"`
	MarketId            string          `json:"marketId,omitempty"`
	SelectionId         int64           `json:"selectionId,omitempty"`
	Handicap            float64         `json:"handicap,omitempty"`
	PriceSize           PriceSize       `json:"priceSize,omitempty"`
	BspLiability        float64         `json:"bspLiability,omitempty"`
	Side                Side            `json:"side,omitempty"`
	Status              OrderStatus     `json:"status,omitempty"`
	PersistenceType     PersistenceType `json:"persistenceType,omitempty"`
	OrderType           OrderType       `json:"orderType,omitempty"`
	PlacedDate          time.Time       `json:"placedDate,omitempty"`
	MatchedDate         time.Time       `json:"matchedDate,omitempty"`
	AveragePriceMatched float64         `json:"averagePriceMatched,omitempty"`
	SizeMatched         float64         `json:"sizeMatched,omitempty"`
	SizeRemaining       float64         `json:"sizeRemaining,omitempty"`
	SizeLapsed          float64         `json:"sizeLapsed,omitempty"`
	SizeCancelled       float64         `json:"sizeCancelled,omitempty"`
	SizeVoided          float64         `json:"sizeVoided,omitempty"`
	RegulatorAuthCode   string          `json:"regulatorAuthCode,omitempty"`
	RegulatorCode       string          `json:"regulatorCode,omitempty"`
}

type PlaceInstruction struct {
	OrderType          OrderType          `json:"orderType,omitempty"`
	SelectionId        int64              `json:"selectionId,omitempty"`
	Handicap           float64            `json:"handicap,omitempty"`
	Side               Side               `json:"side,omitempty"`
	LimitOrder         LimitOrder         `json:"limitOrder,omitempty"`
	LimitOnCloseOrder  LimitOnCloseOrder  `json:"limitOnCloseOrder,omitempty"`
	MarketOnCloseOrder MarketOnCloseOrder `json:"marketOnCloseOrder,omitempty"`
}

type PlaceExecutionReport struct {
	CustomerRef        string                   `json:"customerRef,omitempty"`
	Status             ExecutionReportStatus    `json:"status,omitempty"`
	ErrorCode          ExecutionReportErrorCode `json:"errorCode,omitempty"`
	MarketId           string                   `json:"marketId,omitempty"`
	InstructionReports []PlaceInstructionReport `json:"instructionReports,omitempty"`
}

type LimitOrder struct {
	Size            float64         `json:"size,omitempty"`
	Price           float64         `json:"price,omitempty"`
	PersistenceType PersistenceType `json:"persistenceType,omitempty"`
}

type LimitOnCloseOrder struct {
	Liability float64 `json:"liability,omitempty"`
	Price     float64 `json:"price,omitempty"`
}

type MarketOnCloseOrder struct {
	Liability float64 `json:"liability,omitempty"`
}

type PlaceInstructionReport struct {
	Status              InstructionReportStatus    `json:"status,omitempty"`
	ErrorCode           InstructionReportErrorCode `json:"errorCode,omitempty"`
	Instruction         PlaceInstruction           `json:"instruction,omitempty"`
	BetId               string                     `json:"betId,omitempty"`
	PlacedDate          time.Time                  `json:"placedDate,omitempty"`
	AveragePriceMatched float64                    `json:"averagePriceMatched,omitempty"`
	SizeMatched         float64                    `json:"sizeMatched,omitempty"`
}

type CancelInstruction struct {
	BetId         string  `json:"betId,omitempty"`
	SizeReduction float64 `json:"sizeReduction,omitempty"`
}

type CancelExecutionReport struct {
	CustomerRef        string                    `json:"customerRef,omitempty"`
	Status             ExecutionReportStatus     `json:"status,omitempty"`
	ErrorCode          ExecutionReportErrorCode  `json:"errorCode,omitempty"`
	MarketId           string                    `json:"marketId,omitempty"`
	InstructionReports []CancelInstructionReport `json:"instructionReports,omitempty"`
}

type ReplaceInstruction struct {
	BetId    string  `json:"betId,omitempty"`
	NewPrice float64 `json:"newPrice,omitempty"`
}

type ReplaceExecutionReport struct {
	CustomerRef        string                     `json:"customerRef,omitempty"`
	Status             ExecutionReportStatus      `json:"status,omitempty"`
	ErrorCode          ExecutionReportErrorCode   `json:"errorCode,omitempty"`
	MarketId           string                     `json:"marketId,omitempty"`
	InstructionReports []ReplaceInstructionReport `json:"instructionReports,omitempty"`
}

type ReplaceInstructionReport struct {
	Status                  InstructionReportStatus    `json:"status,omitempty"`
	ErrorCode               InstructionReportErrorCode `json:"errorCode,omitempty"`
	CancelInstructionReport CancelInstructionReport    `json:"cancelInstructionReport,omitempty"`
	PlaceInstructionReport  PlaceInstructionReport     `json:"placeInstructionReport,omitempty"`
}

type CancelInstructionReport struct {
	Status        InstructionReportStatus    `json:"status,omitempty"`
	ErrorCode     InstructionReportErrorCode `json:"errorCode,omitempty"`
	Instruction   CancelInstruction          `json:"instruction,omitempty"`
	SizeCancelled float64                    `json:"sizeCancelled,omitempty"`
	CancelledDate time.Time                  `json:"cancelledDate,omitempty"`
}

type UpdateInstruction struct {
	BetId              string          `json:"betId,omitempty"`
	NewPersistenceType PersistenceType `json:"newPersistenceType,omitempty"`
}

type UpdateExecutionReport struct {
	CustomerRef        string                    `json:"customerRef,omitempty"`
	Status             ExecutionReportStatus     `json:"status,omitempty"`
	ErrorCode          ExecutionReportErrorCode  `json:"errorCode,omitempty"`
	MarketId           string                    `json:"marketId,omitempty"`
	InstructionReports []UpdateInstructionReport `json:"instructionReports,omitempty"`
}

type UpdateInstructionReport struct {
	Status      InstructionReportStatus    `json:"status,omitempty"`
	ErrorCode   InstructionReportErrorCode `json:"errorCode,omitempty"`
	Instruction UpdateInstruction          `json:"instruction,omitempty"`
}

type PriceProjection struct {
	PriceData             []PriceData           `json:"priceData,omitempty"`
	ExBestOffersOverrides ExBestOffersOverrides `json:"exBestOffersOverrides,omitempty"`
	Virtualise            bool                  `json:"virtualise,omitempty"`
	RolloverStakes        bool                  `json:"rolloverStakes,omitempty"`
}

type ExBestOffersOverrides struct {
	BestPricesDepth          int         `json:"bestPricesDepth,omitempty"`
	RollupModel              RollupModel `json:"rollupModel,omitempty"`
	RollupLimit              int         `json:"rollupLimit,omitempty"`
	RollupLiabilityThreshold float64     `json:"rollupLiabilityThreshold,omitempty"`
	RollupLiabilityFactor    int         `json:"rollupLiabilityFactor,omitempty"`
}

type MarketProfitAndLoss struct {
	MarketId          string                `json:"marketId,omitempty"`
	CommissionApplied float64               `json:"commissionApplied,omitempty"`
	ProfitAndLosses   []RunnerProfitAndLoss `json:"profitAndLosses,omitempty"`
}

type RunnerProfitAndLoss struct {
	SelectionId SelectionId `json:"selectionId,omitempty"`
	IfWin       float64     `json:"ifWin,omitempty"`
	IfLose      float64     `json:"ifLose,omitempty"`
	IfPlace     float64     `json:"ifPlace,omitempty"`
}

type MarketType string
type Venue string
type MarketId string
type SelectionId int64
type Handicap float64
type EventId string
type EventTypeId string
type CountryCode string
type ExchangeId string
type CompetitionId string
type Price float64
type Size float64
type BetId string
type MatchId string
