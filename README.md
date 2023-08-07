[![golangci-lint](https://github.com/g8rswimmer/go-tree-map/actions/workflows/golangci-lint.yml/badge.svg)](https://github.com/g8rswimmer/go-tree-map/actions/workflows/golangci-lint.yml)
[![go-test](https://github.com/g8rswimmer/go-tree-map/actions/workflows/go-test.yml/badge.svg)](https://github.com/g8rswimmer/go-tree-map/actions/workflows/go-test.yml)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)
# Go Tree Map
This Map is implemented using a [red-black tree](https://en.wikipedia.org/wiki/Red%E2%80%93black_tree) which allows for the keys to be ordered.  When interating through the the maps pairs, key and value, those pairs will be return in descending order.  This can be useful when iteration through the map needs to be deterministic.  

For example, go's map implementation is a hash map.  This is great with O(1) operations like putting and getting pairs, but when iterating through the pairs the order is not deterministic.

The tree map uses generics to allow for [Ordered](https://pkg.go.dev/golang.org/x/exp/constraints#Ordered) keys and `any` values.

Tree map function signatures are based off of [sync.Map](https://pkg.go.dev/sync) package.

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
```
#### Output
```
Tree Map Example Key: int Value: string
Tree Map Size: 7
Tree Map Get key[88] Value: put 5
Tree Map Get key[88] Value: change put
Tree Map Ordered Iteration
Key: 2 Value: put 4
Key: 4 Value: put 7
Key: 7 Value: put 2
Key: 11 Value: put 6
Key: 26 Value: put 3
Key: 78 Value: put 1
Key: 88 Value: change put
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
	fmt.Printf("Tree Map Get key[88] Value: %s\n", v)

	fmt.Println("Tree Map Ordered Iteration")
	handler := func(k string, v example) error {
		fmt.Printf("Key: %s Value: %s\n", k, v)
		return nil
	}
	if err := m.Range(handler); err != nil {
		panic(err)
	}

}
```
#### Output
```
Tree Map Example Key: string Value: struct
Tree Map Size: 7
Tree Map Get key[c] Value: is
Tree Map Ordered Iteration
Key: b Value: this
Key: c Value: is
Key: d Value: an
Key: h Value: example
Key: k Value: of
Key: l Value: tree
Key: v Value: map
```