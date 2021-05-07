package blockchain

import (
	"reflect"
	"testing"
)

func TestBlockchain(t *testing.T) {
	t.Run("start with genesis block", func(t *testing.T) {
		bc := NewBlockchain()
		got := bc.chain[0]
		expected := genesis()

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("got:%v, expected:%v", got, expected)
		}
	})

	t.Run("adds a new block", func(t *testing.T) {
		bc := NewBlockchain()
		data := "foo"
		bc.addBlock(data)

		if bc.chain[len(bc.chain)-1].data != data {
			t.Errorf("got:%s, expected:%s", bc.chain[len(bc.chain)-1].data, data)
		}
	})

	t.Run("validates a valid chain", func(t *testing.T) {
		bc2 := NewBlockchain()
		bc2.addBlock("foo")

		if isValidChain(bc2.chain) == false {
			t.Errorf("chain is invalid")
		}
	})

	t.Run("invalidates a chain with a corrupt genesis block", func(t *testing.T) {
		bc2 := NewBlockchain()
		bc2.chain[0].data = "Bad data"

		if isValidChain(bc2.chain) != false {
			t.Errorf("not working chain validation")
		}
	})

	t.Run("invalidates a corrupt chain", func(t *testing.T) {
		bc2 := NewBlockchain()
		bc2.addBlock("foo")
		bc2.chain[1].data = "Not foo"

		if isValidChain(bc2.chain) == true {
			t.Errorf("not working chain validation")
		}
	})

	t.Run("replaces the chain with a valid chain", func(t *testing.T) {
		bc := NewBlockchain()
		bc2 := NewBlockchain()
		bc2.addBlock("goo")
		bc.replaceChain(bc2.chain)

		if !reflect.DeepEqual(bc.chain, bc2.chain) {
			t.Errorf("got:%v, expected:%v", bc.chain, bc2.chain)
		}
	})

	t.Run("does not replace the chain with one of less than or equal to length", func(t *testing.T) {
		bc := NewBlockchain()
		bc2 := NewBlockchain()
		bc.addBlock("foo")
		bc.replaceChain(bc2.chain)

		if reflect.DeepEqual(bc.chain, bc2.chain) {
			t.Errorf("got:%v, expected:%v", bc.chain, bc2.chain)
		}
	})

}
