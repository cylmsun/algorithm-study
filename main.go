package main

import "fmt"

func getRepeatCharIndex(str string) int {
	hMap := make(map[int32]int)
	for i, c := range str {
		hMap[c-'a']++
		fmt.Printf("   过程1: %d,%d,%d \n\n", hMap[c], c, i)
	}
	for i, c := range str {
		fmt.Printf("   过程2: %d,%d,%d \n\n", hMap[c], c, i)
		if hMap[c-'a'] == 1 {
			return i
		}
	}
	return -1
}

func main() {
	index := getRepeatCharIndex("aabb")
	fmt.Printf("结果: %d \n\n", index)
}
