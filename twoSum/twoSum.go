package twoSum

func twoSum(nums []int, target int) []int {
	var bb []int
	for k, v := range nums {
		numsS := nums[k+1:]

		b := target - v
		for a, c := range numsS {
			if c == b {
				bb = append(bb, k)
				bb = append(bb, (k+1)+a)
				return bb
			}

		}

	}
	return bb
}
