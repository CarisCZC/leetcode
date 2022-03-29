//leetcode 53，最大子数组之和
package main

// 自己没写出来，该解法非原创,贪心算法
// func maxSubArray(nums []int) (max int) {
// 	max = nums[0]
// 	tmpMax := 0
// 	for _, v := range nums {
// 		if tmpMax < 0 {
// 			tmpMax = v
// 		} else {
// 			tmpMax += v
// 		}
// 		if tmpMax > max {
// 			max = tmpMax
// 		}
// 	}
// 	return
// }

// 分治
func maxSubArray(nums []int) int {
	return get(nums, 0, len(nums)-1).mSum
}

func pushUp(l, r Status) Status {
	iSum := l.iSum + r.iSum
	lSum := max(l.lSum, l.iSum+r.lSum)
	rSum := max(r.rSum, r.iSum+l.rSum)
	mSum := max(max(l.mSum, r.mSum), l.rSum+r.lSum)
	return Status{lSum, rSum, mSum, iSum}
}

func get(nums []int, l, r int) Status {
	if l == r {
		return Status{nums[l], nums[l], nums[l], nums[l]}
	}
	m := (l + r) >> 1
	lSub := get(nums, l, m)
	rSub := get(nums, m+1, r)
	return pushUp(lSub, rSub)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

type Status struct {
	lSum, rSum, mSum, iSum int
}

// 作者：LeetCode-Solution
// 链接：https://leetcode-cn.com/problems/maximum-subarray/solution/zui-da-zi-xu-he-by-leetcode-solution/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

// 高分示例 动态规划
//func maxSubArray(nums []int) int {
// 	max := nums[0]
// 	for i := 1; i < len(nums); i++ {
// 		if nums[i]+nums[i-1] > nums[i] {
// 			nums[i] += nums[i-1]
// 		}
// 		if nums[i] > max {
// 			max = nums[i]
// 		}
// 	}
// 	return max
// }

// func main() {
// 	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}

// 	fmt.Println(maxSubArray(nums))

// }
