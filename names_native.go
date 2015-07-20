// +build !js

package main

import (
	"os"
	"strings"
)

// Languages returns the user's preferred languages in BCP 47 format in
// priority order.
func Languages() (langs []string) {
	add := func(lang string) {
		for _, l := range strings.Split(lang, ":") {
			// Remove encoding (we only support UTF-8).
			if i := strings.IndexRune(l, '.'); i != -1 {
				l = l[:i]
			}
			// Skip empty locales or the "C" locale.
			if l == "" || l == "C" {
				continue
			}
			// Add the locale.
			langs = append(langs, strings.Replace(l, "_", "-", -1))
			// Add the base language if it is a dialect.
			if i := strings.IndexRune(l, '_'); i != -1 {
				langs = append(langs, l[:i])
			}
		}
	}
	add(os.Getenv("LANGUAGE"))
	add(os.Getenv("LC_MESSAGES"))
	add(os.Getenv("LC_ALL"))
	add(os.Getenv("LANG"))
	return
}
