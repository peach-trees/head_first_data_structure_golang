package common

import (
	"bytes"
	"encoding/gob"
	"hash"
	"math/big"
)

func HashFunc(key interface{}, hash hash.Hash) *big.Int {
	buf := bytes.NewBuffer(nil)
	enc := gob.NewEncoder(buf)
	_ = enc.Encode(key)
	hashBytes := hash.Sum(buf.Bytes())
	return new(big.Int).SetBytes(hashBytes)
}
