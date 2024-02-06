package main

import "fmt"

func main() {
	sl := []int{}
	for i := 1; i <= 50; i++ {
		sl = append(sl, i)
	}
	fmt.Println("sl", sl)
	var ans []int
	if len(sl) >= 10 {
		for i := len(sl) - 1; i >= 0; i-- {
			ans = append(ans, sl[i])
			if len(ans) == 10 {
				break
			}
		}
	}
	Swap(ans)
	fmt.Println(ans)
}

func Swap(arr []int) {
	l := 0
	r := len(arr) - 1
	for l <= r {
		temP := arr[l]
		arr[l] = arr[r]
		arr[r] = temP
		l++
		r--
	}

}
