package futuremexc

import "fmt"

type APIError struct {
	Success  bool   `json:"success"`
	Code     int64  `json:"code"`
	Message  string `json:"message"`
	Response []byte `json:"-"`
}

func (e APIError) Error() string {
	if e.IsValid() {
		return fmt.Sprintf("<APIError> code=%d, msg=%s", e.Code, e.Message)
	}
	return fmt.Sprintf("<APIError> rsp=%s", string(e.Response))
}

func (e APIError) IsValid() bool {
	return e.Success
}
