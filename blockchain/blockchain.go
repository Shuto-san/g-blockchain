package blockchain

import (
	"log"
	"reflect"
)

/*
Blockchain is a chain of Blocks
*/
type Blockchain struct {
	chain []*Block
}

/*
NewBlockchain is create new blockchain instance
*/
func NewBlockchain() *Blockchain {
	chain := make([]*Block, 0)
	block := genesis()
	chain = append(chain, block)
	return &Blockchain{
		chain: chain,
	}
}

func (blockchain *Blockchain) addBlock(data string) *Block {
	block := mineBlock(blockchain.chain[len(blockchain.chain)-1], data)
	blockchain.chain = append(blockchain.chain, block)

	return block
}

func isValidChain(chain []*Block) bool {
	if !reflect.DeepEqual(chain[0], genesis()) {
		return false
	}

	for i := 1; i < len(chain); i++ {
		block := chain[i]
		lastBlock := chain[i-1]

		if block.lastHash != lastBlock.hash || block.hash != blockHash(block) {
			return false
		}
	}
	return true
}

func (blockchain *Blockchain) replaceChain(newChain []*Block) {
	if len(newChain) <= len(blockchain.chain) {
		log.Println("Received chain is not longer than the current chain.")
		return
	} else if !isValidChain(newChain) {
		log.Println("The received chain is not valid.")
		return
	}

	log.Println("Replacing blockchain with the new chain.")
	blockchain.chain = newChain
}
