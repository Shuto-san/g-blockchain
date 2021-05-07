package blockchain

import (
	"log"
	"reflect"
)

/*
Blockchain is a chain of Blocks
*/
type Blockchain struct {
	Chain []*Block
}

/*
NewBlockchain is create new blockchain instance
*/
func NewBlockchain() *Blockchain {
	chain := make([]*Block, 0)
	block := genesis()
	chain = append(chain, block)
	return &Blockchain{
		Chain: chain,
	}
}

func (blockchain *Blockchain) AddBlock(data string) *Block {
	block := mineBlock(blockchain.Chain[len(blockchain.Chain)-1], data)
	blockchain.Chain = append(blockchain.Chain, block)

	return block
}

func isValidChain(chain []*Block) bool {
	if !reflect.DeepEqual(chain[0], genesis()) {
		return false
	}

	for i := 1; i < len(chain); i++ {
		block := chain[i]
		lastBlock := chain[i-1]

		if block.LastHash != lastBlock.Hash || block.Hash != blockHash(block) {
			return false
		}
	}
	return true
}

func (blockchain *Blockchain) replaceChain(newChain []*Block) {
	if len(newChain) <= len(blockchain.Chain) {
		log.Println("Received chain is not longer than the current chain.")
		return
	} else if !isValidChain(newChain) {
		log.Println("The received chain is not valid.")
		return
	}

	log.Println("Replacing blockchain with the new chain.")
	blockchain.Chain = newChain
}
