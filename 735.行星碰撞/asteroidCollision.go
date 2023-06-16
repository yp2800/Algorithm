package leetcode

func asteroidCollision(asteroids []int) []int {
	var st []int
	for _, asteroid := range asteroids {
		alive := true
		for alive && asteroid < 0 && len(st) > 0 && st[len(st)-1] > 0 {
			alive = st[len(st)-1] < -asteroid
			if st[len(st)-1] <= -asteroid {
				st = st[:len(st)-1]
			}
		}
		if alive {
			st = append(st, asteroid)
		}
	}
	return st
}
