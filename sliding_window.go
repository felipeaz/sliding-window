package slidingwindow

// minSizeSubArraySumV1 returns the minimal length of a  subarray whose sum is greater than or equal to target.
// If there is no such subarray, returns 0 instead.
func minSizeSubArraySumV1(target int, nums []int) int {
	arrSize := len(nums)
	minLen := 0
	for startPos := 0; startPos < arrSize; startPos++ {
		sum := nums[startPos]
		endPos := startPos + 1
		for sum < target && endPos < arrSize {
			sum += nums[endPos]
			endPos++
		}
		if sum >= target {
			subArrayLen := len(nums[startPos:endPos])
			if minLen == 0 {
				minLen = subArrayLen
				continue
			}
			if subArrayLen < minLen {
				minLen = subArrayLen
			}
		}
	}
	return minLen
}
