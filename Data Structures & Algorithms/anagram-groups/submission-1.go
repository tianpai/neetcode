func groupAnagrams(strs []string) [][]string {
	var r [][]string
	m := make(map[[26]int][]string)

	for _, s := range strs {
		var k [26]int
		for _, c := range s {
			fmt.Println(c - 'a')
			k[c-'a']++
		}
		m[k] = append(m[k], s)
	}
	for _, words := range m {
		r = append(r, words)
	}

	return r
}