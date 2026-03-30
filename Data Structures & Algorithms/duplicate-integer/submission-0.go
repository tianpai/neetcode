func hasDuplicate(nums []int) bool {
    m := make(map[int]int)
    for _, v := range nums {
        m[v] = m[v]+1
        if m[v] > 1 {
            return true
        }
    }
    return false
}