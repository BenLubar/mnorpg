// +build js

package main

import "github.com/gopherjs/gopherjs/js"

// Languages returns the user's preferred languages in BCP 47 format in
// priority order.
func Languages() []string {
	return js.Global.Get("navigator").Get("languages").Interface().([]string)
}
