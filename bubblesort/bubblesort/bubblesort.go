package math

func BubbleSort(sli []int) {
	for i := 0; i < len(sli); i++ {
		for j := 0; j < len(sli)-i-1; j++ {
			if sli[j] > sli[j+1] {
				Swap(sli, j)
			}
		}
	}
}

func Swap(sli []int, i int) {
	if i < len(sli) {
		temp := sli[i+1]
		sli[i+1] = sli[i]
		sli[i] = temp
	}
}