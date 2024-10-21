package main

import "fmt"

func twoSum(nums []int, target int) []int {
	//create a map that stores the indices of the elements
	indexMap := make(map[int]int)

	for i, num := range nums { //iterates over nums slice
		diff := target - num
		if j, found := indexMap[diff]; found { //checks if diff is in indexMap
			return []int{i, j} //returns answer
		}
		indexMap[num] = i //store the index of the current num
	}
	//if no solution, return nil
	return nil
}

func main() {
	//Test Case #1
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(nums, target)
	fmt.Println(result) //Output: [1, 0]

	//Test Case #2
	nums = []int{3, 2, 4}
	target = 6
	result = twoSum(nums, target)
	fmt.Println(result) //Output: [2, 1]

	//Test Case #3
	nums = []int{3, 3}
	target = 6
	result = twoSum(nums, target)
	fmt.Println(result) //Output: [1, 0]
}
