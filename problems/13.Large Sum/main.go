package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/dyusembaevd/project-euler/pkg"
)

func main() {
	solve()
}

func solve() {
	in := bufio.NewReader(os.Stdin)
	sum := "0"
	for i := 0; i < 100; i++ {
		str := pkg.ReadString(in)
		sum = stringsSum(sum, str)
	}

	fmt.Println(sum)

	n := sum[:10]
	num, _ := strconv.Atoi(n)
	fmt.Println(num)
}

func stringsSum(str1, str2 string) string {
	r1 := pkg.ReverseString(str1)
	r2 := pkg.ReverseString(str2)

	builder := strings.Builder{}
	var acc byte

	l1 := len(r1)
	l2 := len(r2)

	for i := 0; i < l1 && i < l2; i++ {
		sum := (r1[i] - '0') + (r2[i] - '0') + acc
		builder.WriteByte(sum%10 + '0')
		if sum > 9 {
			acc = 1
		} else {
			acc = 0
		}
	}

	var i int
	var ptr string
	if l1 > l2 {
		i = l2
		ptr = r1
	} else {
		i = l1
		ptr = r2
	}

	for ; i < len(ptr); i++ {
		sum := (ptr[i] - '0') + acc
		builder.WriteByte(sum%10 + '0')
		if sum > 9 {
			acc = 1
		} else {
			acc = 0
		}
	}

	if acc == 1 {
		builder.WriteByte(acc + '0')
	}

	return pkg.ReverseString(builder.String())
}
