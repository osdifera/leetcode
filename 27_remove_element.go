package main

import (
	"fmt"
)

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	index := 0
	for i := 0; i < len(nums); i++  {
		if nums[i] != val {
			nums[index] = nums[i]
			index++
		}
	}
	return index
}

func main27(){
	nums := []int{0,1,2,2,3,0,4,6,2}
	val := 2
	fmt.Println(removeElement(nums,val))
}
