package stringmodule

import "github.com/karur4n-sandbox/github-actions-monorepo-go/coremodule"

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func ReverseHello() string {
	return Reverse(coremodule.Hello())
}
