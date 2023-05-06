package leetcode

func findKthLargest(nums []int, k int) int {
	// 快速排序
	//quickSort(nums, 0, len(nums)-1)
	//return nums[len(nums)-k]

	// 快速排序演进版
	quickSort2(nums, 0, len(nums)-1, len(nums)-k)
	return nums[len(nums)-k]

	// 大顶堆排序
	//heapSort(nums, len(nums))
	//return nums[len(nums)-k]

	// 堆排序演进版
	//heapSort2(nums, len(nums), k)
	//fmt.Printf("%+v", nums)
	//return nums[len(nums)-k]

}

// 快速排序
func quickSort(nums []int, start, end int) {
	if start < end {
		i, j := start, end
		povit := nums[(i+j)/2]
		for i <= j {
			for nums[i] < povit {
				i++
			}
			for nums[j] > povit {
				j--
			}
			if i <= j {
				nums[i], nums[j] = nums[j], nums[i]
				i++
				j--
			}
			if start < j {
				quickSort(nums, start, j)
			}
			if end > i {
				quickSort(nums, i, end)
			}
		}
	}
	return
}

// 快速排序演进版
func quickSort2(nums []int, start, end, k int) int {
	if start < end {
		i, j := start, end
		povit := nums[(i+j)/2]
		for i <= j {
			for nums[i] < povit {
				i++
			}
			for nums[j] > povit {
				j--
			}
			if i <= j {
				nums[i], nums[j] = nums[j], nums[i]
				i++
				j--
			}
			if i+1 == k {
				return nums[i]
			}
			if start < j {
				quickSort2(nums, start, j, 0)
			}
			if end > i {
				quickSort2(nums, i, end, 0)
			}
		}
	}
	return 0
}

// 堆排序 - 大顶堆
func heapify(nums []int, n int, i int) {
	max := i
	l := 2*i + 1
	r := 2*i + 2
	if l < n && nums[l] > nums[max] {
		max = l
	}
	if r < n && nums[r] > nums[max] {
		max = r
	}
	if max != i {
		nums[max], nums[i] = nums[i], nums[max]
		heapify(nums, n, max)
	}
}

func heapSort(nums []int, n int) {
	//调整堆
	for i := n/2 - 1; i >= 0; i-- {
		heapify(nums, n, i)
	}
	// 把堆顶和最后一个元素交换，再调整堆
	for i := n - 1; i > 0; i-- {
		nums[i], nums[0] = nums[0], nums[i]
		heapify(nums, i, 0)
	}
}

// 堆排序演进版
// 堆排序 - 大顶堆
func heapify2(nums []int, n int, i int) {
	max := i
	l := 2*i + 1
	r := 2*i + 2
	if l < n && nums[l] > nums[max] {
		max = l
	}
	if r < n && nums[r] > nums[max] {
		max = r
	}
	if max != i {
		nums[max], nums[i] = nums[i], nums[max]
		heapify2(nums, n, max)
	}
}

func heapSort2(nums []int, n, k int) {
	//调整堆
	for i := n/2 - 1; i >= 0; i-- {
		heapify2(nums, n, i)
	}
	// 把堆顶和最后一个元素交换，再调整堆
	for i := n - 1; i >= n-k; i-- {
		nums[i], nums[0] = nums[0], nums[i]
		heapify2(nums, i, 0)
	}
}
