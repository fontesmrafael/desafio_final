package main

import "fmt"

func main() {
	x := soma(12, 24, 36)
	y := multiplica(100, 100)
	w := subtrai(5, 20)
	z := divide(200)
	fmt.Println(x, y, w, z)
}
func soma(i ...int) int {
	total := 0
	for _, v := range i {
		total += v
	}
	return total
}
func subtrai(i ...int) int {
	total := 0
	for _, v := range i {
		total = v - total
	}
	return total
}
func multiplica(i ...int) int {
	total := 1
	for _, v := range i {
		total = total * v
	}
	return total
}
func divide(i ...int) int {
	total := 1
	for _, v := range i {
		total = v / total
	}
	return total
}
