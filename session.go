package gobet

import (
	"errors"
	"fmt"
	"time"
)

const loginUrl string = "https://identitysso.betfair.com/api/login"
const logoutUrl string = "https://identitysso.betfair.com/api/logout"
const keepAliveUrl string = "https://identitysso.betfair.com/api/keepAlive"

type SessionResponse struct {
	Token   string `json:"token"`
	Product string `json:"product"`
	Status  string `json:"status"`
	Error   string `json:"error"`
}

type Session struct {
	username, password, appKey, Token string
}

func (s Session) String() string {
	if s.Token != "" {
		return fmt.Sprintf("User %s IS logged in! Token: %s\n", s.username, s.Token)
	} else {
		return fmt.Sprintf("User %s IS NOT logged in!", s.username)
	}
}

func (s *Session) Login() error {
	var jsonStr = fmt.Sprintf("username=%s&password=%s", s.username, s.password)
	var headers = map[string]string{"Accept": "application/json", "X-Application": s.appKey, "Content-Type": "application/x-www-form-urlencoded"}
	var sessionResp SessionResponse
	err := DoRequest("POST", loginUrl, jsonStr, headers, &sessionResp)
	s.Token = sessionResp.Token
	fmt.Println(sessionResp)
	return err
}

func (s *Session) Logout() error {
	var headers = map[string]string{"Accept": "application/json", "X-Application": s.appKey, "X-Authentication": s.Token}
	var sessionResp SessionResponse
	err := DoRequest("GET", logoutUrl, "", headers, &sessionResp)
	fmt.Println(sessionResp)
	return err
}

func (s *Session) KeepAlive() error {
	var headers = map[string]string{"Accept": "application/json", "X-Application": s.appKey, "X-Authentication": s.Token}
	var sessionResp SessionResponse
	err := DoRequest("GET", keepAliveUrl, "", headers, &sessionResp)
	if err == nil && sessionResp.Status == "FAIL" {
		err = errors.New("KeepAlive Failed")
	}
	fmt.Println(sessionResp)
	return err
}

func (s *Session) DoKeepAliveEvery(interval time.Duration) error {
	tickerChannel := time.Tick(interval)
	go func() {
		for _ = range tickerChannel {
			for retries := 0; retries < 5; retries++ {
				err := s.KeepAlive()
				if err != nil {
					fmt.Println("Keep Alive failed. Retry ", retries+1, "in ", 3*(retries+1), "secs")
					// retry for 5 times with linear back-off
					time.Sleep(time.Second * time.Duration(3*(retries+1)))
				} else {
					break
				}

			}
		}
	}()
	return nil
}

func NewSession(username, password, appKey string) *Session {
	return &Session{username: username, password: password, appKey: appKey}
}
