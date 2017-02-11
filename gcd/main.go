package main

import "fmt"

func main() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)
	fmt.Println(gcd(a, b))
}

func gcd(a, b int) int {
	if a < b {
		tmp := a
		a = b
		b = tmp
	}

	for {
		r := a % b
		if r == 0 {
			return b
		}

		a = b
		b = r
	}
}
