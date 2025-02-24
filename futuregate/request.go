package futuregate

import (
	"bytes"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"net/http"
	"time"

	"github.com/Betarost/onetrades/utils"
)

func CreateFullURL(r *utils.Request) error {
	fullURL := fmt.Sprintf("%s%s", r.BaseURL, r.Endpoint)
	r.Timestamp = time.Now().Unix() - r.TimeOffset
	// if r.SecType == utils.SecTypeSigned {
	// 	r.SetParam("timestamp", r.Timestamp)
	// }
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
		h := sha512.New()
		hashedPayload := hex.EncodeToString(h.Sum(nil))
		raw := fmt.Sprintf("%s\n%s\n%s\n%s\n%d", r.Method, r.Endpoint, r.QueryString, hashedPayload, r.Timestamp)

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

	header.Set("Accept", "application/json")
	header.Set("Content-Type", "application/json")
	if r.SecType == utils.SecTypeSigned {

		header.Set("KEY", r.TmpApi)
		header.Set("SIGN", r.Sign)
		header.Set("Timestamp", fmt.Sprintf("%d", r.Timestamp))
	}
	r.TmpApi = ""
	r.Header = header
	return nil
}

func CreateReq(r *utils.Request) error {
	return nil
}
