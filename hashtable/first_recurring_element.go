package hashtable

func firstRecurringInt(arr []int) int {
	m := make(map[int]bool)

	for _, el := range arr {
		if _, prs := m[el]; prs {
			return el
		}
		m[el] = true
	}

	return -1
}
