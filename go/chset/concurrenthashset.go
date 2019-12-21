package chset

import (
	cmap "github.com/orcaman/concurrent-map"
)

type ConcurrentHashSet struct {
	Map cmap.ConcurrentMap
}

func New() *ConcurrentHashSet {
	hashSet := new(ConcurrentHashSet)
	hashSet.Map = cmap.New()
	return hashSet
}

func (hashSet ConcurrentHashSet) Add(value string) {
	hashSet.Map.Set(value, nil)
}

func (hashSet ConcurrentHashSet) Contains(value string) bool {
	return hashSet.Map.Has(value)
}

func (hashSet ConcurrentHashSet) Remove(value string) {
	hashSet.Map.Remove(value)
}

func (hashSet ConcurrentHashSet) Length() int {
	return hashSet.Map.Count()
}
