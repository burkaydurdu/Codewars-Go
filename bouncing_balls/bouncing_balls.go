package bouncing_balls

func BouncingBall(h, bounce, window float64) int {
	// control
	if h <= 0 || h <= window || bounce <= 0 || bounce >= 1 {
		return -1
	}

	bounceCount := 1

	for h > window {
		h = h * bounce
		if h > window {
			bounceCount += 2
		}
	}
	return bounceCount
}
