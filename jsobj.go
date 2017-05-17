// +build js

//compile with gopherjs build -m

package main

import (
	"bytes"

	"github.com/empijei/identigen/identities"
	"github.com/gopherjs/gopherjs/js"
)

func main() {
	js.Global.Set("identigen", map[string]interface{}{
		"RandomPeople": RandomPeopleJS,
		"MainModule":   MainModuleJS,
	})
}

func RandomPeopleJS(minage, maxage int, count int) *js.Object {
	ppl, err := identities.RandomPeople(minage, maxage, count)
	if err != nil {
		panic(err)
	}
	return js.MakeWrapper(ppl)
}

func MainModuleJS(args map[string]interface{}) string {
	out := bytes.NewBuffer(nil)
	//<rant>
	//I officially hate javascript for the following code
	for key, value := range args {
		switch value := value.(type) {
		case float64:
			args[key] = int(value)
		}
	}
	//</rant>
	err := identities.MainModule(args, out)
	if err != nil {
		panic(err)
	}
	return string(out.Bytes())
}
