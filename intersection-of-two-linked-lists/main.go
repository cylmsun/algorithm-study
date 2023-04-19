package main

import "fmt"

type nodeList struct {
	value any
	next  *nodeList
}

func intersection(l1 *nodeList, l2 *nodeList) *nodeList {
	// 检查有没有环
	l1Loop := checkLoop(l1)
	l2Loop := checkLoop(l2)

	// 一个有一个没有必然不相交
	if (l1Loop && !l2Loop) || (!l1Loop && l2Loop) {
		return nil
	}

	// 都没有环的话
	if !l1Loop && !l2Loop {
		p1 := l1
		p2 := l2
		for p1 != p2 {
			if p1 != nil && p2 != nil {
				p1 = p1.next
				p2 = p2.next
			}

			// 代表长度不同
			if p1 == nil {
				p1 = l2
			}
			if p2 == nil {
				p2 = l1
			}
		}

		// p1可能是nil,代表不相交，此时p2也是nil
		return p1
	}

	// 都是环的话
	if l1Loop && l2Loop {
		// todo
	}

	return nil
}

func checkLoop(l *nodeList) bool {
	slow := l.next
	fast := l.next.next

	for slow != fast {
		if slow == nil || fast == nil {
			return false
		}
		slow = slow.next
		fast = fast.next.next
	}
	return true
}

func main() {
	var a *nodeList
	fmt.Printf("sss: %t \n", a == nil)
	//intersection(nil, nil)
	//len1 := 6
	//len2 := 10
}
