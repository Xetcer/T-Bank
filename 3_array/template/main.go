package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {

	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Run(in, out)

}

func Run(in *bufio.Reader, out *bufio.Writer) {
	var n int
	var str string
	fmt.Fscanln(in, &n) // Получаем длину массива
	// fmt.Println("N", n)
	// Считываем строку целиком
	bytes, _ := in.ReadBytes('\n')
	str = string(bytes[:]) // отбрасываем символ переноса строки
	// fmt.Println("str", str)

	// Парсим строку с числами
	arr := make([]int, n)
	for i, valStr := range sort.StringSlice(strings.Fields(str)) {
		fmt.Sscanf(valStr, "%d", &arr[i])
	}
	maxUniqNum := maxUniqueElements(arr)
	fmt.Fprintln(out, maxUniqNum)
}

// Функция для нахождения максимального числа уникальных элементов
func maxUniqueElements(nums []int) int {
	seen := make(map[int]int) // Множество уникальных чисел

	for _, num := range nums {
		for num > 0 {
			seen[num]++
			num /= 2
		}
		seen[0]++ // Обязательно добавляем 0
	}
	// fmt.Println("for str nums", nums, seen)
	if len(seen) > len(nums) {
		return len(nums)
	} else {
		return len(seen)
	}
}
