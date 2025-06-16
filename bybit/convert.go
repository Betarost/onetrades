package bybit

import (
	"fmt"
	"strings"

	"github.com/Betarost/onetrades/entity"
)

func convertAccountInfo(in accountInfo) (out entity.AccountInformation) {

	out.UID = fmt.Sprintf("%d", in.UserID)
	out.Label = in.Note
	out.IP = strings.Join(in.Ips, ",")
	out.CanRead = true

	if in.ReadOnly == 0 {
		out.CanTrade = true
	}

	for _, item := range in.Permissions.Spot {
		if item == "SpotTrade" {
			out.PermSpot = true
			break
		}
	}

	for _, item := range in.Permissions.Derivatives {
		if item == "DerivativesTrade" {
			out.PermFutures = true
			break
		}
	}

	for _, item := range in.Permissions.Wallet {
		if item == "AccountTransfer" {
			out.CanTransfer = true
			break
		}
	}

	return out
}
