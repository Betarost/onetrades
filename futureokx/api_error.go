package futureokx

import "fmt"

type APIError struct {
	Code     string `json:"code"`
	Message  string `json:"msg"`
	Response []byte `json:"-"`
}

func (e APIError) Error() string {
	if e.IsValid() {
		return fmt.Sprintf("<APIError> code=%s, msg=%s", e.Code, e.Message)
	}
	return fmt.Sprintf("<APIError> rsp=%s", string(e.Response))
}

func (e APIError) IsValid() bool {
	return e.Message == "" && e.Code == "0"
}
