package main

import (
	"fmt"

	tm "github.com/g8rswimmer/go-tree-map"
)

type example struct {
	name string
}

func (e example) String() string {
	return e.name
}

func main() {
	fmt.Println("Tree Map Example Key: string Value: struct")
	m := tm.NewMap[string, example]()

	// insert into the map
	m.Put("v", example{"map"})
	m.Put("d", example{"an"})
	m.Put("l", example{"tree"})
	m.Put("k", example{"of"})
	m.Put("b", example{"this"})
	m.Put("h", example{"example"})
	m.Put("c", example{"is"})

	fmt.Printf("Tree Map Size: %d\n", m.Len())

	v, ok := m.Get("c")
	if !ok {
		panic("unable to find value")
	}
	fmt.Printf("Tree Map Get key[c] Value: %s\n", v)

	fmt.Println("Tree Map Ordered Iteration")
	handler := func(k string, v example) error {
		fmt.Printf("Key: %s Value: %s\n", k, v)
		return nil
	}
	if err := m.Range(handler); err != nil {
		panic(err)
	}

}
