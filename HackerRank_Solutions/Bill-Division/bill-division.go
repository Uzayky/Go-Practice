package Bill_Division

import "fmt"

// https://www.hackerrank.com/challenges/bon-appetit/problem

func BonAppetit(bill []int32, k int32, b int32) {

	var sum int32
	for i, v := range bill {
		if int32(i) != k {
			sum += v
		}
	}
	if sum/2 == b {
		fmt.Println("Bon Appetit")
	} else {
		fmt.Println(b - sum/2)
	}
}
