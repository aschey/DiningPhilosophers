package main

import (
	cmap "github.com/orcaman/concurrent-map"
)

type ConcurrentHashSet struct {
	conmap cmap.ConcurrentMap
}

func NewConcurrentHashSet() ConcurrentHashSet {
	hashSet := ConcurrentHashSet{conmap: cmap.New()}
	//hashSet.conmap = cmap.New()
	return hashSet
}

func (hashSet ConcurrentHashSet) Add(value string) {
	hashSet.conmap.Set(value, nil)
}

func (hashSet ConcurrentHashSet) Contains(value string) bool {
	return hashSet.conmap.Has(value)
}

func (hashSet ConcurrentHashSet) Remove(value string) {
	hashSet.conmap.Remove(value)
}

func (hashSet ConcurrentHashSet) Length() int {
	return hashSet.conmap.Count()
}
