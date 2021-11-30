package cpuminer

import (
	"crypto/sha256"
	"errors"
	"strings"
	t "time"
)

const hextable = "0123456789abcdef"

var (
	errBadBlock = errors.New("bad block")
	errNotFound = errors.New("not found")
)

func CalcHash(version uint32, prevBlockHash string, merkleRoot string, time uint32, bits uint32, nonce uint32) (string, error) {

	if len(prevBlockHash) != 64 {
		return "", errBadBlock
	}

	if len(merkleRoot) != 64 {
		return "", errBadBlock
	}

	var header [80]byte

	putUint32(header[:], uint32(version))
	putHashString(header[4:], prevBlockHash)
	putHashString(header[36:], merkleRoot)
	putUint32(header[68:], uint32(time))
	putUint32(header[72:], uint32(bits))
	putUint32(header[76:], uint32(nonce))

	h := sha256.New()
	if _, err := h.Write(header[:]); err != nil {
		return "", err
	}

	h2 := sha256.New()
	if _, err := h2.Write(h.Sum(nil)); err != nil {
		return "", err
	}

	bytes := h2.Sum(nil)

	hash := encodeToString(bytes)

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

func encodeToString(b []byte) string {
	var s string
	for i := len(b) - 1; i >= 0; i-- {
		s += byteToHex(b[i])
	}
	return s
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

func byteToHex(b byte) string {
	return string(hextable[b>>4]) + string(hextable[b&0x0f])
}

type Result struct {
	Hash  string
	Nonce uint32
}

func mineOne(done chan Result, version uint32, prevBlockHash string, merkleRoot string, time uint32, bits uint32, nonce uint32) {
	hash, _ := CalcHash(version, prevBlockHash, merkleRoot, time, bits, nonce)
	if strings.HasPrefix(hash, "0000000000000000") {
		done <- Result{Hash: hash, Nonce: nonce}
	}
}

func Mine(progress func(start, end, current uint32), version uint32, prevBlockHash string, merkleRoot string, time uint32, bits uint32, nonceStart uint32, nonceEnd uint32) (string, uint32, error) {
	done := make(chan Result)

	for nonce := nonceStart; nonce < nonceEnd; nonce++ {
		select {
		case result := <-done:
			return result.Hash, result.Nonce, nil
		default:
			progress(nonceStart, nonceEnd, nonce)
			go mineOne(done, version, prevBlockHash, merkleRoot, time, bits, nonce)
		}
	}

	select {
	case result := <-done:
		return result.Hash, result.Nonce, nil
	case <-t.After(t.Second * 10):
		return "", 0, errNotFound
	}
}
