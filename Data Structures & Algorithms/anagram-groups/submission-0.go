func groupAnagrams(strs []string) [][]string {
	groups := make(map[string][]string)
	mask := make([]int, 26)
	key := []byte{}
	for _, word := range strs {
		for i := range word {
			mask[int(word[i] - 'a')]++
		}
		for _, freq := range mask {
			key = append(key, []byte(strconv.Itoa(freq))...)
			key = append(key, '.')
		}
		keyStr := string(key)
		groups[keyStr] = append(groups[keyStr], word)
		for i := range mask {
			mask[i] = 0
		}
		key = key[:0]
	}
	result := make([][]string, 0, len(groups))
	for _, group := range groups {
		result = append(result, group)
	}
	return result
}
