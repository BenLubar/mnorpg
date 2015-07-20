//go:generate stringer -type ItemRarity

package main

type ItemRarity uint8

const (
	ItemGray ItemRarity = iota
	ItemWhite
	ItemGreen
	ItemBlue
	ItemPurple
	ItemOrange
	numItemRarity
)
