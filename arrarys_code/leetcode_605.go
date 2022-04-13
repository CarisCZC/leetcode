// leetcode 605, 种花问题
package main

// 直接找连续3个0版本
func canPlaceFlowers(flowerbed []int, n int) bool {
	tmpFlower := make([]int, 0, len(flowerbed)+2)
	tmpFlower = append(tmpFlower, 0)
	tmpFlower = append(tmpFlower, flowerbed...)
	tmpFlower = append(tmpFlower, 0)

	for i := 1; i < len(tmpFlower)-1; i++ {
		if tmpFlower[i-1] == 0 && tmpFlower[i] == 0 && tmpFlower[i+1] == 0 {
			tmpFlower[i] = 1
			n--
		}
	}
	if n <= 0 {
		return true
	}
	return false
}

// 官方解法，贪心算法
// func canPlaceFlowers(flowerbed []int, n int) bool {
// 	count := 0
// 	m := len(flowerbed)
// 	prev := -1
// 	for i := 0; i < m; i++ {
// 		if flowerbed[i] == 1 {
// 			if prev < 0 { //遇到的第一朵花,前面都是0,能种花的个数是（l-1）/2，而(l-1)=i
// 				count += i / 2
// 			} else { // 此时遇到的情况就是区间情况，中间有i-prev-1个元素,设为x，
// 				//如果是奇数，能种x-1/2，如果是偶数，也是x-1/2（四个位置只能种一朵）
// 				count += (i - prev - 2) / 2
// 			}
// 			prev = i
// 		}
// 	}
// 	if prev < 0 { //一朵花也没有种,此时种每个偶数位下标
// 		count += (m + 1) / 2
// 	} else { //处理最后一个区间
// 		count += (m - prev - 1) / 2
// 	}
// 	return count >= n
// }
