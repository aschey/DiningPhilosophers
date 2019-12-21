package main

import "fmt"

func main() {
	hashSet := NewConcurrentHashSet()
	hashSet.Add("test")
	hashSet.Add("test")
	fmt.Printf("%t %d", hashSet.Contains("test"), hashSet.Length())
}
