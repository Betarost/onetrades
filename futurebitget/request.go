package futurebitget

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"

	"github.com/Betarost/onetrades/utils"
)

func CreateFullURL(r *utils.Request) error {
	fullURL := fmt.Sprintf("%s%s", r.BaseURL, r.Endpoint)
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
		sf, err := utils.SignFunc(utils.KeyTypeHmacBase64)
		if err != nil {
			return err
		}

		var payload strings.Builder
		payload.WriteString(fmt.Sprintf("%d", r.Timestamp))
		payload.WriteString(r.Method)
		payload.WriteString(fmt.Sprintf("%s?%s", r.Endpoint, r.QueryString))
		if r.BodyString != "" && r.BodyString != "?" {
			payload.WriteString(r.BodyString)
		}
		// raw := fmt.Sprintf("%d%s%s?%s", r.Timestamp, r.Method, r.Endpoint, r.QueryString)
		raw := payload.String()
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
		header.Set("ACCESS-KEY", r.TmpApi)
		header.Set("ACCESS-SIGN", r.Sign)
		header.Set("ACCESS-PASSPHRASE", r.TmpMemo)
		header.Set("ACCESS-TIMESTAMP", fmt.Sprintf("%d", r.Timestamp))
		header.Set("locale", "en-US")
	}
	r.TmpApi = ""
	r.TmpMemo = ""
	r.Header = header
	return nil
}

func CreateReq(r *utils.Request) error {
	return nil
}
