// leetcode 832 反转图像

package main

//
func flipAndInvertImage(image [][]int) [][]int {
	for i:=0;i<len(image);i++{
		reverse(image[i]);
		for j:=0;j<len(image[i]);j++{
			image[i][j] = 1- image[i][j]
		}
	}
	return image;
}


//先逆序
func reverse(line []int) []int{
	for i,j:=0,len(line)-1;i<j;i,j = i+1,j-1{
		line[i],line[j] = line[j],line[i]
	}
	return line
}
