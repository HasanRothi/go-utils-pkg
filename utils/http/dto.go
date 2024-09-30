package httputils

type ApiHelpers struct {
	Query    map[string]string
	Headers  map[string]string
	Body     map[string]interface{}
	BodyByte []byte
}
