package main

import "diningphilosophers/chset"

import "fmt"

func main() {
	hashSet := chset.New()
	hashSet.Add("test")
	fmt.Printf("%t", hashSet.Contains("test"))
}
