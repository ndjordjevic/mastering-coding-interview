package arrays

// [0,3,4,31], [3,4,6,30] -> [0,3,3,4,4,6,30,31]
func mergeTwoSortedArrays(arr1 []int, arr2 []int) []int {
	merged := make([]int, 0, len(arr1)+len(arr2))

	i, j := 0, 0

	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			merged = append(merged, arr1[i])
			i++
		} else {
			merged = append(merged, arr2[j])
			j++
		}
	}

	for i < len(arr1) {
		merged = append(merged, arr1[i])
		i++
	}

	for j < len(arr2) {
		merged = append(merged, arr2[j])
		j++
	}

	return merged
}

// O(n+m), O(n+m)
