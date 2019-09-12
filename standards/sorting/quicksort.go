package sorting

type quickSort struct{}

func (qs *quickSort) Sort(a []int) {
	if len(a) <= 1 {
		return
	}
	// choose the median element as pivot
	pivotIdx := len(a) / 2
	i := 0
	for i < pivotIdx {
		if a[i] > a[pivotIdx] {
			// put the value on the right of the pivot
			a[pivotIdx], a[pivotIdx-1] = a[i], a[pivotIdx]
			pivotIdx--
		} else {
			i++
		}
	}
	qs.Sort(a[:pivotIdx])
	qs.Sort(a[pivotIdx:])
}
