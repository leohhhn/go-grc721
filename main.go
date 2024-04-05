package grc721

import "github.com/gnolang/gno/examples/gno.land/p/demo/avl"

type GRC721 struct {
	name              string
	symbol            string
	owners            *avl.Tree // tokenID > std.Address
	balances          *avl.Tree // std.Address > # of owned tokens
	tokenApprovals    *avl.Tree // tokenID > std.Address
	operatorApprovals *avl.Tree // "OwnerAddress:OperatorAddress" -> bool
	tokenURIs         *avl.Tree // tokenID > URI
}

func NewGRC721Token(name, symbol string) *GRC721 {
	return &GRC721{
		name:              name,
		symbol:            symbol,
		owners:            avl.NewTree(),
		balances:          avl.NewTree(),
		tokenApprovals:    avl.NewTree(),
		operatorApprovals: avl.NewTree(),
		tokenURIs:         avl.NewTree(),
	}
}
