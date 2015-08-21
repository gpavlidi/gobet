package gobet

import (
	"bytes" //needed for debug
	"encoding/json"
	"io/ioutil" //needed for debug
	"log"
	"net/http"
	"strings"
)

func DoRequest(verb, url, payload string, headers map[string]string, outputStructPointer interface{}, logger *log.Logger) error {
	if logger == nil {
		logger = log.New(ioutil.Discard, "", 0)
	}

	req, err := http.NewRequest(verb, url, strings.NewReader(payload))
	for headerKey, headerVal := range headers {
		req.Header.Set(headerKey, headerVal)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logger.Println(err)
		return err
	}
	defer resp.Body.Close()

	// debug request and response
	//logger.Println(req, payload)
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	//logger.Println("\n\n", string(bodyBytes), "\n\n")
	resp.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	err = json.NewDecoder(resp.Body).Decode(outputStructPointer)
	if err != nil {
		logger.Println(err, "Got this instead: ", "Request", req, payload, "Response", string(bodyBytes))
		return err
	}
	return nil
}
