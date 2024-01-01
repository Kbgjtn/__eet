package cases


/*
Example 1:

	Input: nums = [1,2,3,4]
	Output: 3
	Explanation: We have 3 arithmetic slices in nums: [1, 2, 3], [2, 3, 4] and [1,2,3,4] itself.

Example 2:

	Input: nums = [1]
	Output: 0
*/

func NumberOfArithmeticSlices(nums []int) (r int) {
    var nOfSlice, current int = 0, 0

    for i := 2; i < len(nums); i++ {
        if nums[i] - nums[i-1] == nums[i-1] - nums[i-2] {
            current++
            nOfSlice += current
        } else {
            current = 0
        }
    }

    return nOfSlice
}
