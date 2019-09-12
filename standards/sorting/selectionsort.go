package sorting

type selectionSort struct{}

func (ss *selectionSort) Sort(a []int) {
	if len(a) <= 1 {
		return
	}

	maxIdx := 0
	for i, val := range a {
		if val >= a[maxIdx] {
			maxIdx = i
		}
	}
	a[len(a)-1], a[maxIdx] = a[maxIdx], a[len(a)-1]
	ss.Sort(a[:len(a)-1])
}
