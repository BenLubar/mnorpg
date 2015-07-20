package main

type names struct {
	ItemRarities    [numItemRarity]string
	MonsterRarities [numMonsterRarity]string
}

var namesByLanguage = map[string]*names{
	"en": &namesEnglish,
}

func Names(languages ...string) *names {
	for _, l := range languages {
		if n, ok := namesByLanguage[l]; ok {
			return n
		}
	}
	return &namesEnglish
}
