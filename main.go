package main

import (
	"fmt"
	"github.com/ipfs/go-cid"
	"github.com/multiformats/go-multihash"
	"log"
	"os"

	carindex "github.com/ipld/go-car/v2/index"
)

func main() {
	///Users/lijunlong/code/venus-market/dagstore/baga6ea4seaqb3arf55fs4vglzzqyytm7mnga5rjt5ckdvcbdr4ffc4il5snskcy.full.idx
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
		return
	}
	index, err := carindex.ReadFrom(f)
	if err != nil {
		log.Fatal(err)
		return
	}

	iterableIdx, _ := index.(carindex.IterableIndex)
	iterableIdx.ForEach(func(m multihash.Multihash, u uint64) error {
		fmt.Println("PayloadCid:", cid.NewCidV0(m).String(), " Offset:", u)
		return nil
	})

	//lotus client retrieve --miner f01127678 QmNLmq479KD5kGySxf7gjWjmSsxmdoPXYnQLQUkqddDAT6 test.car
}
