func makeKey(s string) string {
	var c [26]int
	for i := 0; i < len(s); i++ {
		c[s[i]-'a']++
	}

	parts := make([]string, 26)
	for i, v := range c {
		parts[i] = strconv.Itoa(v)
	}
	return strings.Join(parts, "#")
}

func toStrList(idxList []int, s []string) []string {
	var r []string
	for i := range idxList {
		fmt.Println(s[i])
	}
	return r
}

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]int)
	for i, s := range strs {
		m[makeKey(s)] = append(m[makeKey(s)], i)
	}
	var r [][]string
	for _, e := range m {
		var sublist []string
		for _, idx := range e {
			sublist = append(sublist, strs[idx])
		}
		r = append(r, sublist)
	}
	return r
}
