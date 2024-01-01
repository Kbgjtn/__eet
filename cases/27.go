package cases

func RemoveElement(nums []int, val int) (r int) {
    for i := 0; i < len(nums); i++ {
        if nums[i] == val {
            // slice element and destructive append
            nums = append(nums[:i], nums[i+1:]...)
            i--
        }
    }

    return len(nums)
}
