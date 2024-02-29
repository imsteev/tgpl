package main

import "sync"

// Example on page 163 - "Composing types by struct embedding"
// Anonymous structs can't explicitly *define* methods...
// ...but struct embeds make it possible for them to *have* methods.
// For e.g:
var cache = struct {
	sync.Mutex
	mapping map[string]string
}{mapping: make(map[string]string)}

func Lookup(key string) string {
	cache.Lock()
	v := cache.mapping[key]
	cache.Unlock()
	return v
}
