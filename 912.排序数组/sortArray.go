package leetcode

import (
	"math"
	"math/rand"
	"strconv"
	"time"
)

// 快排模板1
func sortArray(nums []int) []int {
	quic_sort(nums, 0, len(nums)-1)
	return nums
}

func quic_sort(nums []int, l, r int) {
	if l >= r {
		return
	}
	rand.Seed(time.Now().Unix())
	pivot := rand.Intn(r-l) + l
	nums[r], nums[pivot] = nums[pivot], nums[r]
	i := l - 1
	for j := l; j < r; j++ {
		if nums[j] < nums[r] {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	i++
	nums[i], nums[r] = nums[r], nums[i]
	quic_sort(nums, l, i-1)
	quic_sort(nums, i+1, r)
}

// 快排模板2
func sortArray2(nums []int) []int {
	quic_sort2(nums, 0, len(nums)-1)
	return nums
}
func quic_sort2(nums []int, l, r int) {
	if l >= r {
		return
	}
	i, j := l, r
	nums[i], nums[(i+j)>>1] = nums[(i+j)>>1], nums[i]
	pivot := nums[i]
	for i < j {
		for i < j && pivot <= nums[j] {
			j--
		}
		nums[i] = nums[j]
		for i < j && nums[i] <= pivot {
			i++
		}
		nums[j] = nums[i]
	}
	nums[i] = pivot
	quic_sort2(nums, l, i-1)
	quic_sort2(nums, i+1, r)
}

// 快排模板3
func sortArray3(nums []int) []int {
	quic_sort3(nums, 0, len(nums)-1)
	return nums
}
func quic_sort3(nums []int, l, r int) {
	if l >= r {
		return
	}
	pivot := nums[(l+r)>>1]
	i, j := l-1, r+1
	for i < j {
		for {
			i++
			if nums[i] >= pivot {
				break
			}
		}
		for {
			j--
			if nums[j] <= pivot {
				break
			}
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	quic_sort3(nums, l, j)
	quic_sort3(nums, j+1, r)
}

// 归并排序
func sortArray4(nums []int) []int {
	merge_sort(nums, 0, len(nums)-1)
	return nums
}
func merge_sort(nums []int, l, r int) {
	if l < r {
		mid := (l + r) >> 1
		merge_sort(nums, l, mid)
		merge_sort(nums, mid+1, r)
		i, j := l, mid+1
		tmp := []int{}
		for i <= mid || j <= r {
			if i > mid || (j <= r && nums[j] < nums[i]) {
				tmp = append(tmp, nums[j])
				j++
			} else {
				tmp = append(tmp, nums[i])
				i++
			}
		}
		copy(nums[l:r+1], tmp)
	}
}

// 堆排序
func sortArray5(nums []int) []int {
	heap_sort(nums)
	return nums
}
func heap_sort(nums []int) []int {
	lens := len(nums) - 1
	for i := lens >> 1; i >= 0; i-- {
		down(nums, i, lens)
	}
	for j := lens; j >= 1; j-- {
		nums[0], nums[j] = nums[j], nums[0]
		lens--
		down(nums, 0, lens)
	}
	return nums
}

func down(nums []int, i, lens int) {
	max := i
	if i<<1+1 <= lens && nums[i<<1+1] > nums[max] {
		max = i<<1 + 1
	}
	if i<<1+2 <= lens && nums[i<<1+2] > nums[max] {
		max = i<<1 + 2
	}
	if max != i {
		nums[max], nums[i] = nums[i], nums[max]
		down(nums, max, lens)
	}
}

// 选择排序
func sortArray6(nums []int) []int {
	select_sort(nums)
	return nums
}
func select_sort(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[minIndex] {
				minIndex = j
			}
		}
		nums[i], nums[minIndex] = nums[minIndex], nums[i]
	}
}

// 插入排序
func sortArray7(nums []int) []int {
	insert_sort(nums)
	return nums
}
func insert_sort(nums []int) {
	lens := len(nums)
	for i := 1; i < lens; i++ {
		tmp := nums[i]
		j := i - 1
		for j >= 0 && nums[j] > tmp {
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1] = tmp
	}
}

// 冒泡排序
func sortArray8(nums []int) []int {
	bubble_sort(nums)
	return nums
}
func bubble_sort(nums []int) {
	lens := len(nums)
	for i := 0; i < lens-1; i++ {
		exchange := false
		for j := 0; j < lens-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				exchange = true
			}
		}
		if !exchange {
			break
		}
	}
}

// 希尔排序
func sortArray9(nums []int) []int {
	shell_sort(nums)
	return nums
}
func shell_sort(nums []int) {
	for gap := len(nums) >> 1; gap > 0; gap >>= 1 {
		for i := gap; i < len(nums); i++ {
			tmp := nums[i]
			j := i - gap
			for j >= 0 && tmp < nums[j] {
				nums[j+gap] = nums[j]
				j -= gap
			}
			nums[j+gap] = tmp
		}
	}
}

// 计数排序
func sortArray10(nums []int) []int {
	cnt := [100001]int{}
	for i := 0; i < len(nums); i++ {
		cnt[50000+nums[i]]++
	}
	idx := 0
	for i := 0; i < 100001; i++ {
		for cnt[i] > 0 {
			nums[idx] = i - 50000
			idx++
			cnt[i]--
		}
	}
	return nums
}

// 桶排序
func sortArray11(nums []int) []int {
	return bin_sort(nums, 4)
}
func bin_sort(nums []int, bin_num int) []int {
	min, max := nums[0], nums[0]
	for i := 0; i < len(nums); i++ {
		if min > nums[i] {
			min = nums[i]
		}
		if max < nums[i] {
			max = nums[i]
		}
	}
	bin := make([][]int, bin_num+1)
	for j := 0; j < len(nums); j++ {
		n := (nums[j] - min) / ((max - min + 1) / bin_num) // 这里计算位置有些问题，再用if做下判断。
		if n >= bin_num {
			n = bin_num - 1
		}
		bin[n] = append(bin[n], nums[j])
		k := len(bin[n]) - 2
		for k >= 0 && nums[j] < bin[n][k] {
			bin[n][k+1] = bin[n][k]
			k--
		}
		bin[n][k+1] = nums[j]
	}
	o := 0
	for p, q := range bin {
		for t := 0; t < len(q); t++ {
			nums[o] = bin[p][t]
			o++
		}
	}
	return nums
}

// 基数排序
func sortArray12(nums []int) []int {
	return radix_sort(nums)
}
func radix_sort(nums []int) []int {
	max := nums[0]
	for i := 0; i < len(nums); i++ {
		if max < nums[i] {
			max = nums[i]
		}
	}
	for j := 0; j < len(strconv.Itoa(max)); j++ {
		bin := make([][]int, 10)
		for k := 0; k < len(nums); k++ {
			n := nums[k] / int(math.Pow(10, float64(j))) % 10
			bin[n] = append(bin[n], nums[k])
		}
		m := 0
		for p := 0; p < len(bin); p++ {
			for q := 0; q < len(bin[p]);q++{
				nums[m] = bin[p][q]
				m++
			}
		}
	}
	return nums
}
