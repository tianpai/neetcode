func productExceptSelf(nums []int) []int {
    length := len(nums)

    l := make([]int, length)  // left product of i
    l[0] = 1
    l[1] = nums[0]
    for i := 2; i < length; i++ {
        l[i] = l[i-1] * nums[i-1]
    }

    res := make([]int, length)  // result list
    right := 1
    for j := length - 1; j >= 0; j-- {
        res[j] = right * l[j]
        right = right * nums[j]
    }
    return res
}
