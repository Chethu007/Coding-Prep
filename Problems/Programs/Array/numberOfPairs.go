package main

//func numberOfPairs(nums1 []int, nums2 []int, k int) int {
//	res := 0
//	for i := 0; i < len(nums1); i++ {
//		for j := 0; j < len(nums2); j++ {
//			if nums1[i]%(nums2[j]*k) == 0 {
//				res++
//			}
//		}
//	}
//	return res
//}

func numberOfPairs(nums1 []int, nums2 []int, k int) int64 {
	var res int64 = 0
	m := make(map[int]int)
	for i := 0; i < len(nums2); i++ {
		m[nums2[i]*k]++
	}
	for i := 0; i < len(nums1); i++ {
		for key, _ := range m {
			if nums1[i]%key == 0 {
				res++
			}
		}
	}
	return res
}

//Input: nums1 = [1,3,4], nums2 = [1,3,4], k = 1
//Output: 5

func main() {
	nums1 := []int{1, 3, 4}
	nums2 := []int{1, 3, 4}
	println("Res: ", numberOfPairs(nums1, nums2, 1))
}
