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
	m.Store("v", example{"map"})
	m.Store("d", example{"an"})
	m.Store("l", example{"tree"})
	m.Store("k", example{"of"})
	m.Store("b", example{"this"})
	m.Store("h", example{"example"})
	m.Store("c", example{"is"})

	fmt.Printf("Tree Map Size: %d\n", m.Len())

	v, ok := m.Load("c")
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
