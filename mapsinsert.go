package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["ans"] = 34
	fmt.Println("the value", m["ans"])
	m["ans"] = 45
	fmt.Println("the value ", m["ans"])
	delete(m, "ans")
	fmt.Println("the value", m["ans"])

}
