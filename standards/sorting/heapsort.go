package sorting

type heapSort struct {
}

func (hs *heapSort) Sort(array []int) {
	hs.BuildHeap(array)

	for length := len(array); length > 1; length-- {
		hs.RemoveTop(array, length)
	}
}

func (hs *heapSort) BuildHeap(array []int) {
	for i := len(array) / 2; i >= 0; i-- {
		hs.Heapify(array, i, len(array))
	}
}

func (hs *heapSort) RemoveTop(array []int, length int) {
	var lastIndex = length - 1
	array[0], array[lastIndex] = array[lastIndex], array[0]
	hs.Heapify(array, 0, lastIndex)
}

func (hs *heapSort) Heapify(array []int, root, length int) {
	var max = root
	var l, r = hs.Left(array, root), hs.Right(array, root)

	if l < length && array[l] > array[max] {
		max = l
	}

	if r < length && array[r] > array[max] {
		max = r
	}

	if max != root {
		array[root], array[max] = array[max], array[root]
		hs.Heapify(array, max, length)
	}
}

func (*heapSort) Left(array []int, root int) int {
	return (root * 2) + 1
}

func (*heapSort) Right(array []int, root int) int {
	return (root * 2) + 2
}
