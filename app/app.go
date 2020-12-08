package app

var Public string = "Public"
var private string = "private"

type App struct {
	Num int
}

func Average(s []int) int {
	sum := 0
	for _, v := range s {
		sum += v
	}
	sum /= len(s)
	return int(sum)
}
