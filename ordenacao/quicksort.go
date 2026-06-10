package sort

import "fmt"

func QuickSort(v []int, ini int, fim int) {
	if ini < fim {
		idxPivot := Partition(v, ini, fim)
		QuickSort(v, ini, idxPivot -1)
		QuickSort(v, idxPivot + 1, fim)
	}
}

func Partition(v []int, ini int, fim int) int {
	pivot := v[fim]
	fmt.Println(pivot)
	idxPivot := ini
	for i := ini; i < fim; i++{
		if  v[i] < pivot{
			v[i], v[idxPivot] = v[idxPivot], v[i] 
			idxPivot++
		}
	}
	v[fim], v[idxPivot] = v[idxPivot], v[fim]
	fmt.Println("partition: " ,v[ini:fim])
	fmt.Println(v)
	fmt.Println("--------------------")
	return idxPivot
}