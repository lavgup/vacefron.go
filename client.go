package vacefron

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const baseURL = "https://vacefron.nl/api/"

// makeRequest makes a request to the API
func makeRequest(endpoint string, params url.Values) (interface{}, error) {
	endpointURL := baseURL + endpoint

	if len(params) > 0 {
		endpointURL = endpointURL + "?" + params.Encode()
	}

	res, err := http.Get(endpointURL)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("request failed, status code %s", strconv.Itoa(res.StatusCode))
	}

	defer func() {
		rerr := res.Body.Close()
		if rerr != nil {
			err = rerr
		}
	}()
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	content := http.DetectContentType(body)
	// JSON response
	if strings.HasPrefix(content, "text/plain") {
		var data map[string]interface{}

		if err := json.Unmarshal(body, &data); err != nil {
			return nil, err
		}

		return data, nil
	}

	return body, nil
}
