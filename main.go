package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	len1 := len(nums1)
	len2 := len(nums2)
	targetNum := make([]int, 0, len2)

	i := 0
	j := 0
	index := 0

	for i < len1 && j < len2 {

		if nums1[i] < nums2[j] {
			targetNum = append(targetNum, nums1[i])
		} else {
			targetNum = append(targetNum, nums2[j])
		}
		index = index + 1
		i = i + 1
		j = j + 1
	}
	if i == len1 {
		for k := i; k < len2; k++ {
			targetNum = append(targetNum, nums2[k])
		}
	}
	if j == len2 {
		for k := j; k < len1; k++ {
			targetNum = append(targetNum, nums1[k])
		}
	}
	fmt.Println(targetNum)
	return 2.4
}

func insliceIndex() {

}

func main() {

	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	findMedianSortedArrays(nums1, nums2)
	//fmt.Println(findMedianSortedArrays("abcabcbb"))

}
