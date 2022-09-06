// leetcode 88, 合并两个有序数组
package main

// import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	tmpNums1 := make([]int, m)
	copy(tmpNums1, nums1)
	for i, j := 0, 0; i < m || j < n; {
		if i == m {
			for ; j < n; j++ {
				nums1[j+i] = nums2[j]
			}
			return
		}
		if j == n {
			for ; i < m; i++ {
				nums1[j+i] = tmpNums1[i]
			}
			return
		}
		if tmpNums1[i] <= nums2[j] {
			nums1[i+j] = tmpNums1[i]
			i++
		} else {
			nums1[i+j] = nums2[j]
			j++
		}
	}
}

// func main() {
// 	nums1 := []int{1, 3, 4, 0, 0, 0}
// 	nums2 := []int{2, 2, 6}
// 	merge(nums1, 3, nums2, 3)
// 	fmt.Println(nums1)
// }
