package gobet

import (
	"encoding/json"
	"net/http"
	"strings"
)

func DoRequest(verb, url, payload string, headers map[string]string, responsePtr interface{}) error {

	req, err := http.NewRequest(verb, url, strings.NewReader(payload))
	for headerKey, headerVal := range headers {
		req.Header.Set(headerKey, headerVal)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
		//panic(err)
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(responsePtr)
	if err != nil {
		return err
	}
	return nil
}
