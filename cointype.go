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

	// Testnet
	CtTestnetMock        = "TMCK"
	CtTestnetBitcoin     = "TBTC"
	CtTestnetBitcoinCash = "TBCH"
	CtTestnetLitecoin    = "TLTC"
	CtTestnetZCash       = "TZEC"
	CtTestnetEthereum    = "TETH"
	CtTestnetMonero      = "TXMR"
	CtTestnetDash        = "TDASH"
)

var codeMap = map[string]CoinType{
	"MCK":   CtMock,
	"BTC":   CtBitcoin,
	"BCH":   CtBitcoinCash,
	"LTC":   CtLitecoin,
	"ZEC":   CtZCash,
	"ETH":   CtEthereum,
	"XMR":   CtMonero,
	"DASH":  CtDash,
	"TMCK":  CtTestnetMock,
	"TBTC":  CtTestnetBitcoin,
	"TBCH":  CtTestnetBitcoinCash,
	"TLTC":  CtTestnetLitecoin,
	"TZEC":  CtTestnetZCash,
	"TETH":  CtTestnetEthereum,
	"TXMR":  CtTestnetMonero,
	"TDASH": CtTestnetDash,
}
