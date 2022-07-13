// leetcode 733 图像渲染

package main

//可以写一个递归

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	// indexs = append(indexs, index{sr, sc})
	// if indexs != nil {
	// tmp := indexs[0]
	old := image[sr][sc]
	image[sr][sc] = color
	if sc > 0 && image[sr][sc-1] == old && image[sr][sc-1] != color {

		floodFill(image, sr, sc-1, color)
	}
	if sc < len(image[0])-1 && image[sr][sc+1] == old && image[sr][sc+1] != color {
		floodFill(image, sr, sc+1, color)
	}
	if sr > 0 && image[sr-1][sc] == old && image[sr-1][sc] != color {
		floodFill(image, sr-1, sc, color)
	}
	if sr < len(image)-1 && image[sr+1][sc] == old && image[sr+1][sc] != color {
		floodFill(image, sr+1, sc, color)
	}
	// indexs = indexs[1:]
	// }
	return image
}
