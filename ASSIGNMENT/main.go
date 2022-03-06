package main

import "fmt"

type intSlice []int

func generateIntSlice() intSlice {
	result := intSlice{}
	for i := 0; i < 11; i++ {
		result = append(result, i)
	}
	return result
}

func (inS intSlice) printAns() {
	for _, num := range inS {
		if num%2 == 0 {
			fmt.Println(num, "is even")
		} else {
			fmt.Println(num, "is odd")
		}
	}
}

func main() {
	si := generateIntSlice()
	si.printAns()
}
