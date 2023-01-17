package util

import (
	"io/ioutil"
	"net/http"

	"github.com/wwkeyboard/sunsetwx/logs"
)

func QueryAPI(param map[string]string, method string, url string) ([]byte, error) {
	logs.Log.Error("QueryAPI error, error message:")
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		logs.Log.Error("QueryAPI error, error message: %s", err)
	}

	q := req.URL.Query()
	// appending to existing query args
	for key, item := range param {
		q.Add(key, item)
	}

	// assign encoded query string to http request
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		logs.Log.Error("QueryAPI Do error, error message: %s", err)
		return nil, err
	}

	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.Log.Error("QueryAPI ReadAll error, error message: %s", err)
	}

	return responseBody, err
}
