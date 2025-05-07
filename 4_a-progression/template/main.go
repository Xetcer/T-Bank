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

	fmt.Fprint(out, findPartCount(arr))
}

// isProgression проверяет, образует ли последовательность чисел a, b, c арифметическую прогрессию
func isProgression(a, b, c int) bool {
	return b-a == c-b
}

// findPartCount находит количество подотрезков с арифметическими прогрессиями
func findPartCount(arr []int) int {
	count := 0
	n := len(arr)
	for start := 0; start < n-2; start++ { // Проходим по всему массиву
		for end := start + 2; end < n; end++ { // Выбираем первую позицию начала подотрезка
			subArr := arr[start : end+1]         // Просматриваем весь подотрезок от 'start' до 'end'
			for j := 0; j < len(subArr)-2; j++ { // Проверяем все возможные сочетания внутри подотрезка
				a := subArr[j]
				b := subArr[j+1]
				c := subArr[j+2]
				if isProgression(a, b, c) { // Проверяем, является ли это арифметической прогрессией
					// fmt.Println(a, b, c)
					count++
				}
			}
		}
	}
	return count
}
