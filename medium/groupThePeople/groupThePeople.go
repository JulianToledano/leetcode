package main

func main() {
	groupThePeople([]int{3, 3, 3, 3, 3, 1, 3})
}
func groupThePeople(groupSizes []int) [][]int {
	var groups [][]int

	m := make(map[int][]int)
	for i, size := range groupSizes {
		m[size] = append(m[size], i)

		if len(m[size]) == size {
			groups = append(groups, m[size])
			delete(m, size)
		}
	}

	return groups
}
