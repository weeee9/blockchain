package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

// BlockChain for blockchain
type BlockChain struct {
	blocks []*Block
}

// Block type
type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

func main() {
	chain := InitBlockchain()

	chain.AddBlock("First Block after Genesis")
	chain.AddBlock("Second Block after Genesis")
	chain.AddBlock("Third Block after Genesis")

	chain.Info()
}

// InitBlockchain init a new blockchain
func InitBlockchain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}

// Genesis for genesis block
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

// CreateBlock create a new block
func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash}
	block.DeriveHash()
	return block
}

// DeriveHash ...
func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

// AddBlock add a new block to blockchain
func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	newBlock := CreateBlock(data, prevBlock.Hash)
	chain.blocks = append(chain.blocks, newBlock)
}

// Info printf all block's info in chain
func (chain *BlockChain) Info() {
	for _, b := range chain.blocks {
		b.Info()
	}
}

// Info print block's info
func (b *Block) Info() {
	fmt.Printf("Block's PrevHash:\t%x\n", b.PrevHash)
	fmt.Printf("Block's Data:\t\t%s\n", b.Data)
	fmt.Printf("Block's Hahs:\t\t%x\n\n", b.Hash)
}
