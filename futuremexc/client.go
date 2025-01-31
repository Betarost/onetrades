package futuremexc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	BaseApiMainUrl = "https://contract.mexc.com"
)

const (
	apikeyHeader    = "ApiKey"
	timestampKey    = "timestamp"
	timestampHeader = "Request-Time"
	signatureKey    = "signature"
	signatureHeader = "signature"
	recvWindowKey   = "recvWindow"
)

func currentTimestamp() int64 {
	return time.Now().UnixMilli()
	// return int64(time.Nanosecond) * time.Now().UnixNano() / int64(time.Millisecond)
}

func getApiEndpoint() string {
	return BaseApiMainUrl
}

type doFunc func(req *http.Request) (*http.Response, error)

type Client struct {
	APIKey     string
	SecretKey  string
	KeyType    string
	BaseURL    string
	UserAgent  string
	HTTPClient *http.Client
	Debug      bool
	Logger     *log.Logger
	TimeOffset int64
	do         doFunc
}

func (c *Client) debug(format string, v ...interface{}) {
	if c.Debug {
		c.Logger.Printf(format, v...)
	}
}

func (c *Client) parseRequest(r *request, opts ...RequestOption) (err error) {
	// set request options from user
	for _, opt := range opts {
		opt(r)
	}
	err = r.validate()
	if err != nil {
		return err
	}

	fullURL := fmt.Sprintf("%s%s", c.BaseURL, r.endpoint)
	timestamp := fmt.Sprintf("%d", currentTimestamp()-c.TimeOffset)

	if r.recvWindow > 0 {
		r.setParam(recvWindowKey, r.recvWindow)
	}
	// if r.secType == secTypeSigned {
	// 	r.setParam(timestampKey, timestamp)
	// }
	queryString := r.query.Encode()
	body := &bytes.Buffer{}
	bodyString := r.form.Encode()
	header := http.Header{}
	if r.header != nil {
		header = r.header.Clone()
	}
	header.Set("Content-Type", "application/json")
	header.Set(timestampHeader, timestamp)

	if bodyString != "" {
		header.Set("Content-Type", "application/json")
		body = bytes.NewBufferString(bodyString)
	}
	if r.secType == secTypeAPIKey || r.secType == secTypeSigned {
		header.Set(apikeyHeader, c.APIKey)
	}
	kt := c.KeyType
	if kt == "" {
		kt = KeyTypeHmac
	}
	sf, err := SignFunc(kt)
	if err != nil {
		return err
	}
	if r.secType == secTypeSigned {
		// raw := fmt.Sprintf("%s%s", queryString, bodyString)
		raw := fmt.Sprintf("%s%s%s", c.APIKey, timestamp, bodyString)
		sign, err := sf(c.SecretKey, raw)
		if err != nil {
			return err
		}
		header.Set(signatureHeader, *sign)
		// v := url.Values{}
		// v.Set(signatureKey, *sign)
		// if queryString == "" {
		// 	queryString = v.Encode()
		// } else {
		// 	queryString = fmt.Sprintf("%s&%s", queryString, v.Encode())
		// }
	}
	if queryString != "" {
		fullURL = fmt.Sprintf("%s?%s", fullURL, queryString)
	}
	c.debug("full url: %s, body: %s\n", fullURL, bodyString)

	r.fullURL = fullURL
	r.header = header
	r.body = body
	return nil
}

func (c *Client) callAPI(ctx context.Context, r *request, opts ...RequestOption) (data []byte, header *http.Header, err error) {
	err = c.parseRequest(r, opts...)
	if err != nil {
		return []byte{}, &http.Header{}, err
	}
	req, err := http.NewRequest(r.method, r.fullURL, r.body)
	if err != nil {
		return []byte{}, &http.Header{}, err
	}
	req = req.WithContext(ctx)
	req.Header = r.header
	c.debug("request: %#v\n", req)
	f := c.do
	if f == nil {
		f = c.HTTPClient.Do
	}
	res, err := f(req)
	if err != nil {
		return []byte{}, &http.Header{}, err
	}
	data, err = io.ReadAll(res.Body)
	if err != nil {
		return []byte{}, &http.Header{}, err
	}
	defer func() {
		cerr := res.Body.Close()
		// Only overwrite the returned error if the original error was nil and an
		// error occurred while closing the body.
		if err == nil && cerr != nil {
			err = cerr
		}
	}()
	c.debug("response: %#v\n", res)
	c.debug("response body: %s\n", string(data))
	c.debug("response status code: %d\n", res.StatusCode)

	if res.StatusCode >= http.StatusBadRequest {
		apiErr := new(APIError)
		e := json.Unmarshal(data, apiErr)
		if e != nil {
			c.debug("failed to unmarshal json: %s\n", e)
		}
		if !apiErr.IsValid() {
			apiErr.Response = data
		}
		return nil, &res.Header, apiErr
	}

	apiErr := new(APIError)
	e := json.Unmarshal(data, apiErr)
	if e != nil {
		c.debug("failed to unmarshal json: %s\n", e)
	}
	if !apiErr.IsValid() {
		c.debug("==sss====: %+v\n", apiErr)
		apiErr.Response = data
		return nil, &res.Header, apiErr
	}
	return data, &res.Header, nil
}

func NewClient(apiKey, secretKey string) *Client {
	return &Client{
		APIKey:     apiKey,
		SecretKey:  secretKey,
		KeyType:    KeyTypeHmac,
		BaseURL:    getApiEndpoint(),
		UserAgent:  "Onetrades/golang",
		HTTPClient: http.DefaultClient,
		Logger:     log.New(os.Stderr, "Mexc-onetrades ", log.LstdFlags),
	}
}

func (c *Client) NewGetAccountAssets() *GetAccountAssets {
	return &GetAccountAssets{c: c}
}
