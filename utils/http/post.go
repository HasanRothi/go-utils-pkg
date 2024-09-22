package httputils

import (
	"encoding/json"
	"errors"
	"github.com/nahid/gohttp"
	"strconv"
)

func CommonExternalPostCall(url string, data map[string]interface{}, queryParams ...ApiHelpers) (int, error, map[string]interface{}) {
	req := gohttp.NewRequest()
	var response map[string]interface{}

	// Initialize default headers
	defaultHeaders := map[string]string{
		"content-type": "application/json",
	}
	query := map[string]string{}
	headers := map[string]string{}

	if queryParams != nil && len(queryParams) > 0 {
		query = queryParams[0].Query
		headers = queryParams[0].Headers
		if headers == nil {
			headers = make(map[string]string, 2)
		}
	}

	for k, v := range defaultHeaders {
		headers[k] = v
	}

	apiRes, err := req.Headers(headers).Query(query).Get(url)

	if err == nil {
		body, _ := apiRes.GetBodyAsString()
		if apiRes.GetStatusCode() >= 500 {
			return -1, errors.New(strconv.Itoa(apiRes.GetStatusCode()) + " | " + body), nil
		}
		if apiRes.GetStatusCode() >= 400 {
			return -1, errors.New(strconv.Itoa(apiRes.GetStatusCode()) + " | " + body), nil
		}
		json.Unmarshal([]byte(body), &response)
		return 201, nil, response
	}
	return 400, err, response
}
