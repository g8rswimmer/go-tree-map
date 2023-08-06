package main

import (
	"fmt"

	tm "github.com/g8rswimmer/go-tree-map"
)

func main() {
	fmt.Println("Tree Map Example Key: int Value: string")
	m := tm.NewMap[int, string]()

	// insert into the map
	m.Store(78, "put 1")
	m.Store(7, "put 2")
	m.Store(26, "put 3")
	m.Store(2, "put 4")
	m.Store(88, "put 5")
	m.Store(11, "put 6")
	m.Store(4, "put 7")

	fmt.Printf("Tree Map Size: %d\n", m.Len())

	v, ok := m.Load(88)
	if !ok {
		panic("unable to find value")
	}
	fmt.Printf("Tree Map Get key[88] Value: %s\n", v)

	m.Store(88, "change put")
	v, ok = m.Load(88)
	if !ok {
		panic("unable to find value")
	}
	fmt.Printf("Tree Map Get key[88] Value: %s\n", v)

	fmt.Println("Tree Map Ordered Iteration")
	handler := func(k int, v string) error {
		fmt.Printf("Key: %d Value: %s\n", k, v)
		return nil
	}
	if err := m.Range(handler); err != nil {
		panic(err)
	}

}
