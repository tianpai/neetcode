func twoSum(nums []int, target int) []int {
	// m stores: complement -> index
	// complement = target - current_value (the number we need to find to pair with current)
	m := make(map[int]int)

	// the question assume there is only exactly one pair exists.
	for i, v := range nums {
		// Check if current number v completes a pair with any previous number
		// If v exists in map, some earlier number at m[v] needed exactly this v to sum to target
		if idx, ok := m[v]; ok {
			return []int{idx, i}
		}
		// No match yet. Store what future numbers need to match with current v:
		// complement = target - v
		// If we later see this complement, we know current index i pairs with it
		m[target-v] = i
	}
	return nil
}