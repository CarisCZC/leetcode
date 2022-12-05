// https://leetcode.cn/problems/duplicate-zeros/
package main

func duplicateZeros(arr []int)  {
	n:=len(arr)
	top :=0
	i:=-1
	for top<n{
		i++
		if arr[i]!=0{
			top++
		}else{
			top+=2
		}
	}
	j := n-1
	// top=n+1 说明最后一位是0
	if top == n+1{
		arr[j] =0
		j--
		i--
	}
	for j>=0{
		arr[j] = arr[i]
		j--
		if arr[i] == 0{
			arr[j] = arr[i]
			j--
		}
		i--
	}
}

