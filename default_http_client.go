package http_client_go

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type defaultHttpClient struct {
	isHttps   bool
	method    Method
	host      string
	url       string
	header    map[string]string
	urlParams map[string]string
	body      interface{}
	resp      []byte
	err       error
}

func (d *defaultHttpClient) WithMethod(m Method) Interface {
	d.method = m
	return d
}

func (d *defaultHttpClient) WithHost(host string) Interface {
	d.host = host
	return d
}

func (d *defaultHttpClient) WithURL(url string) Interface {
	d.url = url
	return d
}

func (d *defaultHttpClient) AddHeader(key, value string) Interface {
	if d.header == nil {
		d.header = map[string]string{}
	}

	d.header[key] = value
	return d
}

func (d *defaultHttpClient) WithHeaders(headers map[string]string) Interface {
	if d.header == nil {
		d.header = map[string]string{}
	}

	d.header = headers
	return d
}

func (d *defaultHttpClient) WithBody(body interface{}) Interface {
	d.body = body
	return d
}

func (d *defaultHttpClient) WithUrlParams(params map[string]string) Interface {
	if d.urlParams == nil {
		d.urlParams = map[string]string{}
	}

	d.urlParams = params
	return d
}

func (d *defaultHttpClient) AddUrlParams(key, value string) Interface {
	if d.urlParams == nil {
		d.urlParams = map[string]string{}
	}

	d.urlParams[key] = value
	return d
}

func (d *defaultHttpClient) WithHttps() Interface {
	d.isHttps = true
	return d
}

func (d *defaultHttpClient) WithContentTypeJson() Interface {
	if d.header == nil {
		d.header = map[string]string{}
	}

	d.header["Content-Type"] = "application/json"
	return d
}

func (d *defaultHttpClient) WithAuthorization(authInfo string) Interface {
	if d.header == nil {
		d.header = map[string]string{}
	}

	d.header[HeaderAuthorization] = authInfo

	return d
}

func (d *defaultHttpClient) Do(ctx context.Context) Interface {
	d.err = nil

	err := d.validate()
	if err != nil {
		d.err = err
		return d
	}
	d.complete()

	var baseURL string
	var requestURL string
	if d.isHttps {
		baseURL = "https://" + d.host + d.url
	} else {
		baseURL = "http://" + d.host + d.url
	}

	if d.urlParams != nil {
		params := url.Values{}
		for k, v := range d.urlParams {
			params.Set(k, v)
		}
		requestURL = baseURL + "?" + params.Encode()
	} else {
		requestURL = baseURL
	}

	reqBody, err := json.Marshal(d.body)
	if err != nil {
		d.err = err
		return d
	}

	req, err := http.NewRequest(d.method.String(), requestURL, bytes.NewBuffer(reqBody))
	if err != nil {
		d.err = errors.Wrap(err, "Error create http request")
		return d
	}

	for k, v := range d.header {
		req.Header.Set(k, v)
	}

	// 发送请求并获取响应
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		d.err = errors.Wrap(err, "Error sending request")
		return d
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		d.err = errors.New(fmt.Sprintf("Response status code not 200, is %v", resp.StatusCode))
		return d
	}

	// 读取响应内容
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		d.err = errors.Wrap(err, "Error reading response")
		return d
	}

	d.resp = responseBody
	return d
}

func (d *defaultHttpClient) RespBytes() ([]byte, error) {
	if d.err != nil {
		return nil, d.err
	}

	return d.resp, nil
}

func (d *defaultHttpClient) RespMap() (map[string]interface{}, error) {
	if d.err != nil {
		return nil, d.err
	}

	respMap := make(map[string]interface{})
	err := json.Unmarshal(d.resp, &respMap)
	if err != nil {
		return nil, err
	}

	return respMap, nil
}

func (d *defaultHttpClient) Error() error {
	return d.err
}

func (d *defaultHttpClient) validate() error {
	if !d.method.WasSupported() {
		return ErrMethodNotSupport
	}

	if d.host == "" {
		return ErrHostIsEmpty
	}

	d.host = strings.TrimPrefix(d.host, "https://")
	d.host = strings.TrimPrefix(d.host, "http://")

	if d.url == "" {
		return ErrURLIsEmpty
	}

	return nil
}

func (d *defaultHttpClient) complete() {
	if !strings.HasPrefix(d.url, "/") {
		d.url = "/" + d.url
	}

	d.url = strings.TrimSuffix(d.url, "/")

	return
}
