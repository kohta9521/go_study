package main

import "fmt"

func main() {
	// var nums [3]int = [3]int(1, 2, 3)

	// fmt.Println(nums[0])
	// fmt.Println(len(nums))


	//スライスの変数宣言
	var nums1 []int
	//1, 2, 3,の要素を持つスライスを作成して代入
	nums2 := [] int{1, 2, 3}
	//あるいは既存の配列やスライスからも範囲アクセスでスライスを作成
	nums3 := nums[0:2]//配列から
	nums4 := nums2[1:3]//スライスから

	//配列と同じようにブラケットで要素取得可能
	fmt.Println(nums2[1]) // 2

	nums2[0] = 100
	fmt.Println(len(nums2)) // 3
	nums2 = append(nums2, 4)

}