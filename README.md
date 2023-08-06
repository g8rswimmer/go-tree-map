# Go Tree Map
This Map is implemented using a [red-black tree](https://en.wikipedia.org/wiki/Red%E2%80%93black_tree) which allows for the keys to be ordered.  When interating through the the maps pairs, key and value, those pairs will be return in descending order.  This can be useful when iteration through the map needs to be deterministic.  

For example, go's map implementation is a hash map.  This is great with O(1) operations like putting and getting pairs, but when iterating through the pairs the order is not deterministic.

The tree map uses generics to allow for [Ordered](https://pkg.go.dev/golang.org/x/exp/constraints#Ordered) keys and `any` values.

## Examples
### Key: Int Value: String
```go
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
```

### Key: string Value: struct
```go
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
	fmt.Printf("Tree Map Get key[88] Value: %s\n", v)

	fmt.Println("Tree Map Ordered Iteration")
	handler := func(k string, v example) error {
		fmt.Printf("Key: %s Value: %s\n", k, v)
		return nil
	}
	if err := m.Iterate(handler); err != nil {
		panic(err)
	}

}
```