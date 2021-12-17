package miner

import (
	"errors"
	"runtime"

	"github.com/YaroslavGaponov/cpuminer/pkg/bitcoin"
	"github.com/YaroslavGaponov/cpuminer/pkg/progressbar"
)

var errNotFound = errors.New("hash is not found")

type Miner struct {
	zbytes int
	zbits  int
	block  bitcoin.Block
}

func New(block bitcoin.Block, zbits int) *Miner {
	return &Miner{
		zbytes: zbits / 8,
		zbits:  zbits % 8,
		block:  block,
	}
}

func (m *Miner) mine(block bitcoin.Block, in chan uint32, out chan uint32) {
main:
	for {
		select {
		case nonce := <-in:
			if hash, err := bitcoin.CalcHash(block, nonce); err == nil {
				for i, j := 0, len(hash)-1; i < m.zbytes; i, j = i+1, j-1 {
					if hash[j] != 0 {
						continue main
					}
				}
				if (hash[len(hash)-m.zbytes-1] >> m.zbits) != 0 {
					continue main
				}
				out <- nonce
			}
		}
	}
}

func (m *Miner) Mine(from, to uint32) (uint32, error) {
	bar := progressbar.New("Mining", int(from), int(to))
	bar.Begin()
	defer bar.Done()

	in := make(chan uint32)
	out := make(chan uint32)
	for i := 0; i < runtime.NumCPU()<<1; i++ {
		go m.mine(m.block, in, out)
	}

	for nonce := from; nonce < to; nonce++ {
		select {
		case found := <-out:
			return found, nil
		default:
			in <- nonce
			bar.Update(int(nonce))
		}
	}
	return 0, errNotFound
}
