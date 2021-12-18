package miner

import (
	"errors"
	"runtime"

	"github.com/YaroslavGaponov/cpuminer/pkg/bitcoin"
	"github.com/YaroslavGaponov/cpuminer/pkg/progressbar"
)

const SIZE = 500

var errNotFound = errors.New("hash is not found")

type Miner struct {
	zbytes int
	zbits  int
	block  bitcoin.Block
}

type Task struct {
	start uint32
	end   uint32
}

func New(block bitcoin.Block, zbits int) *Miner {
	return &Miner{
		zbytes: zbits / 8,
		zbits:  zbits % 8,
		block:  block,
	}
}

func (m *Miner) mine(block bitcoin.Block, in chan Task, out chan uint32) {
	bc := bitcoin.New(block)
	for {
		select {
		case task := <-in:
			main: for nonce := task.start; nonce < task.end; nonce++ {
				if hash, err := bc.CalcHash(nonce); err == nil {
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
}

func (m *Miner) Mine(from, to uint32) (uint32, error) {
	bar := progressbar.New("Mining", int(from), int(to))
	bar.Begin()
	defer bar.Done()

	in := make(chan Task)
	out := make(chan uint32)
	for i := 0; i < runtime.NumCPU(); i++ {
		go m.mine(m.block, in, out)
	}

	for nonce := from; nonce < to; nonce += SIZE {
		select {
		case found := <-out:
			return found, nil
		default:
			in <- Task{start: nonce, end: nonce + SIZE}
			bar.Update(int(nonce))
		}
	}
	return 0, errNotFound
}
