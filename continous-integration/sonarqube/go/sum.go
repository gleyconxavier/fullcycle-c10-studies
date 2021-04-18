package _go

import "fmt"

func main(a, b int) (err error) {
	fmt.Println(sum(2, 2))

	return
}

func sum(a, b int) int {
	return a + b
}
