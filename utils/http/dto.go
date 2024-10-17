package httputils

import (
	"encoding/json"
	"errors"
	"github.com/nahid/gohttp"
)

type Postman struct {
	Url          string
	Queries      map[string]string
	Headers      map[string]string
	Body         map[string]interface{}
	BodyByte     []byte
	StatusCode   int
	ErrorMessage error
	Response     map[string]interface{}
}

func (postman *Postman) validateAllRequest() error {
	if postman.Url == "" {
		return errors.New("empty url")
	}
	return nil
}

func (postman *Postman) validateGetRequest() error {
	if err := postman.validateAllRequest(); err != nil {
		return err
	}
	return nil
}

func (postman *Postman) validatePostRequest() error {
	if err := postman.validateAllRequest(); err != nil {
		return err
	}

	if len(postman.Body) == 0 && len(postman.BodyByte) == 0 {
		return errors.New("missing body")
	}
	if len(postman.Body) != 0 && len(postman.BodyByte) != 0 {
		return errors.New("provide any of body_map or body_byte")
	}

	return nil
}

func (postman *Postman) setDefaultHeaders() {
	if postman.Headers == nil {
		postman.Headers = make(map[string]string, 2)
	}

	// Initialize default headers
	defaultHeaders := map[string]string{
		"content-type": "application/json",
	}

	for k, v := range defaultHeaders {
		postman.Headers[k] = v
	}
}

func (postman *Postman) setApiReturnData(res *gohttp.Response, err error) {
	body, _ := res.GetBodyAsString()

	postman.Response = nil
	postman.ErrorMessage = nil
	postman.StatusCode = 200

	if res != nil {
		postman.StatusCode = res.GetStatusCode()
	}
	if err != nil {
		postman.ErrorMessage = errors.New(err.Error())
	}

	var responseJson map[string]interface{}
	json.Unmarshal([]byte(body), &responseJson)
	postman.Response = responseJson
}
