package slidingwindow

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinSizeSubArraySumV1(t *testing.T) {
	// Input: target = 7, nums = [2,3,1,2,4,3]
	// Output: 2
	// Explanation: The subarray [4,3] has the minimal length under the problem constraint.
	target := 7
	nums := []int{2, 3, 1, 2, 4, 3}
	expectedOutput := 2

	out := minSizeSubArraySumV1(target, nums)

	assert.Equal(t, expectedOutput, out)
}
