package main

import "fmt"

type Test struct {
	Name  string
	Value int
}

func Tfunc(a interface{}) {
	b := a.(Test)
	fmt.Printf("%+v\n", b)
}

func main() {
	hashSet := NewConcurrentHashSet()
	hashSet.Add("test")
	hashSet.Add("test2")
	fmt.Printf("%t %d\n", hashSet.Contains("test"), hashSet.Length())
	Tfunc(Test{"test", 1})
}
