package bitcoin

import (
	"crypto/sha256"
	"errors"
)

const hextable = "0123456789abcdef"

var (
	errBadBlock = errors.New("bad block")
	errNotFound = errors.New("not found")
)

type Block struct {
	Hash          string `json:"hash"`
	Version       uint32 `json:"ver"`
	PrevBlockHash string `json:"prev_hash"`
	MerkleRoot    string `json:"mrkl_root"`
	Time          uint32 `json:"time"`
	Bits          uint32 `json:"bits"`
	Nonce         uint32 `json:"nonce"`
}

func CalcHash(block Block, nonce uint32) ([]byte, error) {

	if len(block.PrevBlockHash) != 64 {
		return nil, errBadBlock
	}

	if len(block.MerkleRoot) != 64 {
		return nil, errBadBlock
	}

	var header [80]byte

	putUint32(header[:], block.Version)
	putHashString(header[4:], block.PrevBlockHash)
	putHashString(header[36:], block.MerkleRoot)
	putUint32(header[68:], block.Time)
	putUint32(header[72:], block.Bits)
	putUint32(header[76:], nonce)

	h := sha256.New()
	if _, err := h.Write(header[:]); err != nil {
		return nil, err
	}

	h2 := sha256.New()
	if _, err := h2.Write(h.Sum(nil)); err != nil {
		return nil, err
	}

	hash := h2.Sum(nil)

	return hash, nil
}

func putUint32(b []byte, v uint32) {
	b[3] = byte(v >> 24)
	b[2] = byte(v >> 16)
	b[1] = byte(v >> 8)
	b[0] = byte(v)
}

func putHashString(r []byte, s string) {
	for i, j := len(s)-1, 0; i > 0; i, j = i-2, j+1 {
		r[j] = (hexToByte(s[i-1]) << 4) | hexToByte(s[i])
	}
}

func hexToByte(c byte) byte {
	if c >= '0' && c <= '9' {
		return c - '0'
	}
	if c >= 'a' && c <= 'f' {
		return c - 'a' + 10
	}
	if c >= 'A' && c <= 'F' {
		return c - 'A' + 10
	}
	return 0
}
