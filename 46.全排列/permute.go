package leetcode

// 回溯
func permute(nums []int) [][]int {
	n := len(nums)

	var res, path = make([][]int, 0), make([]int, 0)
	var used = make([]bool, n)

	var dfs func()
	dfs = func() {
		if len(path) == n {
			var temp = make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}

		for i := range nums {
			if used[i] {
				continue
			}

			path = append(path, nums[i])
			used[i] = true
			dfs()
			//回溯的过程中，将当前的节点从path中删除
			path = path[:len(path)-1]
			used[i] = false
		}
	}
	dfs()
	return res
}
