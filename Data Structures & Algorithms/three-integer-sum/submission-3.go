import "slices"

func threeSum(nums []int) [][]int {
    slices.Sort(nums)
    // fix a_i, and start as a_j and a_k where i < j
    var res [][]int
    for i := 0; i < len(nums) -2; i++ {
        if nums[i] > 0{ // the rest nums must be >= 0, then sum > 0
            break
        }
        // start from the second iteration, skip duplicated nums[i] value
        if i > 0 && nums[i] == nums[i-1]{
			continue
		}

        j := i+1
        k := len(nums)-1
        for j < k {
            sum := nums[i] + nums[j] + nums[k]
            if sum == 0 {
                res = append(res, []int{nums[i], nums[j], nums[k]})
                // move j and k here but then check duplicated valus
                j++
                k--
                // keep move j and k until the tuple is different
                // to avoid duplicated tuples
                // NOTE: nums[j-1] and nums[k+1] are previous pairs used 
                // before we move j and k above.
                for j < k && nums[j] == nums[j-1]{
                    j++
                }
                for j < k && nums[k] == nums[k+1]{
                    k--
                }
            } else if sum < 0 { // j is too small but k is at max
                j++ 
            } else { // j is smallest but k is too big
                k--
            }
        }
    }
    // duplications handled in the loop
    // slices.SortFunc(res, slices.Compare)
    // res = slices.CompactFunc(res, slices.Equal)
    return res
}
