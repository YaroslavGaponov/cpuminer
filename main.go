package main

import (
	"fmt"
	"strings"
	"time"

	"./cpuminer"
)

type Block struct {
	Hash          string
	Version       uint32
	PrevBlockHash string
	MerkleRoot    string
	Time          uint32
	Bits          uint32
	Nonce         uint32
}

func progress() func(start, end, current uint32) {
	s := time.Now()
	var percent int = 0
	return func(start, end, current uint32) {
		percent2 := int(float64(current-start) / float64(end-start) * 100)
		if percent2 > percent {
			percent = percent2
			speed := uint32(float64(current-start) / time.Since(s).Seconds())
			indicator := "[" + strings.Repeat("=", percent) + ">" + strings.Repeat(" ", 100-percent) + "]"
			fmt.Printf("\rProgress %s %d %% | %d ops/sec", indicator, percent, speed)
		}
	}
}

func main() {

	block := Block{
		Hash:          "00000000000000001e8d6829a8a21adc5d38d0a473b144b6765798e61f98bd1d",
		Version:       1,
		PrevBlockHash: "00000000000008a3a41b85b8b29ad444def299fee21793cd8b9e567eab02cd81",
		MerkleRoot:    "2b12fcf1b09288fcaff797d71e950e71ae42b91e8bdb2304758dfcffc2b620e3",
		Time:          1305998791,
		Bits:          440711666,
		Nonce:         2504433986,
	}

	fmt.Println("Starting mining...")

	var start, end uint32 = 2400000000, 2600000000
	fmt.Printf("Mining nonce from %d to %d\n", start, end)

	if hash, nonce, err := cpuminer.Mine(progress(), block.Version, block.PrevBlockHash, block.MerkleRoot, block.Time, block.Bits, start, end); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("\nHash is %s\nNonce is %d", hash, nonce)
	}
}
