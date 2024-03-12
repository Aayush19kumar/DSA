
func NextSmaller(arr []int) []int {
	ans := make([]int, len(arr))
	stack := make([]int, 0)
	for i := len(arr) - 1; i >= 0; i-- {
		for len(stack) > 0 && arr[i] <= stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			ans[i] = -1
		} else {
			ans[i] = stack[len(stack)-1]
		}
		stack = append(stack, arr[i])
	}
	return ans
}

func NextGreater(arr []int) []int {
	ans := make([]int, len(arr))
	stack := make([]int, 0)
	for i := len(arr) - 1; i >= 0; i-- {
		for len(stack) > 0 && arr[i] >= stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			ans[i] = -1
		} else {
			ans[i] = stack[len(stack)-1]
		}
		stack = append(stack, arr[i])
	}
	return ans
}
