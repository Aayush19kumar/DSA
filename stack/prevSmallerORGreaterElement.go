
func PreSmaller(arr []int) []int {
	ans := make([]int, len(arr))
	stack := make([]int, 0, len(arr))

	for i := 0; i < len(arr); i++ {
		for len(stack) > 0 && stack[len(stack)-1] >= arr[i] {
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

func PrevGreater(arr []int) []int {
	ans := make([]int, len(arr))
	stack := make([]int, 0)

	for i := 0; i < len(arr); i++ {
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
