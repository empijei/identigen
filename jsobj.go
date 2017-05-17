// +build js

package main

import (
	"github.com/empijei/identigen/identities"
	"github.com/gopherjs/gopherjs/js"
)

func main() {
	js.Global.Set("identigen", map[string]interface{}{
		"RandomPeople": RandomPeopleJS,
	})
}

func RandomPeopleJS(minage, maxage int, count int) *js.Object {
	ppl, err := identities.RandomPeople(minage, maxage, count)
	if err != nil {
		panic(err)
	}
	return js.MakeWrapper(ppl)
}
