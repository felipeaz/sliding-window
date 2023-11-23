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

// lengthOfLongestSubstring returns the Longest Substring Without Repeating Characters (OPTIMIZED)
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	runes := []rune(s)
	// maps rune > position in array
	substringMap := makeSubstringMapFromRunes(runes, 0, 0)
	longestLen := 1

	startPos := 0
	endPos := 1
	currentLen := 1
	for endPos < len(runes) {
		chPos, isMapped := substringMap[runes[endPos]]
		// if the analyzed character repeats, another char map is initialized starting from the repeated elem position + 1
		// This optimizes the solution as we don't need to go through the entire array again if a char repeats
		// reducing the time complexity to O(N) whereas the common solution complexity would be O(N^N)
		if isMapped {
			startPos = chPos + 1
			substringMap = makeSubstringMapFromRunes(runes, startPos, endPos)
			currentLen = len(substringMap)
		} else {
			currentLen++
		}
		substringMap[runes[endPos]] = endPos
		endPos++

		if currentLen > longestLen {
			longestLen = currentLen
		}
	}
	return longestLen
}

func makeSubstringMapFromRunes(runes []rune, startPos, endPos int) map[rune]int {
	m := make(map[rune]int)
	for i := startPos; i <= endPos; i++ {
		m[runes[i]] = i
	}
	return m
}
