package onetrades

import (
	"github.com/Betarost/onetrades/optionokx"
)

func NewOptionOKXClient(apiKey, secretKey, memo string) *optionokx.Client {
	return optionokx.NewClient(apiKey, secretKey, memo)
}
