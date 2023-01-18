package util

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/wwkeyboard/sunsetwx/logs"
)

func PostAPI(param []byte, url string) ([]byte, error) {
	var responseBody []byte
	var err error
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(param))
	if err != nil {
		logs.Log.Error("QueryAPI POST error, error message: %s", err)
	}

	defer resp.Body.Close()

	responseBody, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.Log.Error("QueryAPI ReadAll error, error message: %s", err)
	}

	return responseBody, err
}

func GetAPI(param map[string]string, url string) ([]byte, error) {
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		logs.Log.Error("GetAPI error, error message: %s", err)
	}

	q := req.URL.Query()
	for key, item := range param {
		q.Add(key, item)
	}

	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		logs.Log.Error("GetAPI Do error, error message: %s", err)
		return nil, err
	}

	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.Log.Error("GetAPI ReadAll error, error message: %s", err)
	}

	return responseBody, err
}
