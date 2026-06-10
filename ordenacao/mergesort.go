package sort

import "fmt"

func MergeSort(v []int) {
	if len(v) != 1 {
		len_E := len(v) / 2
		len_D := len(v) - len_E

		fmt.Println(len_E)
		fmt.Println(len_D)

		v_E := make([]int, len_E)
		v_D := make([]int, len_D)

		for i := 0; i < len_E; i++ {
			v_E[i] = v[i]
		}
		for i := 0; i < len_D; i++ {
			v_D[i] = v[i+len_E]
		}

		fmt.Println(v_E)
		fmt.Println(v_D)

		MergeSort(v_E)
		MergeSort(v_D)
		merge(v, v_E, v_D)
		fmt.Println("merged:", v)
	}
}

func merge(new_vec []int, esq []int, dir []int) {

	// é preciso criar um index para iterar o esq e um para o dir
	iesq := 0
	idir := 0
	for iesq < len(esq) && idir < len(dir) {
		if esq[iesq] < dir[idir] {
			new_vec[iesq+idir] = esq[iesq]
			iesq++
		} else {
			new_vec[iesq+idir] = dir[idir]
			idir++
		}
	}
	for idir < len(dir) {
		new_vec[iesq+idir] = dir[idir]
		idir++
	}
	for iesq < len(esq) {
		new_vec[iesq+idir] = esq[iesq]
		iesq++
	}
}
