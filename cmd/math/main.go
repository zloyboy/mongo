package main

import "fmt"

func main() {
	var x1, v1, x2, v2 int
	fmt.Scan(&x1)
	fmt.Scan(&v1)
	fmt.Scan(&x2)
	fmt.Scan(&v2)

	s1 := x2 - x1
	s2 := (x2 + v2) - (x1 + v1)

	if x1 == x2 {
		fmt.Println("YES")
	} else if 0 <= s1 {
		if s2 < s1 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	} else {
		if s1 < s2 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
