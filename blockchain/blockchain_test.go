package blockchain

import (
	"reflect"
	"testing"
)

func TestBlockchain(t *testing.T) {
	t.Run("start with genesis block", func(t *testing.T) {
		bc := NewBlockchain()
		got := bc.Chain[0]
		expected := genesis()

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("got:%v, expected:%v", got, expected)
		}
	})

	t.Run("adds a new block", func(t *testing.T) {
		bc := NewBlockchain()
		data := "foo"
		bc.AddBlock(data)

		if bc.Chain[len(bc.Chain)-1].Data != data {
			t.Errorf("got:%s, expected:%s", bc.Chain[len(bc.Chain)-1].Data, data)
		}
	})

	t.Run("validates a valid chain", func(t *testing.T) {
		bc2 := NewBlockchain()
		bc2.AddBlock("foo")

		if isValidChain(bc2.Chain) == false {
			t.Errorf("chain is invalid")
		}
	})

	t.Run("invalidates a chain with a corrupt genesis block", func(t *testing.T) {
		bc2 := NewBlockchain()
		bc2.Chain[0].Data = "Bad data"

		if isValidChain(bc2.Chain) != false {
			t.Errorf("not working chain validation")
		}
	})

	t.Run("invalidates a corrupt chain", func(t *testing.T) {
		bc2 := NewBlockchain()
		bc2.AddBlock("foo")
		bc2.Chain[1].Data = "Not foo"

		if isValidChain(bc2.Chain) == true {
			t.Errorf("not working chain validation")
		}
	})

	t.Run("replaces the chain with a valid chain", func(t *testing.T) {
		bc := NewBlockchain()
		bc2 := NewBlockchain()
		bc2.AddBlock("goo")
		bc.replaceChain(bc2.Chain)

		if !reflect.DeepEqual(bc.Chain, bc2.Chain) {
			t.Errorf("got:%v, expected:%v", bc.Chain, bc2.Chain)
		}
	})

	t.Run("does not replace the chain with one of less than or equal to length", func(t *testing.T) {
		bc := NewBlockchain()
		bc2 := NewBlockchain()
		bc.AddBlock("foo")
		bc.replaceChain(bc2.Chain)

		if reflect.DeepEqual(bc.Chain, bc2.Chain) {
			t.Errorf("got:%v, expected:%v", bc.Chain, bc2.Chain)
		}
	})

}
