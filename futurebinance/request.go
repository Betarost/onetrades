package futurebinance

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"

	"github.com/Betarost/onetrades/utils"
)

func CreateFullURL(r *utils.Request) error {
	fullURL := fmt.Sprintf("%s%s", r.BaseURL, r.Endpoint)
	if r.RecvWindow > 0 {
		r.SetParam("recvWindow", r.RecvWindow)
	}
	r.Timestamp = utils.CurrentTimestamp() - r.TimeOffset
	if r.SecType == utils.SecTypeSigned {
		r.SetParam("timestamp", r.Timestamp)
	}
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
		sf, err := utils.SignFunc(utils.KeyTypeHmac)
		if err != nil {
			return err
		}
		raw := fmt.Sprintf("%s%s", r.QueryString, r.BodyString)
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
		header.Set("X-MBX-APIKEY", r.TmpApi)
	}
	r.TmpApi = ""
	r.Header = header
	return nil
}

func CreateReq(r *utils.Request) error {
	if r.SecType == utils.SecTypeSigned {
		fullURL := fmt.Sprintf("%s%s", r.BaseURL, r.Endpoint)
		v := url.Values{}
		v.Set("signature", r.Sign)
		if r.QueryString == "" {
			r.QueryString = v.Encode()
		} else {
			r.QueryString = fmt.Sprintf("%s&%s", r.QueryString, v.Encode())
		}
		fullURL = fmt.Sprintf("%s?%s", fullURL, r.QueryString)
		r.FullURL = fullURL
	}
	return nil
}
