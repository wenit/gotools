package strings

import (
	"fmt"
	"github.com/gofrs/uuid"
	"strconv"
	"strings"
)

var HexChars = []string{"a", "b", "c", "d", "e", "f",
	"g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s",
	"t", "u", "v", "w", "x", "y", "z", "0", "1", "2", "3", "4", "5",
	"6", "7", "8", "9", "A", "B", "C", "D", "E", "F", "G", "H", "I",
	"J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V",
	"W", "X", "Y", "Z"}

// Hex2Dec 16进制字符串转int
func Hex2Dec(val string) int {
	n, err := strconv.ParseUint(val, 16, 32)
	if err != nil {
		fmt.Println(err)
	}
	return int(n)
}

// 生成uuid
func UUID() string {
	uid, _ := uuid.NewV4()
	return uid.String()
}

// 获取短uuid，8位长度
func ShortUUID() string {
	s := UUID()
	s = strings.ReplaceAll(s, "-", "")
	bf := strings.Builder{}
	for i := 0; i < 8; i++ {
		s2 := s[i*4 : i*4+4]
		x := Hex2Dec(s2)
		bf.WriteString(HexChars[x%0x3e])
	}
	return bf.String()
}
