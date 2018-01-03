package cache

import "sync"

// Map is a go-built in map data structure with a read-write lock.
type Map struct {
	sync.RWMutex
	data map[string]interface{}
}

var (
	once sync.Once
	cMap Map
)
