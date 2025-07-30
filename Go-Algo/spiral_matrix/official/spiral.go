// 官方解1看不懂 2解和自己的一样
package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}

	var (
		//矩阵的行数和列数
		rows, columns = len(matrix), len(matrix[0])
		//返回结果
		order = make([]int, rows*columns)
		//下标索引
		index = 0
		//各个边界
		top, bottom, left, right = 0, rows - 1, 0, columns - 1
	)
	//由外层到内层
	for top <= bottom && left <= right {
		//1、左上到右上遍历
		for column := left; column <= right; column++ {
			order[index] = matrix[top][column]
			index++
		}
		//2、右上到右下遍历
		for row := top + 1; row <= bottom; row++ {
			order[index] = matrix[row][right]
			index++
		}
		if left < right && top < bottom {
			//3、右下到左下遍历
			for column := right - 1; column >= left+1; column-- {
				order[index] = matrix[bottom][column]
				index++
			}
			//4、右下到左上遍历
			for row := bottom; row >= top+1; row-- {
				order[index] = matrix[row][left]
				index++
			}
		}
		top++
		bottom--
		left++
		right--
	}
	return order
}

func main() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Print(spiralOrder(matrix))
}
