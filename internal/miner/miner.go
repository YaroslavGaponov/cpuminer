package miner

import (
	"errors"
	"runtime"
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

func mine(block bitcoin.Block, in chan uint32, out chan uint32) {
	for {
		select {
		case nonce := <-in:
			if hash, err := bitcoin.CalcHash(block, nonce); err == nil {
				if strings.HasPrefix(hash, "0000000000000000") {
					out <- nonce
				}
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
		go mine(m.block, in, out)
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
