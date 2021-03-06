package leetcode

//Forward declaration of isBadVersion API.
//@param   version   your guess about first bad version
//@return 	 	      true if current version is bad
//		          false if current version is good
func isBadVersion(version int) bool { return true }

func firstBadVersion(n int) int {
	for i := 0; i < n; i++ {
		if isBadVersion(n) {
			return n
		}
	}
	return n
}

func firstBadVersion2(n int) int {
	left, right := 1, n

	for left < right {
		mid := left + (right-left)/2
		if isBadVersion(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}
