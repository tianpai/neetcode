func productExceptSelf(nums []int) []int {
    length := len(nums)

    l := make([]int, length)  // left product of i
    l[0] = 1
    l[1] = nums[0]
    for i := 2; i < length; i++ {
        l[i] = l[i-1] * nums[i-1]
    }

    r := make([]int, length)  // right product of i
    r[length-1] = 1
    r[length-2] = nums[length-1]
    for j := length - 3; j >= 0; j-- {
        r[j] = r[j+1] * nums[j+1]
    }

    var res []int
    k := 0
    for k < length {
        res = append(res, l[k] * r[k])
        k++
    }
    return res
}
