package gobet

import (
	"encoding/json"
	"fmt"
)

const accountsUrlPrefix string = "https://api.betfair.com/exchange/account/rest/v1.0/"

func (s *Session) CreateDeveloperAppKeys(inputParams AccountsInputParams) (DeveloperApp, error) {
	response := DeveloperApp{}
	err := s.doAccountsRequest("createDeveloperAppKeys", inputParams, &response)
	return response, err
}

func (s *Session) GetDeveloperAppKeys(inputParams AccountsInputParams) ([]DeveloperApp, error) {
	response := []DeveloperApp{}
	err := s.doAccountsRequest("getDeveloperAppKeys", inputParams, &response)
	return response, err
}

func (s *Session) GetAccountFunds(inputParams AccountsInputParams) (AccountFundsResponse, error) {
	response := AccountFundsResponse{}
	err := s.doAccountsRequest("getAccountFunds", inputParams, &response)
	return response, err
}

func (s *Session) TransferFunds(inputParams AccountsInputParams) (TransferResponse, error) {
	response := TransferResponse{}
	err := s.doAccountsRequest("transferFunds", inputParams, &response)
	return response, err
}

func (s *Session) GetAccountDetails(inputParams AccountsInputParams) (AccountDetailsResponse, error) {
	response := AccountDetailsResponse{}
	err := s.doAccountsRequest("getAccountDetails", inputParams, &response)
	return response, err
}

func (s *Session) GetVendorClientId(inputParams AccountsInputParams) (string, error) {
	response := ""
	err := s.doAccountsRequest("getVendorClientId", inputParams, &response)
	return response, err
}

func (s *Session) GetApplicationSubscriptionToken(inputParams AccountsInputParams) (string, error) {
	response := ""
	err := s.doAccountsRequest("getApplicationSubscriptionToken", inputParams, &response)
	return response, err
}

func (s *Session) ActivateApplicationSubscription(inputParams AccountsInputParams) (Status, error) {
	var response Status
	err := s.doAccountsRequest("activateApplicationSubscription", inputParams, &response)
	return response, err
}

func (s *Session) CancelApplicationSubscription(inputParams AccountsInputParams) (Status, error) {
	var response Status
	err := s.doAccountsRequest("cancelApplicationSubscription", inputParams, &response)
	return response, err
}

func (s *Session) UpdateApplicationSubscription(inputParams AccountsInputParams) (string, error) {
	response := ""
	err := s.doAccountsRequest("updateApplicationSubscription", inputParams, &response)
	return response, err
}

func (s *Session) ListApplicationSubscriptionTokens(inputParams AccountsInputParams) ([]ApplicationSubscription, error) {
	response := []ApplicationSubscription{}
	err := s.doAccountsRequest("listApplicationSubscriptionTokens", inputParams, &response)
	return response, err
}

func (s *Session) ListAccountSubscriptionTokens(inputParams AccountsInputParams) ([]AccountSubscription, error) {
	response := []AccountSubscription{}
	err := s.doAccountsRequest("listAccountSubscriptionTokens", inputParams, &response)
	return response, err
}

func (s *Session) GetApplicationSubscriptionHistory(inputParams AccountsInputParams) ([]SubscriptionHistory, error) {
	response := []SubscriptionHistory{}
	err := s.doAccountsRequest("getApplicationSubscriptionHistory", inputParams, &response)
	return response, err
}

func (s *Session) GetAccountStatement(inputParams AccountsInputParams) (AccountStatementReport, error) {
	response := AccountStatementReport{}
	err := s.doAccountsRequest("getAccountStatement", inputParams, &response)
	return response, err
}

func (s *Session) ListCurrencyRates(inputParams AccountsInputParams) ([]CurrencyRate, error) {
	response := []CurrencyRate{}
	err := s.doAccountsRequest("listCurrencyRates", inputParams, &response)
	return response, err
}

func (s *Session) doAccountsRequest(accountsOperation string, inputParams AccountsInputParams, outputStructPtr interface{}) error {
	url := fmt.Sprintf("%s%s/", accountsUrlPrefix, accountsOperation)
	headers := map[string]string{"Content-Type": "application/json", "Accept": "application/json", "X-Application": s.configuration.AppKey, "X-Authentication": s.token}

	payload, err := json.Marshal(inputParams)
	if err != nil {
		return err
	}

	err = DoRequest("POST", url, string(payload), headers, outputStructPtr, s.logger)

	return err
}
