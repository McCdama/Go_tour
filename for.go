package main_for

import "fmt"

func main_for() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}
	fmt.Println("---------")
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}
	fmt.Println("---------")
	for {
		fmt.Println("Loop")
		break
	}
	fmt.Println("---------")
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
