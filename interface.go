package http_client_go

import "context"

type Interface interface {
	// WithMethod specify method
	WithMethod(m Method) Interface
	// WithHost Specify the request host.
	WithHost(host string) Interface
	// WithURL specify the url which will be requested
	WithURL(url string) Interface
	// WithHttps use https
	WithHttps() Interface
	// AddHeader add header
	AddHeader(key, value string) Interface
	// WithContentTypeJson add specify content as application/json
	WithContentTypeJson() Interface
	// WithAuthorization add Authorization header to http client
	WithAuthorization(authInfo string) Interface
	// WithHeaders specify header which will replace the header your added before.
	WithHeaders(headers map[string]string) Interface
	// WithBody specify body which will be used in http request
	WithBody(body interface{}) Interface
	// WithUrlParams specify url params which will be put in url, and replace the url params you added before.
	WithUrlParams(params map[string]string) Interface
	// AddUrlParams add url param
	AddUrlParams(key, value string) Interface
	// Do execute http request.
	Do(ctx context.Context) Interface
	// RespBytes marshal response as bytes and return errors.
	RespBytes() ([]byte, error)
	// RespMap marshal response as map and return errors
	RespMap() (map[string]interface{}, error)
	// Error marshal response and return error that has caused.
	Error() error
}

type Option func(p Interface)

func NewDefaultClient(options ...Option) Interface {
	d := &defaultHttpClient{}
	for _, op := range options {
		op(d)
	}

	return d
}
