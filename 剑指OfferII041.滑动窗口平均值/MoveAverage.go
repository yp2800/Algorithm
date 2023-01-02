package leetcode

type MovingAverage struct {
	size, sum int
	nums      []int
}

/** Initialize your data structure here. */
func Constructor(size int) MovingAverage {
	return MovingAverage{
		size: size,
	}
}
func (this *MovingAverage) Next(val int) float64 {
	if len(this.nums) == this.size {
		this.sum -= this.nums[0]
		this.nums = this.nums[1:]
	}
	this.sum += val
	this.nums = append(this.nums, val)
	return float64(this.sum) / float64(len(this.nums))
}
