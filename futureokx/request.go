package futureokx

import (
	"bytes"
	"fmt"
	"net/http"
	"time"

	"github.com/Betarost/onetrades/utils"
)

func CreateFullURL(r *utils.Request) error {
	fullURL := fmt.Sprintf("%s%s", r.BaseURL, r.Endpoint)
	queryString := r.Query.Encode()
	if queryString != "" {
		fullURL = fmt.Sprintf("%s?%s", fullURL, queryString)
	}
	r.QueryString = queryString
	r.FullURL = fullURL
	return nil
}

func CreateBody(r *utils.Request) error {
	body := &bytes.Buffer{}
	bodyString := r.Form.Encode()
	if bodyString != "" {
		r.Header.Set("Content-Type", "application/json")
		body = bytes.NewBufferString(bodyString)
	}
	r.BodyString = bodyString
	r.Body = body
	return nil
}

func CreateSign(r *utils.Request) error {
	if r.SecType == utils.SecTypeSigned {
		r.Timestamp = utils.CurrentTimestamp() - r.TimeOffset

		sf, err := utils.SignFunc(utils.KeyTypeHmacBase64)
		if err != nil {
			return err
		}
		path := r.Endpoint
		if r.QueryString != "" {
			path = fmt.Sprintf("%s?%s", path, r.QueryString)
		}
		raw := fmt.Sprintf("%s%s%s%s", time.UnixMilli(r.Timestamp).UTC().Format("2006-01-02T15:04:05.999Z07:00"), r.Method, path, r.BodyString)
		sign, err := sf(r.TmpSig, raw)
		if err != nil {
			return err
		}
		r.TmpSig = ""
		r.Sign = *sign
	}
	return nil
}

func CreateHeaders(r *utils.Request) error {
	header := http.Header{}
	if r.Header != nil {
		header = r.Header.Clone()
	}

	header.Set("Content-Type", "application/json")
	if r.SecType == utils.SecTypeSigned {
		header.Set("OK-ACCESS-KEY", r.TmpApi)
		header.Set("OK-ACCESS-PASSPHRASE", r.TmpMemo)
		header.Set("OK-ACCESS-SIGN", r.Sign)
		header.Set("OK-ACCESS-TIMESTAMP", time.UnixMilli(r.Timestamp).UTC().Format("2006-01-02T15:04:05.999Z07:00"))
	}
	r.TmpApi = ""
	r.TmpMemo = ""
	r.Header = header
	return nil
}

func CreateReq(r *utils.Request) error {
	return nil
}
