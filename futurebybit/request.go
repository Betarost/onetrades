package futurebybit

import (
	"bytes"
	"fmt"
	"net/http"

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
		r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		body = bytes.NewBufferString(bodyString)
	}
	r.BodyString = bodyString
	r.Body = body
	return nil
}

func CreateSign(r *utils.Request) error {
	if r.SecType == utils.SecTypeSigned {
		r.Timestamp = utils.CurrentTimestamp() - r.TimeOffset

		sf, err := utils.SignFunc(utils.KeyTypeHmac)
		if err != nil {
			return err
		}
		raw := fmt.Sprintf("%d%s5000%s", r.Timestamp, r.TmpApi, r.QueryString)
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
	if r.BodyString != "" {
		header.Set("Content-Type", "application/x-www-form-urlencoded")
	}
	if r.SecType == utils.SecTypeSigned {
		header.Set("X-BAPI-API-KEY", r.TmpApi)
		header.Set("X-BAPI-SIGN", r.Sign)
		header.Set("X-BAPI-TIMESTAMP", fmt.Sprintf("%d", r.Timestamp))
		header.Set("X-BAPI-RECV-WINDOW", "5000")
	}
	r.TmpApi = ""
	r.Header = header
	return nil
}

func CreateReq(r *utils.Request) error {
	return nil
}
