package main

import (
	"fmt"

	"github.com/YaroslavGaponov/cpuminer/internal/miner"
	"github.com/YaroslavGaponov/cpuminer/pkg/bitcoin"
)

func main() {

	block := bitcoin.Block{
		Hash:          "00000000000000001e8d6829a8a21adc5d38d0a473b144b6765798e61f98bd1d",
		Version:       1,
		PrevBlockHash: "00000000000008a3a41b85b8b29ad444def299fee21793cd8b9e567eab02cd81",
		MerkleRoot:    "2b12fcf1b09288fcaff797d71e950e71ae42b91e8bdb2304758dfcffc2b620e3",
		Time:          1305998791,
		Bits:          440711666,
		Nonce:         2504433986,
	}

	m := miner.New(block)
	if nonce, err := m.Mine(2400000000, 2600000000); err != nil {
		fmt.Printf("Error %v", err)
	} else {
		fmt.Printf("Nonce is %d", nonce)
	}
}
