package main

type MyCalendarTwo struct {
	Interval [][]int
}

func Constructor() MyCalendarTwo {
	var interval [][]int
	x := MyCalendarTwo{Interval: interval}
	return x
}

func (this *MyCalendarTwo) Book(start int, end int) bool {
	curInterval := this.Interval
	count := 0
	for _, v := range curInterval {
		if start == v[0] {
			count++
		} else if start < v[0] && end > v[0] {
			count++
		} else if start > v[0] && start < v[1] {
			count++
		}
	}
	if count > 1 {
		return false
	}
	return true
}

//func main() {
//	obj := Constructor()
//	param_1 := obj.Book(10, 20)
//}

/**
 * Your MyCalendarTwo object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
