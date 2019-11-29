package wallet_interface

import "strings"

// CoinType represents a cryptocurrency that has been
// implemented the wallet interface.
type CoinType string

// CurrencyCode returns the coins currency code.
func (ct CoinType) CurrencyCode() string {
	return strings.ToUpper(string(ct))
}

const (
	// Mainnet
	CtMock        = "MCK"
	CtBitcoin     = "BTC"
	CtBitcoinCash = "BCH"
	CtLitecoin    = "LTC"
	CtZCash       = "ZEC"
	CtEthereum    = "ETH"
	CtMonero      = "XMR"
	CtDash        = "DASH"
)

var Bip44Indexes = map[CoinType]uint32{
	CtMock:        1,
	CtBitcoin:     0,
	CtBitcoinCash: 145,
	CtLitecoin:    2,
	CtZCash:       133,
	CtEthereum:    60,
	CtMonero:      128,
	CtDash:        5,
}

var codeMap = map[string]CoinType{
	"MCK":  CtMock,
	"BTC":  CtBitcoin,
	"BCH":  CtBitcoinCash,
	"LTC":  CtLitecoin,
	"ZEC":  CtZCash,
	"ETH":  CtEthereum,
	"XMR":  CtMonero,
	"DASH": CtDash,
}
