package main

import (
	"fmt"

	"datastructs/structs"
)

func main() {
	dq := structs.Deque{}
	dq.PushFront(2)
	val, _ := dq.PopFront()
	fmt.Println(val)
}
