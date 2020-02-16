package blockchain

import (
	"fmt"
	"strconv"
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
	Nonce    int
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
	block := &Block{[]byte{}, []byte(data), prevHash, 0}
	pow := NewProof(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
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
	fmt.Printf("Block's PrevHash: %x\n", b.PrevHash)
	fmt.Printf("Block's Data: %s\n", b.Data)
	fmt.Printf("Block's Hahs: %x\n", b.Hash)

	pow := NewProof(b)
	fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
	fmt.Println()
}
