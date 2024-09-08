package httputils

import (
	"encoding/json"
	"errors"
	"github.com/nahid/gohttp"
	"strconv"
)

func CommonExternalPostCall(url string, data map[string]interface{}) (int, error, map[string]interface{}) {
	req := gohttp.NewRequest()
	var response map[string]interface{}
	apiRes, err := req.JSON(data).Headers(map[string]string{
		"content-type": "application/json",
	}).Post(url)
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
