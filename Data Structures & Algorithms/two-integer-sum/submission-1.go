func twoSum(nums []int, target int) []int {
    // key: complement
	// value: index of number waiting for this complement
    m := make(map[int]int)
    
    for i, v := range nums {
        if idx, ok := m[v]; ok {  // current v completes a pair
            return []int{idx, i}
        }
        m[target-v] = i  // store what we need to find to pair with current v
    }
    return nil
}