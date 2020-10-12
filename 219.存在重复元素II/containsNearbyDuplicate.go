package leetcode

func containsNearbyDuplicate(nums []int, k int) bool {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j <= i+k && j < len(nums); j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}
	return false
}

func containsNearbyDuplicate2(nums []int, k int) bool {
	km := make(map[int]bool, k)
	var i, j int
	for j < len(nums) {
		for j < len(nums) && j-i <= k {
			if km[nums[j]] {
				return true
			} else {
				km[nums[j]] = true
			}
			j++
		}
		delete(km, nums[i])
		i++
	}
	return false
}

func containsNearbyDuplicate3(nums []int, k int) bool {
	m := make(map[int]int)
	for i, num := range nums {
		if _, ok := m[num]; ok {
			if i - m[num] <= k {
				return true
			}
		}
		m[num] = i
	}
	return false
}
