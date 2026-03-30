func removeDuplicateValues(intSlice []int) []int {
    keys := make(map[int]bool)
    list := []int{}
    for _, entry := range intSlice {
        if _, value := keys[entry]; !value {
            keys[entry] = true
            list = append(list, entry)
        }
    }
    return list
}

func longestConsecutive(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    m := make(map[int]int) // map of a_i and length of this seq
    noDupNums := removeDuplicateValues(nums)
    for _, n := range noDupNums {
        m[n] = 1
    }

    var sk []int // bound by len(nums), starting key of a seq
    for _, n := range noDupNums{
        _, ok := m[n-1] 
        if !ok {
            sk = append(sk, n)
        }
    }

    for _, s := range sk {
        c := s
        for _, ok := m[c+1]; ok; _, ok = m[c+1] {
            m[s]++
            c++
        }
    }

    max := 0 //  length of the max seq
    for _, n := range noDupNums{
        if m[n] > max {
            max = m[n]
        }
    }
    return max
}