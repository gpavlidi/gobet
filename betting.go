package gobet

import (
	"encoding/json"
	"fmt"
)

const bettingUrlPrefix string = "https://api.betfair.com/exchange/betting/rest/v1.0/"

func (s *Session) ListEventTypes(inputParams BettingInputParams) ([]EventTypeResult, error) {
	response := []EventTypeResult{}
	err := s.doBettingRequest("listEventTypes", inputParams, &response)
	return response, err
}

func (s *Session) ListCompetitions(inputParams BettingInputParams) ([]CompetitionResult, error) {
	response := []CompetitionResult{}
	err := s.doBettingRequest("listCompetitions", inputParams, &response)
	return response, err
}

func (s *Session) ListTimeRanges(inputParams BettingInputParams) ([]TimeRangeResult, error) {
	response := []TimeRangeResult{}
	err := s.doBettingRequest("listTimeRanges", inputParams, &response)
	return response, err
}

func (s *Session) ListEvents(inputParams BettingInputParams) ([]EventResult, error) {
	response := []EventResult{}
	err := s.doBettingRequest("listEvents", inputParams, &response)
	return response, err
}

func (s *Session) ListMarketTypes(inputParams BettingInputParams) ([]MarketTypeResult, error) {
	response := []MarketTypeResult{}
	err := s.doBettingRequest("listMarketTypes", inputParams, &response)
	return response, err
}

func (s *Session) ListCountries(inputParams BettingInputParams) ([]CountryCodeResult, error) {
	response := []CountryCodeResult{}
	err := s.doBettingRequest("listCountries", inputParams, &response)
	return response, err
}

func (s *Session) ListVenues(inputParams BettingInputParams) ([]VenueResult, error) {
	response := []VenueResult{}
	err := s.doBettingRequest("listVenues", inputParams, &response)
	return response, err
}

func (s *Session) ListMarketCatalogue(inputParams BettingInputParams) ([]MarketCatalogue, error) {
	response := []MarketCatalogue{}
	err := s.doBettingRequest("listMarketCatalogue", inputParams, &response)
	return response, err
}

func (s *Session) ListMarketBook(inputParams BettingInputParams) ([]MarketBook, error) {
	response := []MarketBook{}
	err := s.doBettingRequest("listMarketBook", inputParams, &response)
	return response, err
}

func (s *Session) ListMarketProfitAndLoss(inputParams BettingInputParams) ([]MarketProfitAndLoss, error) {
	response := []MarketProfitAndLoss{}
	err := s.doBettingRequest("listMarketProfitAndLoss", inputParams, &response)
	return response, err
}

func (s *Session) ListCurrentOrders(inputParams BettingInputParams) (CurrentOrderSummaryReport, error) {
	response := CurrentOrderSummaryReport{}
	err := s.doBettingRequest("listCurrentOrders", inputParams, &response)
	return response, err
}

func (s *Session) ListClearedOrders(inputParams BettingInputParams) (ClearedOrderSummaryReport, error) {
	response := ClearedOrderSummaryReport{}
	err := s.doBettingRequest("listClearedOrders", inputParams, &response)
	return response, err
}

func (s *Session) PlaceOrders(inputParams BettingInputParams) (PlaceExecutionReport, error) {
	response := PlaceExecutionReport{}
	err := s.doBettingRequest("placeOrders", inputParams, &response)
	return response, err
}

func (s *Session) CancelOrders(inputParams BettingInputParams) (CancelExecutionReport, error) {
	response := CancelExecutionReport{}
	err := s.doBettingRequest("cancelOrders", inputParams, &response)
	return response, err
}

func (s *Session) ReplaceOrders(inputParams BettingInputParams) (ReplaceExecutionReport, error) {
	response := ReplaceExecutionReport{}
	err := s.doBettingRequest("replaceOrders", inputParams, &response)
	return response, err
}

func (s *Session) UpdateOrders(inputParams BettingInputParams) (UpdateExecutionReport, error) {
	response := UpdateExecutionReport{}
	err := s.doBettingRequest("updateOrders", inputParams, &response)
	return response, err
}

func (s *Session) doBettingRequest(bettingOperation string, inputParams BettingInputParams, outputStructPtr interface{}) error {
	url := fmt.Sprintf("%s%s/", bettingUrlPrefix, bettingOperation)
	headers := map[string]string{"Content-Type": "application/json", "Accept": "application/json", "X-Application": s.configuration.AppKey, "X-Authentication": s.token}

	payload, err := json.Marshal(inputParams)
	if err != nil {
		return err
	}

	err = DoRequest("POST", url, string(payload), headers, outputStructPtr, s.logger)

	return err
}
