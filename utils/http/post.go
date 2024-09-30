package httputils

import (
	"encoding/json"
	"errors"
	"github.com/nahid/gohttp"
	"strconv"
)

func Post(url string, queryParams ...ApiHelpers) (int, error, map[string]interface{}) {
	if queryParams == nil || (queryParams[0].Body == nil && queryParams[0].BodyByte == nil) {
		return 400, errors.New("missing body"), nil
	}

	if queryParams[0].Body != nil && queryParams[0].BodyByte != nil {
		return 400, errors.New("provide any of body_map or body_byte"), nil
	}

	query := map[string]string{}
	headers := map[string]string{}
	var response map[string]interface{}
	var body map[string]interface{}
	var bodyByte []byte

	//Payload body Map
	if queryParams[0].Body != nil {
		body = queryParams[0].Body
		bodyByte = nil
	}

	//Payload body byte
	if queryParams[0].BodyByte != nil {
		bodyByte = queryParams[0].BodyByte
		body = nil
	}

	//Params
	if queryParams[0].Query != nil {
		query = queryParams[0].Query
	}

	//Headers
	if queryParams[0].Headers != nil {
		headers = queryParams[0].Headers
	}

	if headers == nil {
		headers = make(map[string]string, 2)
	}

	// Initialize default headers
	defaultHeaders := map[string]string{
		"content-type": "application/json",
	}

	for k, v := range defaultHeaders {
		headers[k] = v
	}

	//Api initializer
	var req *gohttp.Request

	//Either json or body byte
	if body != nil {
		req = gohttp.NewRequest().JSON(body)
	} else if bodyByte != nil {
		req = gohttp.NewRequest().Body(bodyByte)
	}

	//Final Hit
	apiRes, err := req.Headers(headers).Query(query).Post(url)

	//Response
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
