package main

import (
	"fmt"
	"./myadd"
)

func swap(x, y string) (string, string) {
	return y, x

}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	sum := myadd.Add(5, 6)
	fmt.Print(sum)
}
