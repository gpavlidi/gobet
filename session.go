package gobet

import (
	"errors"
	"fmt"
	"time"
)

const sessionUrlPrefix string = "https://identitysso.betfair.com/api/"

type SessionResponse struct {
	Token   string `json:"token"`
	Product string `json:"product"`
	Status  string `json:"status"`
	Error   string `json:"error"`
}

// below expects an appKey even though we havent logged in yet!!!
// Passing it a random 16 char appkey
func (s *Session) Login() error {
	payload := fmt.Sprintf("username=%s&password=%s", s.configuration.Username, s.configuration.Password)
	headers := map[string]string{"Accept": "application/json", "X-Application": "XXXXXXXXXXXXXXXX", "Content-Type": "application/x-www-form-urlencoded"}
	url := fmt.Sprintf("%s%s", sessionUrlPrefix, "login")
	response := SessionResponse{}
	err := DoRequest("POST", url, payload, headers, &response, s.logger)
	if err != nil {
		return err
	} else if response.Status == "FAIL" {
		return errors.New("Login Failed")
	}

	s.token = response.Token
	return nil
}

func (s *Session) Logout() error {
	headers := map[string]string{"Accept": "application/json", "X-Application": s.configuration.AppKey, "X-Authentication": s.token}
	url := fmt.Sprintf("%s%s", sessionUrlPrefix, "logout")
	response := SessionResponse{}
	err := DoRequest("GET", url, "", headers, &response, s.logger)
	if err != nil {
		return err
	} else if response.Status == "FAIL" {
		return errors.New("Logout Failed")
	}

	return nil
}

func (s *Session) KeepAlive() error {
	headers := map[string]string{"Accept": "application/json", "X-Application": s.configuration.AppKey, "X-Authentication": s.token}
	url := fmt.Sprintf("%s%s", sessionUrlPrefix, "keepAlive")
	response := SessionResponse{}
	err := DoRequest("GET", url, "", headers, &response, s.logger)
	if err != nil {
		return err
	} else if response.Status == "FAIL" {
		return errors.New("KeepAlive Failed")
	}

	return nil
}

func (s *Session) DoKeepAliveEvery(interval time.Duration) error {
	tickerChannel := time.Tick(interval)
	go func() {
		for _ = range tickerChannel {
			for retries := 0; retries < 5; retries++ {
				err := s.KeepAlive()
				if err != nil {
					s.logger.Println("Keep Alive failed. Retry ", retries+1, "in ", 3*(retries+1), "secs")
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
