package wallet_interface

import "encoding/json"

// Address represents a cryptocurrency address used by OpenBazaar.
type Address struct {
	addr string
	typ  CoinType
}

// NewAddress return a new Address.
func NewAddress(addr string, typ CoinType) Address {
	return Address{addr, typ}
}

// String returns the address's string representation.
func (a Address) String() string {
	return a.addr
}

// CoinType returns the addresses type.
func (a Address) CoinType() CoinType {
	return a.typ
}

// MarshalJSON is used to marshal the address to JSON.
func (a Address) MarshalJSON() ([]byte, error) {
	type addrJSON struct {
		Addr string `json:"address"`
		Typ  string `json:"cointype"`
	}

	c0 := addrJSON{
		Addr: a.addr,
		Typ:  a.typ.CurrencyCode(),
	}
	return json.Marshal(c0)
}

// UnmarshalJSON is used to unmarshal the address from JSON.
func (a Address) UnmarshalJSON(b []byte) error {
	type addrJSON struct {
		Addr string `json:"address"`
		Typ  string `json:"cointype"`
	}

	var c0 addrJSON
	err := json.Unmarshal(b, &c0)
	typ, ok := codeMap[c0.Typ]
	if !ok {
		typ = CoinType(c0.Typ)
	}
	if err == nil {
		a.addr = c0.Addr
		a.typ = typ
	}

	return err
}
