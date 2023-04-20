package main

import (
	"sort"
	"strings"
)

// sortedKeys returns the keys of a map in sorted order.
func sortedKeys[K string, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	return keys
}

// stringPart splits the string by the given character and returns the item
// at the given index.
func stringPart(value, split string, index int) string {
	parts := strings.Split(value, split)
	if index < 0 {
		index = len(parts) + index
	}
	return parts[index]
}
