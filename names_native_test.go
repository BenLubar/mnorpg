// +build !js

package main

import (
	"os"
	"testing"
)

func TestLanguages(t *testing.T) {
	var tests = []struct {
		LANGUAGE    string
		LC_MESSAGES string
		LC_ALL      string
		LANG        string
		Expect      []string
	}{
		{
			LANGUAGE:    "",
			LC_MESSAGES: "",
			LC_ALL:      "",
			LANG:        "",
			Expect:      []string{},
		},
		{
			LANGUAGE:    "",
			LC_MESSAGES: "",
			LC_ALL:      "en_US.UTF-8",
			LANG:        "",
			Expect:      []string{"en-US", "en"},
		},
	}

	for _, test := range tests {
		os.Setenv("LANGUAGE", test.LANGUAGE)
		os.Setenv("LC_MESSAGES", test.LC_MESSAGES)
		os.Setenv("LC_ALL", test.LC_ALL)
		os.Setenv("LANG", test.LANG)

		if l := Languages(); len(l) != len(test.Expect) {
			t.Errorf("test failed for %#v", test)
		} else {
			for i := range l {
				if l[i] != test.Expect[i] {
					t.Errorf("test failed for %#v", test)
					break
				}
			}
		}
	}
}
