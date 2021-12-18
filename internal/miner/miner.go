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

type Result struct {
	nonce uint32
	hash  string
}

func New(block bitcoin.Block, zbits int) *Miner {
	return &Miner{
		zbytes: zbits / 8,
		zbits:  zbits % 8,
		block:  block,
	}
}

func (m *Miner) mine(block bitcoin.Block, in chan Task, out chan Result) {
	bc := bitcoin.New(block)
	for {
		select {
		case task := <-in:
		main:
			for nonce := task.start; nonce < task.end; nonce++ {
				if hash, err := bc.CalcHash(nonce); err == nil {
					for i, j := 0, len(hash)-1; i < m.zbytes; i, j = i+1, j-1 {
						if hash[j] != 0 {
							continue main
						}
					}
					if (hash[len(hash)-m.zbytes-1] >> m.zbits) != 0 {
						continue main
					}
					out <- Result{nonce: nonce, hash: bytesToHex(hash)}
				}
			}
		}
	}
}

func (m *Miner) Mine(from, to uint32) (uint32, string, error) {
	bar := progressbar.New("Mining", from, to)
	bar.Begin()
	defer bar.Done()

	in := make(chan Task)
	out := make(chan Result)
	for i := 0; i < runtime.NumCPU(); i++ {
		go m.mine(m.block, in, out)
	}

	for nonce := from; nonce < to; nonce += SIZE {
		select {
		case result := <-out:
			return result.nonce, result.hash, nil
		default:
			in <- Task{start: nonce, end: nonce + SIZE}
			bar.Update(nonce)
		}
	}
	return 0, "", errNotFound
}

func bytesToHex(b []byte) string {
	var s string
	for i := len(b) - 1; i >= 0; i-- {
		s += byteToHex(b[i])
	}
	return s
}

func byteToHex(b byte) string {
	const hex = "0123456789abcdef"
	return string(hex[b>>4]) + string(hex[b&0xF])
}
