package blockchain

import (
	"fmt"
	"testing"
)

var (
	data      = "bar"
	lastBlock = genesis()
	block     = mineBlock(lastBlock, data)
)

func TestBlock(t *testing.T) {
	fmt.Printf("%s", lastBlock.toString())
	fmt.Printf("%s", block.toString())

	t.Run("it sets the `data` to match the input", func(t *testing.T) {
		if lastBlock.Data != "" {
			t.Errorf("got:%s, expected:%s", lastBlock.Data, "")
		}
	})

	t.Run("it sets the `lastHash` to match the hash of the last block", func(t *testing.T) {
		if block.LastHash != lastBlock.Hash {
			t.Errorf("got:%s, expected:%s", block.LastHash, lastBlock.Hash)
		}
	})
}
