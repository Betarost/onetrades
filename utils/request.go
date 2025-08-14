package utils

import (
	"compress/gzip"
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type secType int

const (
	SecTypeNone secType = iota
	SecTypeAPIKey
	SecTypeSigned
)

type Params map[string]interface{}

type Request struct {
	Method      string
	BaseURL     string
	Endpoint    string
	Proxy       string
	BrokerID    string
	TimeOffset  int64
	Query       url.Values
	QueryString string
	Form        url.Values
	RecvWindow  int64
	Timestamp   int64
	SecType     secType
	Sign        string
	Header      http.Header
	Body        io.Reader
	BodyString  string
	FullURL     string
	TmpSig      string
	TmpApi      string
	TmpMemo     string
}

type RequestOption func(*Request) error

func (r *Request) SetParam(key string, value interface{}) *Request {
	if r.Query == nil {
		r.Query = url.Values{}
	}
	r.Query.Set(key, fmt.Sprintf("%v", value))
	return r
}

func (r *Request) SetParams(m Params) *Request {
	for k, v := range m {
		r.SetParam(k, v)
	}
	return r
}

func (r *Request) SetFormParam(key string, value interface{}) *Request {
	if r.Form == nil {
		r.Form = url.Values{}
	}
	r.Form.Set(key, fmt.Sprintf("%v", value))
	return r
}

func (r *Request) SetFormParams(m Params) *Request {
	for k, v := range m {
		r.SetFormParam(k, v)
	}
	return r
}

func (r *Request) validate() (err error) {
	if r.Query == nil {
		r.Query = url.Values{}
	}
	if r.Form == nil {
		r.Form = url.Values{}
	}
	return nil
}

func (r *Request) ParseRequest(opts ...RequestOption) (err error) {
	err = r.validate()

	if err != nil {
		return err
	}
	for _, opt := range opts {
		e := opt(r)
		if e != nil {
			return e
		}
	}
	return nil
}

func (c *Request) ReadAllBody(r io.Reader, content string) ([]byte, error) {
	switch content {
	case "gzip":
		reader, err := gzip.NewReader(r)
		if err != nil {
			return nil, err
		}
		defer reader.Close()
		buff, err := ioutil.ReadAll(reader)
		return buff, err
	default:
		return io.ReadAll(r)
	}
}

func (c *Request) DoFunc(req *http.Request) (*http.Response, error, *tls.Conn) {
	if c.Proxy != "" {
		proxyURL, err := url.Parse(c.Proxy)
		if err != nil {
			log.Println("DoFunc Err", err)
		} else {
			transport := &http.Transport{
				Proxy: http.ProxyURL(proxyURL),
			}
			client := &http.Client{
				Transport: transport,
			}
			resp, err := client.Do(req)
			return resp, err, nil
		}

	}

	resp, err := http.DefaultClient.Do(req)
	return resp, err, nil
}
