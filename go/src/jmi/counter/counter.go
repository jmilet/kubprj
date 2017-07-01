package counter

import (
	"log"
)

// Run ...
func Run(id int, input chan string) {
	n := 0
	counter := make(map[string]int)

	for key := range input {
		counter[key]++
		n++
		if n%1000 == 0 || key == "/show" {
			sum := sumValues(counter)
			log.Printf("[%d] - %d - %d - %s", id, len(counter), sum, key)
		}
	}
}

func sumValues(m map[string]int) int {
	total := 0
	for _, v := range m {
		total += v
	}

	return total
}
