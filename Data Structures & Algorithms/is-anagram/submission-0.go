func isAnagram(s string, t string) bool {
	m := make([]int, 26)
	for i := range s {
		m[int(s[i] - 'a')]++
	}
	for i := range t {
		m[int(t[i] - 'a')]--
	}
	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}
