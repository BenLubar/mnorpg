package main

import "testing"

func TestNamesByLanguage(t *testing.T) {
	if n := Names("en"); n != &namesEnglish {
		t.Errorf("\"en\" ->\nhave %#v\nwant %#v", n, &namesEnglish)
	}
	if n := Names(); n != &namesEnglish {
		t.Errorf(" ->\nhave %#v\nwant %#v", n, &namesEnglish)
	}
	if n := Names("fake language"); n != &namesEnglish {
		t.Errorf("\"fake language\" ->\nhave %#v\nwant %#v", n, &namesEnglish)
	}
}

func TestNames(t *testing.T) {
	for l, n := range namesByLanguage {
		used := make(map[string]bool)
		for i, s := range n.ItemRarities {
			if s == "" {
				t.Errorf("language %q missing translation for item rarity %v", l, ItemRarity(i))
			} else if used[s] {
				t.Errorf("language %q has multiple item rarities named %q", l, s)
			} else {
				used[s] = true
			}
		}
		used = make(map[string]bool)
		for i, s := range n.MonsterRarities {
			if s == "" {
				t.Errorf("language %q missing translation for monster rarity %v", l, MonsterRarity(i))
			} else if used[s] {
				t.Errorf("language %q has multiple monster rarities named %q", l, s)
			} else {
				used[s] = true
			}
		}
	}
}
