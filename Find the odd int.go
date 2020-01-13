import (
	"sort"
)
func FindOdd(seq []int) int {
	sort.Ints(seq)
	sum := 0
	for _, value := range seq {
		sum = value - sum
	}
	return sum
}
