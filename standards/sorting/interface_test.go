package sorting

import "testing"

func isSorted(a []int) bool {
	if len(a) == 0 {
		return true
	}
	var prev int
	for i, v := range a {
		if i > 0 {
			if v < prev {
				return false
			}
		}
		prev = v
	}
	return true
}

func Test_isSorted(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "empty arr",
			args: args{a: []int{}},
			want: true,
		},
		{
			name: "sorted array",
			args: args{a: []int{1, 2, 3}},
			want: true,
		},
		{
			name: "non sorted array",
			args: args{a: []int{1, 2, 1}},
			want: false,
		},
		{
			name: "sorted array",
			args: args{a: []int{1, 1, 1}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSorted(tt.args.a); got != tt.want {
				t.Errorf("isSorted() = %v, want %v", got, tt.want)
			}
		})
	}
}

type testCaseArray struct {
	name string
	arr  []int
}

func getTestArrays() []testCaseArray {
	return []testCaseArray{
		{
			name: "empty array",
			arr:  []int{},
		},
		{
			name: "identical value array",
			arr:  []int{1, 1, 1},
		},
		{
			name: "already sorted array",
			arr:  []int{1, 2, 3},
		},
		{
			name: "non sorted array",
			arr:  []int{1, 3, 2},
		},
		{
			name: "non sorted array, with identical values",
			arr:  []int{1, 2, 1},
		},
	}
}
