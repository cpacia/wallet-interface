package wallet_interface

import (
	"encoding/hex"
	"encoding/json"
)

// TransactionID represents an ID for a transaction made by the wallet
type TransactionID string

// String returns the string representation of the ID.
func (t *TransactionID) String() string {
	return string(*t)
}

// Transaction is a basic record which is used to convey information
// about a transaction to OpenBazaar. It's designed to be generic
// enough to be used by a variety of different coins.
//
// In the case of multisig transactions OpenBazaar will be using the
// To spend info objects in the from field when spending from a
// multisig.
type Transaction struct {
	ID TransactionID

	From []SpendInfo
	To   []SpendInfo

	Height uint64
}

// SpendInfo represents a transaction data element. This could either
// be an input or an outpoint in the Bitcoin context. The ID can
// be used by the wallet to attach metadata needed construct a
// transaction from this data structure. Again in the bitcoin context
// this would be a serialized outpoint (when this represents an input).
type SpendInfo struct {
	ID []byte

	Address Address
	Amount  Amount

	IsRelevant bool
	IsWatched  bool
}

// MarshalJSON is used to marshal the spend info to JSON.
func (si *SpendInfo) MarshalJSON() ([]byte, error) {
	type addrJSON struct {
		Address  string `json:"address"`
		CoinType string `json:"cointype"`
	}

	type siJSON struct {
		ID      string   `json:"id"`
		Address addrJSON `json:"address"`
		Amount  string   `json:"amount"`
	}

	c0 := siJSON{
		ID: hex.EncodeToString(si.ID),
		Address: addrJSON{
			Address:  si.Address.addr,
			CoinType: si.Address.typ.CurrencyCode(),
		},
		Amount: si.Amount.String(),
	}
	return json.Marshal(c0)
}

// UnmarshalJSON is used to unmarshal the spend info from JSON.
func (si *SpendInfo) UnmarshalJSON(b []byte) error {
	type addrJSON struct {
		Address  string `json:"address"`
		CoinType string `json:"cointype"`
	}
	type siJSON struct {
		ID      string   `json:"id"`
		Address addrJSON `json:"address"`
		Amount  string   `json:"amount"`
	}
	var j siJSON
	err := json.Unmarshal(b, &j)
	if err == nil {
		id, err := hex.DecodeString(j.ID)
		if err != nil {
			return err
		}
		si.ID = id
		si.Address = Address{addr: j.Address.Address, typ: CoinType(j.Address.CoinType)}
		si.Amount = NewAmount(j.Amount)
	}
	return err
}

// EscrowSignature represents a signature for an escrow transaction.
type EscrowSignature struct {
	Index     int
	Signature []byte
}
