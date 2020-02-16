package main

import "github.com/weeee9/blockchain/blockchain"

func main() {
	chain := blockchain.InitBlockchain()

	chain.AddBlock("First Block after Genesis")
	chain.AddBlock("Second Block after Genesis")
	chain.AddBlock("Third Block after Genesis")

	chain.Info()
}
