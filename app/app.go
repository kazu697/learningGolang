package app

func Average(s []int) int {
	sum := 0
	for _, v := range s {
		sum += v
	}
	sum /= len(s)
	return int(sum)
}
