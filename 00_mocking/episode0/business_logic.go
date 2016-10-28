package episode0

import "fmt"

func BusinessLogic(h HashTable) {
	fmt.Println("Business logic doing stuff")
	h.Set("hello", []byte("world"))
}
