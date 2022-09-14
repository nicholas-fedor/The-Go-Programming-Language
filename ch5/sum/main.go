// see page 142

package main

func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

func main() {
	
}

