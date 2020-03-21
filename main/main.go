package main

import "fmt"

func DigitalRoot(n int) int {
	var sum = 0

	for {
		for n > 0 {
			sum += n % 10
			n /= 10
		}

		if sum >= 10 {
			n = sum
			sum = 0
		} else {
			break
		}
	}

	return sum
}

func main() {
	fmt.Println(DigitalRoot(16))
	fmt.Println(DigitalRoot(195))
	fmt.Println(DigitalRoot(992))
	fmt.Println(DigitalRoot(167346))
	fmt.Println(DigitalRoot(0))
}
