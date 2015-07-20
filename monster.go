//go:generate stringer -type MonsterRarity

package main

type MonsterRarity uint8

const (
	MonsterGray MonsterRarity = iota
	MonsterWhite
	MonsterBronze
	MonsterSilver
	MonsterGold
	MonsterPurple
	MonsterPlatinum
	numMonsterRarity
)
