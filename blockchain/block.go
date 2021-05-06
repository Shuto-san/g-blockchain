package blockchain

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

/*
Block is a block including timestamp, last hash, hash, data
*/
type Block struct {
	timestamp int64
	lastHash  string
	hash      string
	data      string
}

/*
NewBlock is create new block instance
*/
func NewBlock(timestamp int64, lastHash string, hash string, data string) *Block {
	return &Block{
		timestamp: timestamp,
		lastHash:  lastHash,
		hash:      hash,
		data:      data,
	}
}

func (block *Block) toString() string {
	return "block ---\n" +
		"Timestamp : " + strconv.FormatInt(block.timestamp, 10) + "\n" +
		"Last Hash : " + string([]rune(block.lastHash)[:10]) + "\n" +
		"Hash      : " + string([]rune(block.hash)[:10]) + "\n" +
		"Data      : " + block.data + "\n"
}

func genesis() *Block {
	block := NewBlock(time.Now().Unix(), "-------", "firsthash", "")
	return block
}

func mineBlock(lastBlock *Block, data string) *Block {
	timestamp := time.Now().Unix()
	lastHash := lastBlock.hash
	hash := hash(timestamp, lastHash, data)
	block := NewBlock(timestamp, lastHash, hash, data)
	return block
}

func hash(timestamp int64, lastHash string, data string) string {
	blockString := strconv.FormatInt(timestamp, 10) + lastHash + data
	sum := sha256.Sum256([]byte(blockString))
	return fmt.Sprintf("%x", sum)
}
