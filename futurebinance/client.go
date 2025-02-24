package futurebinance

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Betarost/onetrades/utils"
)

var (
	TradeName = "BINANCE"
)

type Client struct {
	apiKey     string
	secretKey  string
	KeyType    string
	BaseURL    string
	UserAgent  string
	Debug      bool
	Logger     *log.Logger
	TimeOffset int64
}

func (c *Client) debug(format string, v ...interface{}) {
	if c.Debug {
		c.Logger.Printf(format, v...)
	}
}

func (c *Client) callAPI(ctx context.Context, r *utils.Request, opts ...utils.RequestOption) (data []byte, header *http.Header, err error) {
	r.TmpApi = c.apiKey
	r.TmpSig = c.secretKey
	opts = append(opts, CreateFullURL, CreateBody, CreateSign, CreateHeaders, CreateReq)
	err = r.ParseRequest(opts...)
	if err != nil {
		return []byte{}, &http.Header{}, err

	}
	c.debug("FullURL %s\n", r.FullURL)
	c.debug("Body %s\n", r.BodyString)
	c.debug("Sign %s\n", r.Sign)
	c.debug("Headers %+v\n", r.Header)
	req, err := http.NewRequest(r.Method, r.FullURL, r.Body)
	if err != nil {
		return []byte{}, &http.Header{}, err
	}
	req = req.WithContext(ctx)
	req.Header = r.Header
	c.debug("request: %#v\n", req)

	res, err, conn := r.DoFunc(req)
	if err != nil {
		return []byte{}, &http.Header{}, err
	}
	data, err = r.ReadAllBody(res.Body, res.Header.Get("Content-Encoding"))
	if err != nil {
		return []byte{}, &http.Header{}, err
	}
	defer func() {
		cerr := res.Body.Close()
		if conn != nil {
			conn.Close()
		}
		if err == nil && cerr != nil {
			err = cerr
		}
	}()

	c.debug("response: %#v\n", res)
	c.debug("response body: %s\n", string(data))
	c.debug("response status code: %d\n", res.StatusCode)

	// if res.StatusCode >= http.StatusBadRequest {
	// 	apiErr := new(APIError)
	// 	e := json.Unmarshal(data, apiErr)
	// 	if e != nil {
	// 		c.debug("failed to unmarshal json: %s\n", e)
	// 	}
	// 	if !apiErr.IsValid() {
	// 		apiErr.Response = data
	// 	}
	// 	return nil, &res.Header, apiErr
	// }

	// apiErr := new(APIError)
	// e := json.Unmarshal(data, apiErr)
	// if e != nil {
	// 	c.debug("failed to unmarshal json: %s\n", e)
	// }
	// if !apiErr.IsValid() {
	// 	c.debug("==sss====: %+v\n", apiErr)
	// 	apiErr.Response = data
	// 	return nil, &res.Header, apiErr
	// }
	return data, &res.Header, nil
}

func NewClient(apiKey, secretKey string) *Client {
	return &Client{
		apiKey:    apiKey,
		secretKey: secretKey,
		KeyType:   utils.KeyTypeHmac,
		BaseURL:   utils.GetApiEndpoint(TradeName),
		UserAgent: "Onetrades/golang",
		Logger:    log.New(os.Stderr, fmt.Sprintf("%s-onetrades ", TradeName), log.LstdFlags),
	}
}

func (c *Client) NewGetAccountService() *GetAccountService {
	return &GetAccountService{c: c}
}

func (c *Client) NewGetAccountBalance() *GetAccountBalance {
	return &GetAccountBalance{c: c}
}
