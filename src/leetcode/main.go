/**
 *Create by chensr on 2019/9/24
*/

package main

import (
	"leetcode/dynamicProgramming/63"
	"fmt"
)

func main() {
	temp := [][]int{{0}}
	result := dynamicProgramming.UniquePaths2(temp)
	fmt.Println(result)
}
