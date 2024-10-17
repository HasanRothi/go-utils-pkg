package httputils

import (
	"github.com/nahid/gohttp"
)

func (postman *Postman) Get() {
	if err := postman.validateGetRequest(); err != nil {
		postman.ErrorMessage = err
		postman.StatusCode = 400
	}

	postman.setDefaultHeaders()

	req := gohttp.NewRequest()

	apiRes, err := req.Headers(postman.Headers).Query(postman.Queries).Get(postman.Url)

	postman.setApiReturnData(apiRes, err)
}
