package gobet

import (
	"fmt"
	"time"
)

const inplayUrl string = "http://www.betfair.com/inplayservice/v1/eventTimeline?alt=json&eventId=%s&locale=en_GB&productType=EXCHANGE&regionCode=UK&ts=%d"

type EventTimeline struct {
	EventId            int64         `json:"eventId"`
	EventTypeId        int           `json:"eventTypeId"`
	Score              EventScores   `json:"score"`
	TimeElapsed        int           `json:"timeElapsed"`
	ElapsedRegularTime int           `json:"elapsedRegularTime"`
	ElapsedAddedTime   int           `json:"elapsedAddedTime"`
	UpdateDetails      []EventUpdate `json:"updateDetails"`
	Status             string        `json:"status"`
	InPlayMatchStatus  string        `json:"inPlayMatchStatus"`
}

type EventScores struct {
	Home EventScore `json:"home"`
	Away EventScore `json:"away"`
}

type EventScore struct {
	Name              string `json:"name"`
	Score             string `json:"score"`
	HalfTimeScore     string `json:"halfTimeScore"`
	FullTimeScore     string `json:"fullTimeScore"`
	PenaltiesScore    string `json:"penaltiesScore"`
	PenaltiesSequence []bool `json:"penaltiesSequence"`
	Games             string `json:"games"`
	Sets              string `json:"sets"`
}

type EventUpdate struct {
	UpdateTime time.Time `json:"updateTime"`
	UpdateId   int       `json:"updateId"`
	Team       string    `json:"team"`
	TeamName   string    `json:"teamName"`
	MatchTime  int       `json:"matchTime"`
	Type       string    `json:"type"`
	UpdateType string    `json:"updateType"`
}

func (s *Session) GetEventTimeline(eventId string) (EventTimeline, error) {
	url := fmt.Sprintf(inplayUrl, eventId, time.Now().Unix())
	headers := map[string]string{"Content-Type": "application/json", "Accept": "application/json"}
	timeline := EventTimeline{}
	err := DoRequest("GET", url, "", headers, &timeline, s.logger)

	return timeline, err
}
