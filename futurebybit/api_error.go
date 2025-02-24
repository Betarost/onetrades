package futurebybit

import "fmt"

type APIError struct {
	Code     int64  `json:"retCode"`
	Message  string `json:"retMsg"`
	Response []byte `json:"-"`
}

func (e APIError) Error() string {
	if e.IsValid() {
		return fmt.Sprintf("<APIError> code=%d, msg=%s", e.Code, e.Message)
	}
	return fmt.Sprintf("<APIError> rsp=%s", string(e.Response))
}

func (e APIError) IsValid() bool {
	return e.Message == "OK" && e.Code == 0
}
