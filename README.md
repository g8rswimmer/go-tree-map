# Go Tree Map
This Map is implemented using a red-black tree which allows for the keys to be ordered.  When interating through the the maps pairs, key and value, those pairs will be return in descending order.  This can be useful when iteration through the map needs to be deterministic.  

For example, go's map implementation is a hash map.  This is great with O(1) operations like putting and getting pairs, but when iterating through the pairs the order is not deterministic.

The tree map uses generics to allow for [Ordered](https://pkg.go.dev/golang.org/x/exp/constraints#Ordered) keys and `any` values.
