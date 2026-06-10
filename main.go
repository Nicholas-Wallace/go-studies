package main

import (
	"fmt"
	sort "go-studies/ordenacao"
)

func main()  {
	vet := []int{9,4,3,6,3,2,8,7,1,5}
	sort.QuickSort(vet, 0, len(vet)-1)
	fmt.Println(vet)
}