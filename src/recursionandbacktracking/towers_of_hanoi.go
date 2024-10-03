package recursionandbacktracking

import (
	"fmt"
)

// Tower of hanoi problem

func TOH(n int, from, to, aux string) {
	if n == 1 {
		fmt.Printf("move disk %d from:%s to:%s\n", n, from, to)
		return
	}

	TOH(n-1, from, aux, to)
	fmt.Printf("move disk %d from:%s to:%s\n", n, from, to)
	TOH(n-1, aux, to, from)
}
