package main

//
//import (
//"bufio"
//"bytes"
//"fmt"
//"os"
//)
//
//func main() {
//	reader := bufio.NewReader(os.Stdin)
//
//	tc := readNum(reader)
//
//	var buf bytes.Buffer
//
//	for tc > 0 {
//		tc--
//		n, k := readTwoNums(reader)
//		res := solve(n, k)
//		buf.WriteString(fmt.Sprintf("%d\n", res))
//	}
//
//	fmt.Print(buf.String())
//}
//
//func readString(reader *bufio.Reader) string {
//	s, _ := reader.ReadString('\n')
//	for i := 0; i < len(s); i++ {
//		if s[i] == '\n' || s[i] == '\r' {
//			return s[:i]
//		}
//	}
//	return s
//}
//
//func readNInt64s(reader *bufio.Reader, n int) []int64 {
//	res := make([]int64, n)
//	s, _ := reader.ReadBytes('\n')
//
//	var pos int
//
//	for i := 0; i < n; i++ {
//		pos = readInt64(s, pos, &res[i]) + 1
//	}
//
//	return res
//}
//
//func readInt64(bytes []byte, from int, val *int64) int {
//	i := from
//	var sign int64 = 1
//	if bytes[i] == '-' {
//		sign = -1
//		i++
//	}
//	var tmp int64
//	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
//		tmp = tmp*10 + int64(bytes[i]-'0')
//		i++
//	}
//	*val = tmp * sign
//	return i
//}
//
//func readInt(bytes []byte, from int, val *int) int {
//	i := from
//	sign := 1
//	if bytes[i] == '-' {
//		sign = -1
//		i++
//	}
//	tmp := 0
//	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
//		tmp = tmp*10 + int(bytes[i]-'0')
//		i++
//	}
//	*val = tmp * sign
//	return i
//}
//
//func readNum(reader *bufio.Reader) (a int) {
//	bs, _ := reader.ReadBytes('\n')
//	readInt(bs, 0, &a)
//	return
//}
//
//func readTwoNums(reader *bufio.Reader) (a int, b int) {
//	res := readNNums(reader, 2)
//	a, b = res[0], res[1]
//	return
//}
//
//func readThreeNums(reader *bufio.Reader) (a int, b int, c int) {
//	res := readNNums(reader, 3)
//	a, b, c = res[0], res[1], res[2]
//	return
//}
//
//func readNNums(reader *bufio.Reader, n int) []int {
//	res := make([]int, n)
//	x := 0
//	bs, _ := reader.ReadBytes('\n')
//	for i := 0; i < n; i++ {
//		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
//			x++
//		}
//		x = readInt(bs, x, &res[i])
//	}
//	return res
//}
//
//func solve(n int, k int) int {
//	if k > 30 {
//		return 0
//	}
//	buf := make([]int, k)
//	check := func(x int) bool {
//		buf[k-1] = n
//		buf[k-2] = x
//
//		it := k - 3
//		for it >= 0 {
//			buf[it] = buf[it+2] - buf[it+1]
//			if buf[it] < 0 || buf[it] > buf[it+1] {
//				return false
//			}
//
//			if buf[it] == 0 && it > 0 {
//				return false
//			}
//
//			it--
//		}
//		return it < 0
//	}
//	var res int
//	for i := n; i >= 0; i-- {
//		if check(i) {
//			res++
//		}
//	}
//	return res
//}
