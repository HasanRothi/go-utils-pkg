package httputils

import (
	"github.com/nahid/gohttp"
)

func (postman *Postman) Post() {
	if err := postman.validatePostRequest(); err != nil {
		postman.ErrorMessage = err
		postman.StatusCode = 400
		return
	}

	postman.setDefaultHeaders()

	//Api initializer
	var req *gohttp.Request

	//Either json or body byte
	if postman.Body != nil {
		req = gohttp.NewRequest().JSON(postman.Body)
	} else if postman.BodyByte != nil {
		req = gohttp.NewRequest().Body(postman.BodyByte)
	}

	//Final Hit
	apiRes, err := req.Headers(postman.Headers).Query(postman.Queries).Post(postman.Url)

	postman.setApiReturnData(apiRes, err)

}
