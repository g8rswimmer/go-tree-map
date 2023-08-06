package main

import (
	"fmt"

	tm "github.com/g8rswimmer/go-tree-map"
)

func main() {
	fmt.Println("Tree Map Example Key: int Value: string")
	m := tm.NewMap[int, string]()

	// insert into the map
	m.Put(78, "put 1")
	m.Put(7, "put 2")
	m.Put(26, "put 3")
	m.Put(2, "put 4")
	m.Put(88, "put 5")
	m.Put(11, "put 6")
	m.Put(4, "put 7")

	fmt.Printf("Tree Map Size: %d\n", m.Len())

	v, ok := m.Get(88)
	if !ok {
		panic("unable to find value")
	}
	fmt.Printf("Tree Map Get key[88] Value: %s\n", v)

	m.Put(88, "change put")
	v, ok = m.Get(88)
	if !ok {
		panic("unable to find value")
	}
	fmt.Printf("Tree Map Get key[88] Value: %s\n", v)

	fmt.Println("Tree Map Ordered Iteration")
	handler := func(k int, v string) error {
		fmt.Printf("Key: %d Value: %s\n", k, v)
		return nil
	}
	if err := m.Iterate(handler); err != nil {
		panic(err)
	}

}
