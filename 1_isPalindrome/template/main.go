package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Функция проверяет, является ли строка палиндромом
func isPalindrome(s string) bool {
	strLen := len(s)
	mid := strLen / 2

	for i := 0; i < mid; i++ {
		if s[i] != s[strLen-1-i] {
			return false
		}
	}
	return true
}

// Функция проверяет, является ли строка почти палиндромом
func almostPalindrome(s string) string {
	strLen := len(s)
	for i := 0; i < strLen; i++ {
		// Формируем строку, исключая i-й символ
		sub := strings.Join([]string{s[:i], s[i+1:]}, "")

		// Проверяем, является ли оставшаяся строка палиндромом
		if isPalindrome(sub) {
			return "YES"
		}
	}
	return "NO"
}

func main() {

	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Run(in, out)

}

func Run(in *bufio.Reader, out *bufio.Writer) {
	var str string
	fmt.Fscanln(in, &str)
	result := almostPalindrome(str)
	fmt.Fprint(out, result)
}
