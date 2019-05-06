package main

import (
	"fmt"
	"sort"
)

var list = []int{0,13,21,3,2,4,9,11,70}

func main()  {
	sort.Ints(list)
	fmt.Println(list)
}

