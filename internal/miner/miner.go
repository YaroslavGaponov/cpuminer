package miner

import (
	"github.com/YaroslavGaponov/cpuminer/pkg/bitcoin"
	"github.com/YaroslavGaponov/cpuminer/pkg/progressbar"
)

type Miner struct {
	block *bitcoin.Block
}

func New(block *bitcoin.Block) *Miner {
	return &Miner{
		block: block,
	}
}

func (m *Miner) Mine(from, to uint32) error {
	bar := progressbar.New("Mining", int(from), int(to))
	bar.Begin()
	for nonce := from; nonce < to; nonce++ {
		m.block.Nonce = nonce
		bitcoin.CalcHash(m.block)
		bar.Update(int(nonce))
	}
	bar.Done()
	return nil
}
