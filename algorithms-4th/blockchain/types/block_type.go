package types

import "golang.org/x/crypto/sha3"

type Hash [32]byte
type Tx string
type BlockCommitment string

var EmptyHash = sha3.Sum256(nil)

type BlockHeader struct {
	Version       uint64
	Height        uint64
	PrevBlockHash Hash
	Timestamp     uint64
	Nouce         uint64
	Bits          uint64
	BlockCommitment
}

type Block struct {
	BlockHeader
	Transactions []*Tx
}

type Tx struct {
	TxData
}

type TxData struct {
	Version        uint64
	SerializedSize uint64
	TimeRange      uint64
	Inputs         []*TxInput
	Outputs        []*TxOutput
}
