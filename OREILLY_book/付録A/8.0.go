package main

import (
	"fmt"
)

func main() {
	//あまり使用しない配列
	var nums [3]int = [3]int{1, 2, 3}
	fmt.Println(nums[0])
	fmt.Println(len(nums))

	var nums1 []int
	nums2 := []int{1, 2, 3}
	nums3 := nums[0:2]  //配列から
	nums4 := nums2[1:3] //スライスから

	fmt.Println(nums2[1])

	nums2[0] = 100

	fmt.Println(len(nums2))

	nums2 = append(nums2, 4)
}
