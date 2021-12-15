package miner

import (
	"errors"
	"strings"

	"github.com/YaroslavGaponov/cpuminer/pkg/bitcoin"
	"github.com/YaroslavGaponov/cpuminer/pkg/progressbar"
)

var errNotFound = errors.New("hash is not found")

type Miner struct {
	block bitcoin.Block
}

func New(block bitcoin.Block) *Miner {
	return &Miner{
		block: block,
	}
}

func (m *Miner) Mine(from, to uint32) (uint32, error) {
	bar := progressbar.New("Mining", int(from), int(to))
	bar.Begin()
	defer bar.Done()
	for nonce := from; nonce < to; nonce++ {
		if hash, err := bitcoin.CalcHash(m.block, nonce); err == nil {
			if strings.HasPrefix(hash, "0000000000000000") {
				return nonce, nil
			}
		}
		bar.Update(int(nonce))
	}
	return 0, errNotFound
}
