package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	fmt.Println("is", intersection(nums1, nums2))
}

func intersection(nums1 []int, nums2 []int) []int {
	m1, m2 := map[int]bool{}, map[int]bool{}

	for _, v := range nums1 {
		m1[v] = true
	}
	for _, v := range nums2 {
		m2[v] = true
	}
	ans := []int{}

	for k, _ := range m1 {
		if _, ok := m2[k]; ok {
			ans = append(ans, k)
		}
	}

	return ans
}
