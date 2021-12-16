package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"math"

	"github.com/YaroslavGaponov/cpuminer/internal/miner"
	"github.com/YaroslavGaponov/cpuminer/pkg/bitcoin"
)

func main() {

	fileName := flag.String("file", "", "file with block")
	from := flag.Uint("from", 0, "nonce from")
	to := flag.Uint("to", math.MaxUint32, "nonce to")

	flag.Parse()

	file, err := ioutil.ReadFile(*fileName)
	if err != nil {
		panic(err)
	}

	var block bitcoin.Block
	if err := json.Unmarshal(file, &block); err != nil {
		panic(err)
	}

	fmt.Printf("Nonce from %d to %d\n", *from, *to)

	m := miner.New(block)
	if nonce, err := m.Mine(uint32(*from), uint32(*to)); err != nil {
		fmt.Printf("Error %v", err)
	} else {
		fmt.Printf("Nonce is %d", nonce)
	}
}
