package main

import (
	"container/ring"
	"fmt"
)

func joseph(r *ring.Ring, m int) *ring.Ring {
	if m == 1 {
		return r
	}

	i := 1
	for r.Len() > 1 {
		fmt.Printf("    单次循环开始, value:%d, r.len(): %d \n", r.Value, r.Len())
		if i != m {
			i++
			r = r.Next()
		} else {
			// 删除当前元素 === 前一个的next指向下一个
			r = r.Prev()
			s := r.Unlink(1)

			fmt.Printf("    单次循环结束,当前value:%d, 删除的是:%d \n\n\n", r.Value, s.Value)
			i = 1
			// 恢复之前的prev操作，确保正常
			r = r.Next()
		}
	}

	return r
}

func main() {
	fmt.Println("开始...")
	// 初始化
	r := ring.New(6)
	for i := 0; i < r.Len(); i++ {
		r.Value = i + 1
		r = r.Next()
	}
	fmt.Printf("长度: %d, 起始value: %d \n\n", r.Len(), r.Value)

	ans := joseph(r, 5)

	fmt.Println("检查结果...")
	// 检查结果
	//for i := 0; i < ans.Len(); i++ {
	//	fmt.Println(ans.Value)
	//	ans = ans.Next()
	//}
	fmt.Printf("结果值: %d \n", ans.Next().Value)
}
