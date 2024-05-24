package main

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
)

func main() {
	data1 := [][]int{{59}, {73, 41}, {52, 40, 53}, {26, 53, 6, 34}}
	byteValue, _ := os.ReadFile("/Users/myner/Desktop/GitHub/backend-challenge/files/hard.json")
	data2 := make([][]int, 0)
	json.Unmarshal(byteValue, &data2)
	data3 := make([][]int, 0)

	fmt.Println(resultMaxSum(data1)) // should 237
	fmt.Println(resultMaxSum(data2)) // should 7273
	fmt.Println(resultMaxSum(data3)) // should 0
}

func resultMaxSum(tree [][]int) int {
	if len(tree) == 0 {
		return 0
	}

	for row := len(tree) - 2; row >= 0; row-- {
		for col := 0; col < len(tree[row]); col++ {
			tree[row][col] += int(math.Max(float64(tree[row+1][col]), float64(tree[row+1][col+1])))
		}
	}

	return tree[0][0]
}
