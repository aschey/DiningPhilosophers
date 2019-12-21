package main

import "diningphilosophers/chset"

import "fmt"

func main() {
	hashSet := chset.New()
	hashSet.Add("test")
	hashSet.Add("test")
	fmt.Printf("%t %d", hashSet.Contains("test"), hashSet.Length())
}
