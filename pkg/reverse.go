package reverse

import (
	"strings"
)

func RevSlice(s []string) string {
	str := strings.Join(s, " ")
	return RevString(str)
}

func RevString(s string) string {
	length := len(s)
	b := make([]byte, length)
	b = []byte(s)
	for i :=0; i <length/2; i++{
		j := length-1-i
		b[i] , b[j] = s[j], s[i]
	}
	return string(b)

}