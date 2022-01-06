package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"time"
	"runtime"

	"github.com/YaroslavGaponov/cpuminer/internal/miner"
	"github.com/YaroslavGaponov/cpuminer/pkg/bitcoin"
)

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	fileName := flag.String("file", "", "file with block")
	url := flag.String("url", "", "url with block")
	from := flag.Uint("from", 0, "nonce from")
	to := flag.Uint("to", math.MaxUint32, "nonce to")
	zerobites := flag.Int("zerobits", 13*4, "zerobits")

	flag.Parse()

	var data []byte
	var err error

	if len(*fileName) > 0 {
		data, err = ioutil.ReadFile(*fileName)
		if err != nil {
			panic(err)
		}
	} else if len(*url) > 0 {
		resp, err := http.Get(*url)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		data, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
	} else {
		fmt.Println("You must specify file or url")
		return
	}

	var block bitcoin.Block
	if err := json.Unmarshal(data, &block); err != nil {
		panic(err)
	}

	fmt.Printf("Nonce from %d to %d, zerobits is %d\n", *from, *to, *zerobites)

	start := time.Now()
	m := miner.New(block, *zerobites)
	if nonce, hash, err := m.Mine(uint32(*from), uint32(*to)); err != nil {
		fmt.Printf("Error %v", err)
	} else {
		fmt.Printf("Nonce is %d\n", nonce)
		fmt.Printf("Hash is %s\n", hash)
	}
	fmt.Printf("Time %v\n", time.Since(start))
}
